package goorm
import (
        "database/sql"
        "github.com/rs/zerolog"
    "time"
)

//对象必须实现的接口方法
type DBInterface interface {
    /**    写入数据，返回最新写入数据的自增id    */
    Insert(sql string,bindarray []interface{})(int64,error)
    /**    执行指定的SQL语句，返回影响到的条数    */
    Exec(sql string,bindarray []interface{})(int64,error)
    /**    正在执行sql的部分代码，更新删除写入之类的操作    */
    pdoexec(sql string,bindarray []interface{})sql.Result
    /**    返回一行数据，map类型    */
    SelectOne(sql string,bindarray []interface{})(map[string]string,error)
    /**    查询一行数据返回一个结构体    */
    SelectOneObject(sql string,bindarray []interface{},orm_ptr interface{})error
    /**    查询多行数据，返回map类型    */
    SelectAll(sql string,bindarray []interface{})([]map[string]string,error)
    /**    查询多行数据，返回struct对象的数组    */
    SelectallObject(sql string,bindarray []interface{},orm_ptr interface{})error
    /**    执行Query方法，返回rows    */
    query(sql string,bindarray []interface{})(*sql.Rows,[]interface{},[]sql.RawBytes,[]string)
    /**    查询count    */
    SelectVar(sql string,bindarray []interface{})(string,error)

}

//定义函数的结构体，方便扩展成中间件接
type DB_InsertHandleFunc func(sql string,bindarray []interface{})(int64,error)
type DB_ExecHandleFunc func(sql string,bindarray []interface{})(int64,error)
type DB_SelectOneHandleFunc func(sql string,bindarray []interface{})(map[string]string,error)
type DB_SelectOneObjectHandleFunc func(sql string,bindarray []interface{},orm_ptr interface{})error
type DB_SelectAllHandleFunc func(sql string,bindarray []interface{})([]map[string]string,error)
type DB_SelectallObjectHandleFunc func(sql string,bindarray []interface{},orm_ptr interface{})error
type DB_SelectVarHandleFunc func(sql string,bindarray []interface{})(string,error)

/**
不带事务的数据库操作;
*/
type DB struct
{
    /*数据库配置，中间件*/
    Pdoconfig *PdoconfigMiddleware
    /*非事务链接句柄，db*/
    TX *sql.DB
}






func (this *DB) NewDBMiddleware() *DBMiddleware{
    return &DBMiddleware{DB:this}
}
/**
* 中间件的扩展类middleware
*/
type DBMiddleware struct{
    Insertindex int
    InsertHandleFuncs []DB_InsertHandleFunc
    Execindex int
    ExecHandleFuncs []DB_ExecHandleFunc
    SelectOneindex int
    SelectOneHandleFuncs []DB_SelectOneHandleFunc
    SelectOneObjectindex int
    SelectOneObjectHandleFuncs []DB_SelectOneObjectHandleFunc
    SelectAllindex int
    SelectAllHandleFuncs []DB_SelectAllHandleFunc
    SelectallObjectindex int
    SelectallObjectHandleFuncs []DB_SelectallObjectHandleFunc
    SelectVarindex int
    SelectVarHandleFuncs []DB_SelectVarHandleFunc
    DB *DB
    //日志记录的目标文件
    zloger *zerolog.Logger
}

/**
设置日志写入类。
*/
func (this *DBMiddleware) SetZloger(l *zerolog.Logger) *DBMiddleware {
    Zloger := l.With().Str("library","goorm").
                    Str("class","DBMiddleware").
                    Logger()
    this.zloger = &Zloger
    return this
}


func (this *DBMiddleware) Add_Insert(middlewares ...DB_InsertHandleFunc) DB_InsertHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.InsertHandleFuncs) == 0 {
        this.InsertHandleFuncs = append(this.InsertHandleFuncs, func(sql string,bindarray []interface{}) (int64,error) {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").
                                Str("func","Insert").Str("timeuse",tc).Logger()
                    zloger.Debug().Msg("耗时")
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").
                            Interface("call_args",[]interface{}{sql,bindarray}).
                            Str("func","Insert").Logger()
                zloger.Debug().Msg("调起")
            }
            return this.Next_CALL_Insert(sql,bindarray)
        })
    }

    //
	if this.InsertHandleFuncs == nil {
		this.InsertHandleFuncs = make([]DB_InsertHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.InsertHandleFuncs = append(this.InsertHandleFuncs, mid)
	}
	return this.Next_CALL_Insert
}
/**
* 中间件，替代函数入口
*/
func (this *DBMiddleware) Insert(sql string,bindarray []interface{})(int64,error) {
    this.Insertindex = 0
    return this.Next_CALL_Insert(sql,bindarray)
}

