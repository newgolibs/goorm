package goorm
import (
        "github.com/rs/zerolog"
    "time"
)

//对象必须实现的接口方法
type PageInterface interface {
    /**    配合mysql，在当前数目下，limit 语句的拼接    */
    SqlLImit()string
    /**    设置总条数，且计算总页数    */
    SetTotal(Total int)

}

//定义函数的结构体，方便扩展成中间件接
type Page_SqlLImitHandleFunc func()string
type Page_SetTotalHandleFunc func(Total int)

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
    /*总页数*/
    TotalPages int
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
    SetTotalindex int
    SetTotalHandleFuncs []Page_SetTotalHandleFunc
    Page *Page
    //日志记录的目标文件
    zloger *zerolog.Logger
}

/**
设置日志写入类。
*/
func (this *PageMiddleware) SetZloger(l *zerolog.Logger) *PageMiddleware {
    Zloger := l.With().Str("library","goorm").
                    Str("class","PageMiddleware").
                    Logger()
    this.zloger = &Zloger
    return this
}


func (this *PageMiddleware) Add_SqlLImit(middlewares ...Page_SqlLImitHandleFunc) Page_SqlLImitHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.SqlLImitHandleFuncs) == 0 {
        this.SqlLImitHandleFuncs = append(this.SqlLImitHandleFuncs, func() string {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").Logger()
                    zloger.Debug().Msgf("耗时 - Page.SqlLImit:%+v",tc)
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").Logger()
                zloger.Debug().Msgf("调起 - Page.SqlLImit，参数：%#v ",[]interface{}{})
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
/**
* 中间件，替代函数入口
*/
func (this *PageMiddleware) SqlLImit()string {
    this.SqlLImitindex = 0
    return this.Next_CALL_SqlLImit()
}

/**
*/
func (this *PageMiddleware) Next_CALL_SqlLImit()string{
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.SqlLImitHandleFuncs) == 0 {
		this.Add_SqlLImit(this.Page.SqlLImit)
	} else if this.SqlLImitindex == 0 {
        // 👇👇---- 原始函数入口
		this.SqlLImitHandleFuncs = append(this.SqlLImitHandleFuncs, this.Page.SqlLImit)
	}
    index := this.SqlLImitindex
	if this.SqlLImitindex >= len(this.SqlLImitHandleFuncs) {
		return ""	}

	this.SqlLImitindex++
	return this.SqlLImitHandleFuncs[index]()
}

func (this *PageMiddleware) Add_SetTotal(middlewares ...Page_SetTotalHandleFunc) Page_SetTotalHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.SetTotalHandleFuncs) == 0 {
        this.SetTotalHandleFuncs = append(this.SetTotalHandleFuncs, func(Total int)  {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").Logger()
                    zloger.Debug().Msgf("耗时 - Page.SetTotal:%+v",tc)
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").Logger()
                zloger.Debug().Msgf("调起 - Page.SetTotal，参数：%#v ",[]interface{}{Total})
            }
            this.Next_CALL_SetTotal(Total)
        })
    }

    //
	if this.SetTotalHandleFuncs == nil {
		this.SetTotalHandleFuncs = make([]Page_SetTotalHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.SetTotalHandleFuncs = append(this.SetTotalHandleFuncs, mid)
	}
	return this.Next_CALL_SetTotal
}
/**
* 中间件，替代函数入口
*/
func (this *PageMiddleware) SetTotal(Total int) {
    this.SetTotalindex = 0
    this.Next_CALL_SetTotal(Total)
}

/**
*/
func (this *PageMiddleware) Next_CALL_SetTotal(Total int){
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.SetTotalHandleFuncs) == 0 {
		this.Add_SetTotal(this.Page.SetTotal)
	} else if this.SetTotalindex == 0 {
        // 👇👇---- 原始函数入口
		this.SetTotalHandleFuncs = append(this.SetTotalHandleFuncs, this.Page.SetTotal)
	}
    index := this.SetTotalindex
	if this.SetTotalindex >= len(this.SetTotalHandleFuncs) {
        return
	}

	this.SetTotalindex++
    this.SetTotalHandleFuncs[index](Total)
}

//检测接口是否被完整的实现了，如果没有实现，那么编译不通过
var _ PageInterface =new(Page)
