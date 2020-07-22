package goorm

import (
	"database/sql"
	"fmt"
	"runtime"
)

/**
开启事务
*/
func (this *Pdo) Begin() {
	var err error
	this.Tx, err = this.Pdoconfig.SqldbPool().Begin()
	if err != nil {
		fmt.Printf("err:%+v\n", err)
	}
}

/**
提交事务
*/
func (this *Pdo) Commit() {
	var err error
	err = this.Tx.Commit()
	if err != nil {
		fmt.Printf("err:%+v\n", err)
	}
}

/**    执行Query方法，返回rows    */
func (this *Pdo) query(sqlstring string, bindarray []interface{}) (*sql.Rows, []interface{}, []sql.RawBytes, []string) {
	rows, err := this.Tx.Query(sqlstring, bindarray...)
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
	var rowMaps []map[string]string
	rowMap := make(map[string]string)
	for rows.Next() {
		rows.Scan(scanArgs...)
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col != nil {
				rowMap[columns[i]] = string(col)
			}
			rowMaps = append(rowMaps, rowMap)
		}
	}
	return rowMaps
}

/**    返回一行数据，一般是返回一个结构体    */
func (this *Pdo) SelectOne(sqlstring string, bindarray []interface{}) map[string]string {
	rows, scanArgs, values, columns := this.query(sqlstring, bindarray)
	defer rows.Close()
	// 这个map用来存储一行数据，列名为map的key，map的value为列的值
	rowMap := make(map[string]string)
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

/**
正在执行sql的部分代码
*/
func (this *Pdo) pdoexec(sql string, bindarray []interface{}) sql.Result {
	// defer this.Tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.
	stmt, err := this.Tx.Prepare(sql)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		panic(fmt.Sprintf("\033[41;36merr:%+v %+v:%+v \033[0m\n", []interface{}{err}, file, line))
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	Result, err := stmt.Exec(bindarray...)
	if err != nil {
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