/**
*/
func (this *DBMiddleware) Next_CALL_Insert(sql string,bindarray []interface{})(int64,error){
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.InsertHandleFuncs) == 0 {
		this.Add_Insert(this.DB.Insert)
	} else if this.Insertindex == 0 {
        // 👇👇---- 原始函数入口
		this.InsertHandleFuncs = append(this.InsertHandleFuncs, this.DB.Insert)
	}
    index := this.Insertindex
	if this.Insertindex >= len(this.InsertHandleFuncs) {
		return 0,nil	}

	this.Insertindex++
	return this.InsertHandleFuncs[index](sql,bindarray)
}

func (this *DBMiddleware) Add_Exec(middlewares ...DB_ExecHandleFunc) DB_ExecHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.ExecHandleFuncs) == 0 {
        this.ExecHandleFuncs = append(this.ExecHandleFuncs, func(sql string,bindarray []interface{}) (int64,error) {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").
                                Str("func","Exec").Str("timeuse",tc).Logger()
                    zloger.Debug().Msg("耗时")
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").
                            Interface("call_args",[]interface{}{sql,bindarray}).
                            Str("func","Exec").Logger()
                zloger.Debug().Msg("调起")
            }
            return this.Next_CALL_Exec(sql,bindarray)
        })
    }

    //
	if this.ExecHandleFuncs == nil {
		this.ExecHandleFuncs = make([]DB_ExecHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.ExecHandleFuncs = append(this.ExecHandleFuncs, mid)
	}
	return this.Next_CALL_Exec
}
/**
* 中间件，替代函数入口
*/
func (this *DBMiddleware) Exec(sql string,bindarray []interface{})(int64,error) {
    this.Execindex = 0
    return this.Next_CALL_Exec(sql,bindarray)
}

/**
*/
func (this *DBMiddleware) Next_CALL_Exec(sql string,bindarray []interface{})(int64,error){
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.ExecHandleFuncs) == 0 {
		this.Add_Exec(this.DB.Exec)
	} else if this.Execindex == 0 {
        // 👇👇---- 原始函数入口
		this.ExecHandleFuncs = append(this.ExecHandleFuncs, this.DB.Exec)
	}
    index := this.Execindex
	if this.Execindex >= len(this.ExecHandleFuncs) {
		return 0,nil	}

	this.Execindex++
	return this.ExecHandleFuncs[index](sql,bindarray)
}

func (this *DBMiddleware) Add_SelectOne(middlewares ...DB_SelectOneHandleFunc) DB_SelectOneHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.SelectOneHandleFuncs) == 0 {
        this.SelectOneHandleFuncs = append(this.SelectOneHandleFuncs, func(sql string,bindarray []interface{}) (map[string]string,error) {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").
                                Str("func","SelectOne").Str("timeuse",tc).Logger()
                    zloger.Debug().Msg("耗时")
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").
                            Interface("call_args",[]interface{}{sql,bindarray}).
                            Str("func","SelectOne").Logger()
                zloger.Debug().Msg("调起")
            }
            return this.Next_CALL_SelectOne(sql,bindarray)
        })
    }

    //
	if this.SelectOneHandleFuncs == nil {
		this.SelectOneHandleFuncs = make([]DB_SelectOneHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.SelectOneHandleFuncs = append(this.SelectOneHandleFuncs, mid)
	}
	return this.Next_CALL_SelectOne
}
/**
* 中间件，替代函数入口
*/
func (this *DBMiddleware) SelectOne(sql string,bindarray []interface{})(map[string]string,error) {
    this.SelectOneindex = 0
    return this.Next_CALL_SelectOne(sql,bindarray)
}

/**
*/
func (this *DBMiddleware) Next_CALL_SelectOne(sql string,bindarray []interface{})(map[string]string,error){
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.SelectOneHandleFuncs) == 0 {
		this.Add_SelectOne(this.DB.SelectOne)
	} else if this.SelectOneindex == 0 {
        // 👇👇---- 原始函数入口
		this.SelectOneHandleFuncs = append(this.SelectOneHandleFuncs, this.DB.SelectOne)
	}
    index := this.SelectOneindex
	if this.SelectOneindex >= len(this.SelectOneHandleFuncs) {
		return nil,nil	}

	this.SelectOneindex++
	return this.SelectOneHandleFuncs[index](sql,bindarray)
}

