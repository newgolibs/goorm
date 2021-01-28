package goorm
import (
        "database/sql"
    "time"
)

//对象必须实现的接口方法
type PdoconfigInterface interface {
    /**    链接数据库，拼接起来的字符串    */
    LinkString()string
    /**    链接池    */
    MakeDbPool()*Pdoconfig
    /**    新开事务线程    */
    MakeTX()*sql.Tx
    /**    生成新的pdo对象    */
    NewPdo()*Pdo
    /**    返回命令行下的连接字符串    */
    ShellLinkString()string

}

//定义函数的结构体，方便扩展成中间件接
type Pdoconfig_LinkStringHandleFunc func()string
type Pdoconfig_MakeDbPoolHandleFunc func()*Pdoconfig
type Pdoconfig_MakeTXHandleFunc func()*sql.Tx
type Pdoconfig_NewPdoHandleFunc func()*Pdo
type Pdoconfig_ShellLinkStringHandleFunc func()string

/**
数据库配置;
*/
type Pdoconfig struct
{
    /*服务器地址*/
    Tns string
    /*账户*/
    User string
    /*密码*/
    Password string
    /*数据库*/
    DB string
    /*端口*/
    Port int
    /*数据库链接句柄*/
    Sqldb *sql.DB
}










func (this *Pdoconfig) NewPdoconfigMiddleware() *PdoconfigMiddleware{
    return &PdoconfigMiddleware{Pdoconfig:this}
}
/**
* 中间件的扩展类middleware
*/
type PdoconfigMiddleware struct{
    LinkStringindex int
    LinkStringHandleFuncs []Pdoconfig_LinkStringHandleFunc
    MakeDbPoolindex int
    MakeDbPoolHandleFuncs []Pdoconfig_MakeDbPoolHandleFunc
    MakeTXindex int
    MakeTXHandleFuncs []Pdoconfig_MakeTXHandleFunc
    NewPdoindex int
    NewPdoHandleFuncs []Pdoconfig_NewPdoHandleFunc
    ShellLinkStringindex int
    ShellLinkStringHandleFuncs []Pdoconfig_ShellLinkStringHandleFunc
    Pdoconfig *Pdoconfig
    //日志记录的目标文件
    SQLLogger Logger
}


func (this *PdoconfigMiddleware) Add_LinkString(middlewares ...Pdoconfig_LinkStringHandleFunc) Pdoconfig_LinkStringHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.LinkStringHandleFuncs) == 0 {
        this.LinkStringHandleFuncs = append(this.LinkStringHandleFuncs, func() string {
            defer func(start time.Time) {
                if this.SQLLogger != nil {
                    tc := time.Since(start).String()
                    this.SQLLogger.Debug("耗时 - Pdoconfig.LinkString:%+v",tc)
                }
            }(time.Now())
            if this.SQLLogger != nil {
                this.SQLLogger.Debug("调起 - Pdoconfig.LinkString，参数：%#v ",[]interface{}{})
            }
            return this.Next_CALL_LinkString()
        })
    }

    //
	if this.LinkStringHandleFuncs == nil {
		this.LinkStringHandleFuncs = make([]Pdoconfig_LinkStringHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.LinkStringHandleFuncs = append(this.LinkStringHandleFuncs, mid)
	}
	return this.Next_CALL_LinkString
}
/**
* 中间件，替代函数入口
*/
func (this *PdoconfigMiddleware) LinkString()string {
    this.LinkStringindex = 0
    return this.Next_CALL_LinkString()
}

/**
*/
func (this *PdoconfigMiddleware) Next_CALL_LinkString()string{
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.LinkStringHandleFuncs) == 0 {
		this.Add_LinkString(this.Pdoconfig.LinkString)
	} else if this.LinkStringindex == 0 {
        // 👇👇---- 原始函数入口
		this.LinkStringHandleFuncs = append(this.LinkStringHandleFuncs, this.Pdoconfig.LinkString)
	}
    index := this.LinkStringindex
	if this.LinkStringindex >= len(this.LinkStringHandleFuncs) {
		return ""	}

	this.LinkStringindex++
	return this.LinkStringHandleFuncs[index]()
}

