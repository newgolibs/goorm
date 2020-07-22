package goorm
import (
        "database/sql"
)

type PdoInterface interface {
    /**    开启事务    */
    Begin()
    /**    提交事务    */
    Commit()
    /**    执行指定的SQL语句，返回影响到的条数    */
    Exec(sql string,bindarray []interface{})int
    /**    写入数据，返回最新写入数据的自增id    */
    Insert(sql string,bindarray []interface{})int
    /**    正在执行sql的部分代码，更新删除写入之类的操作    */
    pdoexec(sql string,bindarray []interface{})sql.Result
    /**    执行Query方法，返回rows    */
    query(sql string,bindarray []interface{})(*sql.Rows,[]interface{},[]sql.RawBytes,[]string)
    /**    查询多行数据    */
    SelectAll(sql string,bindarray []interface{})[]map[string]string
    /**    返回一行数据，一般是返回一个结构体    */
    SelectOne(sql string,bindarray []interface{})map[string]string

}

/**
数据库执行，返回数据;
*/
type Pdo struct
{
    /*数据库配置*/
    Pdoconfig Pdoconfig
    /*事务链接句柄*/
    Tx *sql.Tx
}








//检测接口是否被完整的实现了，如果没有实现，那么编译不通过
var _ PdoInterface =new(Pdo)