func (this *DBMiddleware) Add_SelectOneObject(middlewares ...DB_SelectOneObjectHandleFunc) DB_SelectOneObjectHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.SelectOneObjectHandleFuncs) == 0 {
        this.SelectOneObjectHandleFuncs = append(this.SelectOneObjectHandleFuncs, func(sql string,bindarray []interface{},orm_ptr interface{}) error {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").
                                Str("func","SelectOneObject").Str("timeuse",tc).Logger()
                    zloger.Debug().Msg("耗时")
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").
                            Interface("call_args",[]interface{}{sql,bindarray,orm_ptr}).
                            Str("func","SelectOneObject").Logger()
                zloger.Debug().Msg("调起")
            }
            return this.Next_CALL_SelectOneObject(sql,bindarray,orm_ptr)
        })
    }

    //
	if this.SelectOneObjectHandleFuncs == nil {
		this.SelectOneObjectHandleFuncs = make([]DB_SelectOneObjectHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.SelectOneObjectHandleFuncs = append(this.SelectOneObjectHandleFuncs, mid)
	}
	return this.Next_CALL_SelectOneObject
}
/**
* 中间件，替代函数入口
*/
func (this *DBMiddleware) SelectOneObject(sql string,bindarray []interface{},orm_ptr interface{})error {
    this.SelectOneObjectindex = 0
    return this.Next_CALL_SelectOneObject(sql,bindarray,orm_ptr)
}

/**
*/
func (this *DBMiddleware) Next_CALL_SelectOneObject(sql string,bindarray []interface{},orm_ptr interface{})error{
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.SelectOneObjectHandleFuncs) == 0 {
		this.Add_SelectOneObject(this.DB.SelectOneObject)
	} else if this.SelectOneObjectindex == 0 {
        // 👇👇---- 原始函数入口
		this.SelectOneObjectHandleFuncs = append(this.SelectOneObjectHandleFuncs, this.DB.SelectOneObject)
	}
    index := this.SelectOneObjectindex
	if this.SelectOneObjectindex >= len(this.SelectOneObjectHandleFuncs) {
		return nil	}

	this.SelectOneObjectindex++
	return this.SelectOneObjectHandleFuncs[index](sql,bindarray,orm_ptr)
}

func (this *DBMiddleware) Add_SelectAll(middlewares ...DB_SelectAllHandleFunc) DB_SelectAllHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.SelectAllHandleFuncs) == 0 {
        this.SelectAllHandleFuncs = append(this.SelectAllHandleFuncs, func(sql string,bindarray []interface{}) ([]map[string]string,error) {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").
                                Str("func","SelectAll").Str("timeuse",tc).Logger()
                    zloger.Debug().Msg("耗时")
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").
                            Interface("call_args",[]interface{}{sql,bindarray}).
                            Str("func","SelectAll").Logger()
                zloger.Debug().Msg("调起")
            }
            return this.Next_CALL_SelectAll(sql,bindarray)
        })
    }

    //
	if this.SelectAllHandleFuncs == nil {
		this.SelectAllHandleFuncs = make([]DB_SelectAllHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.SelectAllHandleFuncs = append(this.SelectAllHandleFuncs, mid)
	}
	return this.Next_CALL_SelectAll
}
/**
* 中间件，替代函数入口
*/
func (this *DBMiddleware) SelectAll(sql string,bindarray []interface{})([]map[string]string,error) {
    this.SelectAllindex = 0
    return this.Next_CALL_SelectAll(sql,bindarray)
}

/**
*/
func (this *DBMiddleware) Next_CALL_SelectAll(sql string,bindarray []interface{})([]map[string]string,error){
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.SelectAllHandleFuncs) == 0 {
		this.Add_SelectAll(this.DB.SelectAll)
	} else if this.SelectAllindex == 0 {
        // 👇👇---- 原始函数入口
		this.SelectAllHandleFuncs = append(this.SelectAllHandleFuncs, this.DB.SelectAll)
	}
    index := this.SelectAllindex
	if this.SelectAllindex >= len(this.SelectAllHandleFuncs) {
		return nil,nil	}

	this.SelectAllindex++
	return this.SelectAllHandleFuncs[index](sql,bindarray)
}

