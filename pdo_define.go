package goorm

type pdoInterface interface {
    /**    执行指定的SQL语句    */
    Exec()
    /**    写入数据    */
    Insert()

}

/**
数据库执行，返回数据;
*/
type pdo struct
{
    /*数据库链接实例*/
    _db string
    /*绑定的值*/
    Bindarray []string
    /*数据库配置*/
    Pdoconfig pdoconfig
    /*要执行的sql*/
    Sql string
}










//检测接口是否被完整的实现了，如果没有实现，那么编译不通过
var _ pdoInterface =new(pdo)
