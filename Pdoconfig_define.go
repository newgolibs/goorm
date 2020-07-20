package goorm
import (
        "database/sql"
)

type PdoconfigInterface interface {
    /**    链接数据库，拼接起来的字符串    */
    LinkString()string
    /**    链接池    */
    SqldbPool()*sql.DB

}

/**
数据库配置;
*/
type Pdoconfig struct
{
    /*数据库*/
    DB string
    /*密码*/
    Password string
    /*端口*/
    Port int
    /*服务器地址*/
    Tns string
    /*账户*/
    User string
}











//检测接口是否被完整的实现了，如果没有实现，那么编译不通过
var _ PdoconfigInterface =new(Pdoconfig)
