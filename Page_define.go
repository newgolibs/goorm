package goorm
import (
    "time"
)

//对象必须实现的接口方法
type PageInterface interface {
    /**    配合mysql，在当前数目下，limit 语句的拼接    */
    SqlLImit()string

}

//定义函数的结构体，方便扩展成中间件接
type Page_SqlLImitHandleFunc func()string

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







func (this *Page) NewPageMiddleware() *PageMiddleware{
    return &PageMiddleware{Page:this}
}
/**
* 中间件的扩展类middleware
*/
type PageMiddleware struct{
    SqlLImitindex int
    SqlLImitHandleFuncs []Page_SqlLImitHandleFunc
    Page *Page
    //日志记录的目标文件
    SQLLogger Logger
}


func (this *PageMiddleware) Add_SqlLImit(middlewares ...Page_SqlLImitHandleFunc) Page_SqlLImitHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.SqlLImitHandleFuncs) == 0 {
        this.SqlLImitHandleFuncs = append(this.SqlLImitHandleFuncs, func() string {
            defer func(start time.Time) {
                if this.SQLLogger != nil {
                    tc := time.Since(start).String()
                    this.SQLLogger.Debug("耗时 - Page.SqlLImit",tc)
                }
            }(time.Now())
            if this.SQLLogger != nil {
                this.SQLLogger.Debug("调起 - Page.SqlLImit，参数： ",)
            }
            return this.Next_CALL_SqlLImit()
        })
    }

    //
	if this.SqlLImitHandleFuncs == nil {
		this.SqlLImitHandleFuncs = make([]Page_SqlLImitHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.SqlLImitHandleFuncs = append(this.SqlLImitHandleFuncs, mid)
	}
	return this.Next_CALL_SqlLImit
}
func (this *PageMiddleware) Next_SqlLImit()string {
    this.SqlLImitindex = 0
    return this.Next_CALL_SqlLImit()
}
func (this *PageMiddleware) Next_CALL_SqlLImit()string{
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.SqlLImitHandleFuncs) == 0 {
		this.Add_SqlLImit(this.Page.SqlLImit)
	} else if this.SqlLImitindex == 0 {
		this.SqlLImitHandleFuncs = append(this.SqlLImitHandleFuncs, this.Page.SqlLImit)
	}
    index := this.SqlLImitindex
	if this.SqlLImitindex >= len(this.SqlLImitHandleFuncs) {
		return ""	}

	this.SqlLImitindex++
	return this.SqlLImitHandleFuncs[index]()
}

//检测接口是否被完整的实现了，如果没有实现，那么编译不通过
var _ PageInterface =new(Page)