func (this *PdoconfigMiddleware) Add_MakeDbPool(middlewares ...Pdoconfig_MakeDbPoolHandleFunc) Pdoconfig_MakeDbPoolHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.MakeDbPoolHandleFuncs) == 0 {
        this.MakeDbPoolHandleFuncs = append(this.MakeDbPoolHandleFuncs, func() *Pdoconfig {
            defer func(start time.Time) {
                if this.SQLLogger != nil {
                    tc := time.Since(start).String()
                    this.SQLLogger.Debug("耗时 - Pdoconfig.MakeDbPool:%+v",tc)
                }
            }(time.Now())
            if this.SQLLogger != nil {
                this.SQLLogger.Debug("调起 - Pdoconfig.MakeDbPool，参数：%#v ",[]interface{}{})
            }
            return this.Next_CALL_MakeDbPool()
        })
    }

    //
	if this.MakeDbPoolHandleFuncs == nil {
		this.MakeDbPoolHandleFuncs = make([]Pdoconfig_MakeDbPoolHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.MakeDbPoolHandleFuncs = append(this.MakeDbPoolHandleFuncs, mid)
	}
	return this.Next_CALL_MakeDbPool
}
/**
* 中间件，替代函数入口
*/
func (this *PdoconfigMiddleware) MakeDbPool()*Pdoconfig {
    this.MakeDbPoolindex = 0
    return this.Next_CALL_MakeDbPool()
}

/**
*/
func (this *PdoconfigMiddleware) Next_CALL_MakeDbPool()*Pdoconfig{
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.MakeDbPoolHandleFuncs) == 0 {
		this.Add_MakeDbPool(this.Pdoconfig.MakeDbPool)
	} else if this.MakeDbPoolindex == 0 {
        // 👇👇---- 原始函数入口
		this.MakeDbPoolHandleFuncs = append(this.MakeDbPoolHandleFuncs, this.Pdoconfig.MakeDbPool)
	}
    index := this.MakeDbPoolindex
	if this.MakeDbPoolindex >= len(this.MakeDbPoolHandleFuncs) {
		return nil	}

	this.MakeDbPoolindex++
	return this.MakeDbPoolHandleFuncs[index]()
}

func (this *PdoconfigMiddleware) Add_MakeTX(middlewares ...Pdoconfig_MakeTXHandleFunc) Pdoconfig_MakeTXHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.MakeTXHandleFuncs) == 0 {
        this.MakeTXHandleFuncs = append(this.MakeTXHandleFuncs, func() *sql.Tx {
            defer func(start time.Time) {
                if this.SQLLogger != nil {
                    tc := time.Since(start).String()
                    this.SQLLogger.Debug("耗时 - Pdoconfig.MakeTX:%+v",tc)
                }
            }(time.Now())
            if this.SQLLogger != nil {
                this.SQLLogger.Debug("调起 - Pdoconfig.MakeTX，参数：%#v ",[]interface{}{})
            }
            return this.Next_CALL_MakeTX()
        })
    }

    //
	if this.MakeTXHandleFuncs == nil {
		this.MakeTXHandleFuncs = make([]Pdoconfig_MakeTXHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.MakeTXHandleFuncs = append(this.MakeTXHandleFuncs, mid)
	}
	return this.Next_CALL_MakeTX
}
/**
* 中间件，替代函数入口
*/
func (this *PdoconfigMiddleware) MakeTX()*sql.Tx {
    this.MakeTXindex = 0
    return this.Next_CALL_MakeTX()
}

/**
*/
func (this *PdoconfigMiddleware) Next_CALL_MakeTX()*sql.Tx{
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.MakeTXHandleFuncs) == 0 {
		this.Add_MakeTX(this.Pdoconfig.MakeTX)
	} else if this.MakeTXindex == 0 {
        // 👇👇---- 原始函数入口
		this.MakeTXHandleFuncs = append(this.MakeTXHandleFuncs, this.Pdoconfig.MakeTX)
	}
    index := this.MakeTXindex
	if this.MakeTXindex >= len(this.MakeTXHandleFuncs) {
		return nil	}

	this.MakeTXindex++
	return this.MakeTXHandleFuncs[index]()
}

