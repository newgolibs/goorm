package goorm

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
)

/**
开启事务
*/
func (this *Pdo) Begin() {
	var err error
	this.tx, err = this.Pdoconfig.SqldbPool().Begin()
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		panic(fmt.Sprintf("\033[41;36merr:%+v %+v:%+v\033[0m\n", []interface{}{err}, file, line))
	}
}

/**    当运行中，有一条sql错误了，那么回滚，在这个事务期间的所有操作全部报废    */
func (this *Pdo) Rollback() error {
	// 根本没数据库链接，返回
	if this.tx == nil {
		return nil
	}
	var err error
	err = this.tx.Rollback()
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		panic(fmt.Sprintf("\033[41;36merr:%+v %+v:%+v\033[0m\n", []interface{}{err}, file, line))
	}
	this.tx = nil
	return err
}

/**
提交事务
*/
func (this *Pdo) Commit() error {
	// 根本没数据库链接，返回
	if this.tx == nil {
		return nil
	}
	var err error
	err = this.tx.Commit()
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		panic(fmt.Sprintf("\033[41;36merr:%+v %+v:%+v\033[0m\n", []interface{}{err}, file, line))
	}
	// 再开一条事务
	this.tx = nil
	return err
}

/**    执行Query方法，返回rows    */
func (this *Pdo) query(sqlstring string, bindarray []interface{}) (*sql.Rows, []interface{}, []sql.RawBytes, []string) {
	// 如果数据库还没链接，那么初始化一下
	if this.tx == nil {
		this.Begin()
	}
	rows, err := this.tx.Query(sqlstring, bindarray...)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		panic(fmt.Sprintf("\033[41;36merr:%+v %+v:%+v\033[0m\n", []interface{}{err}, file, line))
	}

	columns, _ := rows.Columns()

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	return rows, scanArgs, values, columns
}

/**
查询多行数据
*/
func (this *Pdo) SelectAll(sqlstring string, bindarray []interface{}) []map[string]string {
	rows, scanArgs, values, columns := this.query(sqlstring, bindarray)
	defer rows.Close()
	// 这个map用来存储一行数据，列名为map的key，map的value为列的值
	var rowMaps_All []map[string]string
	var rowMap = make(map[string]string)
	for rows.Next() {
		rows.Scan(scanArgs...)
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col != nil {
				rowMap[columns[i]] = string(col)
			}
		}
		// 因为rowmap的底层是一个value是引用，所以会导致后面的覆盖了前面，写的始终是一个值
		rowMap_temp := make(map[string]string)
		marshal, _ := json.Marshal(rowMap)
		json.Unmarshal(marshal, &rowMap_temp)
		rowMaps_All = append(rowMaps_All, rowMap_temp)
	}
	return rowMaps_All
}

/**    查询多行数据，返回struct对象的数组    */
func (this *Pdo) SelectallObject(sql string, bindarray []interface{}, orm_ptr interface{}) {
	orm_ptr_ref := reflect.ValueOf(orm_ptr)
	orm_ptr_value := orm_ptr_ref.Elem()
	if orm_ptr_value.Kind() != reflect.Slice {
		panic("传入的类型应该是一个切片：[]xxxx，现在是：" + orm_ptr_value.Kind().String())
	}
	all := this.SelectAll(sql, bindarray)
	orm_ptr_type := reflect.TypeOf(orm_ptr)
	for _, item := range all {
		one_orm_ptr := reflect.New(orm_ptr_type.Elem().Elem())
		Map_to_struct(item, one_orm_ptr.Interface())
		orm_ptr_value.Set(reflect.Append(orm_ptr_value, one_orm_ptr.Elem()))
	}

}

/**    返回一行数据，一般是返回一个结构体    */
func (this *Pdo) SelectOne(sqlstring string, bindarray []interface{}) map[string]string {
	rows, scanArgs, values, columns := this.query(sqlstring, bindarray)
	defer rows.Close()
	// 这个map用来存储一行数据，列名为map的key，map的value为列的值
	var rowMap = make(map[string]string)
	for rows.Next() {
		rows.Scan(scanArgs...)
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col != nil {
				rowMap[columns[i]] = string(col)
			}
		}
		break
	}
	return rowMap
}

/**    查询一行数据返回一个结构体    */
func (this *Pdo) SelectOneObject(sql string, bindarray []interface{}, orm_ptr interface{}) {
	one := this.SelectOne(sql, bindarray)
	Map_to_struct(one, orm_ptr)
}

/**
正在执行sql的部分代码
*/
func (this *Pdo) pdoexec(sql string, bindarray []interface{}) sql.Result {
	// 如果数据库还没链接，那么初始化一下
	if this.tx == nil {
		this.Begin()
	}
	// defer this.tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.
	stmt, err := this.tx.Prepare(sql)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		this.Rollback()
		panic(fmt.Sprintf("\033[41;36merr:%+v %+v:%+v \033[0m\n", []interface{}{err}, file, line))
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	Result, err := stmt.Exec(bindarray...)
	if err != nil {
		this.Rollback()
		_, file, line, _ := runtime.Caller(0)
		panic(fmt.Sprintf("\033[41;36merr:%+v %+v:%+v \033[0m\n", []interface{}{err}, file, line))
	}
	return Result
}

/**
  执行指定的SQL语句;
*/
func (this *Pdo) Exec(sql string, bindarray []interface{}) int {
	Result := this.pdoexec(sql, bindarray)
	num, _ := Result.RowsAffected()
	return int(num)
}

/**
  写入数据;
*/
func (this *Pdo) Insert(sql string, bindarray []interface{}) int {
	Result := this.pdoexec(sql, bindarray)
	num, _ := Result.LastInsertId()
	return int(num)
}