func (this *DBMiddleware) Add_SelectallObject(middlewares ...DB_SelectallObjectHandleFunc) DB_SelectallObjectHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.SelectallObjectHandleFuncs) == 0 {
        this.SelectallObjectHandleFuncs = append(this.SelectallObjectHandleFuncs, func(sql string,bindarray []interface{},orm_ptr interface{}) error {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").
                                Str("func","SelectallObject").Str("timeuse",tc).Logger()
                    zloger.Debug().Msg("耗时")
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").
                            Interface("call_args",[]interface{}{sql,bindarray,orm_ptr}).
                            Str("func","SelectallObject").Logger()
                zloger.Debug().Msg("调起")
            }
            return this.Next_CALL_SelectallObject(sql,bindarray,orm_ptr)
        })
    }

    //
	if this.SelectallObjectHandleFuncs == nil {
		this.SelectallObjectHandleFuncs = make([]DB_SelectallObjectHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.SelectallObjectHandleFuncs = append(this.SelectallObjectHandleFuncs, mid)
	}
	return this.Next_CALL_SelectallObject
}
/**
* 中间件，替代函数入口
*/
func (this *DBMiddleware) SelectallObject(sql string,bindarray []interface{},orm_ptr interface{})error {
    this.SelectallObjectindex = 0
    return this.Next_CALL_SelectallObject(sql,bindarray,orm_ptr)
}

/**
*/
func (this *DBMiddleware) Next_CALL_SelectallObject(sql string,bindarray []interface{},orm_ptr interface{})error{
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.SelectallObjectHandleFuncs) == 0 {
		this.Add_SelectallObject(this.DB.SelectallObject)
	} else if this.SelectallObjectindex == 0 {
        // 👇👇---- 原始函数入口
		this.SelectallObjectHandleFuncs = append(this.SelectallObjectHandleFuncs, this.DB.SelectallObject)
	}
    index := this.SelectallObjectindex
	if this.SelectallObjectindex >= len(this.SelectallObjectHandleFuncs) {
		return nil	}

	this.SelectallObjectindex++
	return this.SelectallObjectHandleFuncs[index](sql,bindarray,orm_ptr)
}

func (this *DBMiddleware) Add_SelectVar(middlewares ...DB_SelectVarHandleFunc) DB_SelectVarHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.SelectVarHandleFuncs) == 0 {
        this.SelectVarHandleFuncs = append(this.SelectVarHandleFuncs, func(sql string,bindarray []interface{}) (string,error) {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").
                                Str("func","SelectVar").Str("timeuse",tc).Logger()
                    zloger.Debug().Msg("耗时")
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").
                            Interface("call_args",[]interface{}{sql,bindarray}).
                            Str("func","SelectVar").Logger()
                zloger.Debug().Msg("调起")
            }
            return this.Next_CALL_SelectVar(sql,bindarray)
        })
    }

    //
	if this.SelectVarHandleFuncs == nil {
		this.SelectVarHandleFuncs = make([]DB_SelectVarHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.SelectVarHandleFuncs = append(this.SelectVarHandleFuncs, mid)
	}
	return this.Next_CALL_SelectVar
}
/**
* 中间件，替代函数入口
*/
func (this *DBMiddleware) SelectVar(sql string,bindarray []interface{})(string,error) {
    this.SelectVarindex = 0
    return this.Next_CALL_SelectVar(sql,bindarray)
}

/**
*/
func (this *DBMiddleware) Next_CALL_SelectVar(sql string,bindarray []interface{})(string,error){
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.SelectVarHandleFuncs) == 0 {
		this.Add_SelectVar(this.DB.SelectVar)
	} else if this.SelectVarindex == 0 {
        // 👇👇---- 原始函数入口
		this.SelectVarHandleFuncs = append(this.SelectVarHandleFuncs, this.DB.SelectVar)
	}
    index := this.SelectVarindex
	if this.SelectVarindex >= len(this.SelectVarHandleFuncs) {
		return "",nil	}

	this.SelectVarindex++
	return this.SelectVarHandleFuncs[index](sql,bindarray)
}

//检测接口是否被完整的实现了，如果没有实现，那么编译不通过
var _ DBInterface =new(DB)
