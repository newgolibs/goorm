package goorm

//对象必须实现的接口方法
type PageInterface interface {
    /**    配合mysql，在当前数目下，limit 语句的拼接    */
    SqlLImit()string

}

/**
结合数据库分页;
*/
type Page struct
{
    /**/
    PageID int
    /*总数*/
    Total int
    /*每页条目*/
    Prepage int
}









//检测接口是否被完整的实现了，如果没有实现，那么编译不通过
var _ PageInterface =new(Page)
