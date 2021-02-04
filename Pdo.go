package goorm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"reflect"
)

/**    当运行中，有一条sql错误了，那么回滚，在这个事务期间的所有操作全部报废    */
func (this *Pdo) Rollback() {
	var err error
	err = this.TX.Rollback()
	if err != nil {
		panic("rollback error:" + err.Error())
	}
}

/**    提交事务，并且还继续开启事务    */
func (this *Pdo) Commit_NewTX() {
	this.Commit(nil)
	this.TX = this.Pdoconfig.MakeTX()
}

/**
提交事务
*/
func (this *Pdo) Commit(recover interface{}) {
	var err error
	if recover != nil { // 👈👈---- 发现有错误了
		this.Rollback()
	} else { // 👈👈----  没有错误，提交
		err = this.TX.Commit()
		if err != nil {
			panic("commit error:" + err.Error())
		}
	}
}

/**    执行Query方法，返回rows    */
func (this *Pdo) query(sqlstring string, bindarray []interface{}) (*sql.Rows, []interface{}, []sql.RawBytes, []string) {
	rows, err := this.TX.Query(sqlstring, bindarray...)
	if err != nil {
		panic("query error:" + err.Error())
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
func (this *Pdo) SelectAll(sqlstring string, bindarray []interface{}) ([]map[string]string, error) {
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
	return rowMaps_All, nil
}

/**    查询多行数据，返回struct对象的数组    */
func (this *Pdo) SelectallObject(sql string, bindarray []interface{}, orm_ptr interface{}) error {
	orm_ptr_ref := reflect.ValueOf(orm_ptr)
	orm_ptr_value := orm_ptr_ref.Elem()
	if orm_ptr_value.Kind() != reflect.Slice {
		panic("传入的类型应该是一个切片：[]xxxx，现在是：" + orm_ptr_value.Kind().String())
	}
	all, err := this.SelectAll(sql, bindarray)
	if err != nil {
		return err
	}
	orm_ptr_type := reflect.TypeOf(orm_ptr)
	for _, item := range all {
		one_orm_ptr := reflect.New(orm_ptr_type.Elem().Elem())
		Map_to_struct(item, one_orm_ptr.Interface())
		orm_ptr_value.Set(reflect.Append(orm_ptr_value, one_orm_ptr.Elem()))
	}
	return nil
}

/**    返回一行数据，一般是返回一个结构体    */
func (this *Pdo) SelectOne(sql string, bindarray []interface{}) (map[string]string, error) {
	rows, scanArgs, values, columns := this.query(sql, bindarray)
	defer rows.Close()
	// 这个map用来存储一行数据，列名为map的key，map的value为列的值
	var rowMap = make(map[string]string)
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			panic("rows.Scan error:" + err.Error())
		}
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col != nil {
				rowMap[columns[i]] = string(col)
			}
		}
		break
	}
	return rowMap, nil
}

/**    查询一行数据返回一个结构体    */
func (this *Pdo) SelectOneObject(sql string, bindarray []interface{}, orm_ptr interface{}) error {
	one, err := this.SelectOne(sql, bindarray)
	Map_to_struct(one, orm_ptr)
	return err
}

/**    查询一行数据返回一个结构体    */
func (this *Pdo) SelectVar(sql string, bindarray []interface{}) (string, error) {
	one, err := this.SelectOne(sql, bindarray)
	if err != nil {
		return "0", err
	}
	for _, v := range one {
		return v, nil
	}
	return "0", errors.New("not found")
}

/**
正在执行sql的部分代码
*/
func (this *Pdo) pdoexec(sql string, bindarray []interface{}) sql.Result {
	stmt, err := this.TX.Prepare(sql)
	if err != nil {
		panic("pdoexec error:" + err.Error())
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	Result, err := stmt.Exec(bindarray...)
	if err != nil {
		panic("stmt.Exec error:" + err.Error())
	}
	return Result
}

/**
  执行指定的SQL语句;
*/
func (this *Pdo) Exec(sql string, bindarray []interface{}) (int64, error) {
	Result := this.pdoexec(sql, bindarray)
	num, err := Result.RowsAffected()
	return num, err
}

/**
  写入数据;
*/
func (this *Pdo) Insert(sql string, bindarray []interface{}) (int64, error) {
	Result := this.pdoexec(sql, bindarray)
	num, err := Result.LastInsertId()
	return num, err
}

