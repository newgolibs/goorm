package goorm

type PdoInterface interface {
    /**    执行指定的SQL语句1    */
    Exec(sql string,bindarray []interface{})int
    /**    写入数据    */
    Insert()

}

/**
数据库执行，返回数据;
*/
type Pdo struct
{
    /*数据库配置*/
    Pdoconfig Pdoconfig
}







//检测接口是否被完整的实现了，如果没有实现，那么编译不通过
var _ PdoInterface =new(Pdo)