func (this *PdoconfigMiddleware) Add_NewPdo(middlewares ...Pdoconfig_NewPdoHandleFunc) Pdoconfig_NewPdoHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.NewPdoHandleFuncs) == 0 {
        this.NewPdoHandleFuncs = append(this.NewPdoHandleFuncs, func() *Pdo {
            defer func(start time.Time) {
                if this.SQLLogger != nil {
                    tc := time.Since(start).String()
                    this.SQLLogger.Debug("耗时 - Pdoconfig.NewPdo:%+v",tc)
                }
            }(time.Now())
            if this.SQLLogger != nil {
                this.SQLLogger.Debug("调起 - Pdoconfig.NewPdo，参数：%#v ",[]interface{}{})
            }
            return this.Next_CALL_NewPdo()
        })
    }

    //
	if this.NewPdoHandleFuncs == nil {
		this.NewPdoHandleFuncs = make([]Pdoconfig_NewPdoHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.NewPdoHandleFuncs = append(this.NewPdoHandleFuncs, mid)
	}
	return this.Next_CALL_NewPdo
}
/**
* 中间件，替代函数入口
*/
func (this *PdoconfigMiddleware) NewPdo()*Pdo {
    this.NewPdoindex = 0
    return this.Next_CALL_NewPdo()
}

/**
*/
func (this *PdoconfigMiddleware) Next_CALL_NewPdo()*Pdo{
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.NewPdoHandleFuncs) == 0 {
		this.Add_NewPdo(this.Pdoconfig.NewPdo)
	} else if this.NewPdoindex == 0 {
        // 👇👇---- 原始函数入口
		this.NewPdoHandleFuncs = append(this.NewPdoHandleFuncs, this.Pdoconfig.NewPdo)
	}
    index := this.NewPdoindex
	if this.NewPdoindex >= len(this.NewPdoHandleFuncs) {
		return nil	}

	this.NewPdoindex++
	return this.NewPdoHandleFuncs[index]()
}

func (this *PdoconfigMiddleware) Add_ShellLinkString(middlewares ...Pdoconfig_ShellLinkStringHandleFunc) Pdoconfig_ShellLinkStringHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.ShellLinkStringHandleFuncs) == 0 {
        this.ShellLinkStringHandleFuncs = append(this.ShellLinkStringHandleFuncs, func() string {
            defer func(start time.Time) {
                if this.SQLLogger != nil {
                    tc := time.Since(start).String()
                    this.SQLLogger.Debug("耗时 - Pdoconfig.ShellLinkString:%+v",tc)
                }
            }(time.Now())
            if this.SQLLogger != nil {
                this.SQLLogger.Debug("调起 - Pdoconfig.ShellLinkString，参数：%#v ",[]interface{}{})
            }
            return this.Next_CALL_ShellLinkString()
        })
    }

    //
	if this.ShellLinkStringHandleFuncs == nil {
		this.ShellLinkStringHandleFuncs = make([]Pdoconfig_ShellLinkStringHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.ShellLinkStringHandleFuncs = append(this.ShellLinkStringHandleFuncs, mid)
	}
	return this.Next_CALL_ShellLinkString
}
/**
* 中间件，替代函数入口
*/
func (this *PdoconfigMiddleware) ShellLinkString()string {
    this.ShellLinkStringindex = 0
    return this.Next_CALL_ShellLinkString()
}

/**
*/
func (this *PdoconfigMiddleware) Next_CALL_ShellLinkString()string{
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.ShellLinkStringHandleFuncs) == 0 {
		this.Add_ShellLinkString(this.Pdoconfig.ShellLinkString)
	} else if this.ShellLinkStringindex == 0 {
        // 👇👇---- 原始函数入口
		this.ShellLinkStringHandleFuncs = append(this.ShellLinkStringHandleFuncs, this.Pdoconfig.ShellLinkString)
	}
    index := this.ShellLinkStringindex
	if this.ShellLinkStringindex >= len(this.ShellLinkStringHandleFuncs) {
		return ""	}

	this.ShellLinkStringindex++
	return this.ShellLinkStringHandleFuncs[index]()
}

//检测接口是否被完整的实现了，如果没有实现，那么编译不通过
var _ PdoconfigInterface =new(Pdoconfig)
