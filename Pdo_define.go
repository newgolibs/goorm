package goorm
import (
        "database/sql"
        "github.com/rs/zerolog"
    "time"
)

//对象必须实现的接口方法
type PdoInterface interface {
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
    /**    提交事务    */
    Commit(recover interface{})
    /**    执行Query方法，返回rows    */
    query(sql string,bindarray []interface{})(*sql.Rows,[]interface{},[]sql.RawBytes,[]string)
    /**    当运行中，有一条sql错误了，那么回滚，在这个事务期间的所有操作全部报废    */
    Rollback()
    /**    查询count    */
    SelectVar(sql string,bindarray []interface{})(string,error)
    /**    提交事务，并且还继续开启事务    */
    Commit_NewTX()

}

//定义函数的结构体，方便扩展成中间件接
type Pdo_InsertHandleFunc func(sql string,bindarray []interface{})(int64,error)
type Pdo_ExecHandleFunc func(sql string,bindarray []interface{})(int64,error)
type Pdo_SelectOneHandleFunc func(sql string,bindarray []interface{})(map[string]string,error)
type Pdo_SelectOneObjectHandleFunc func(sql string,bindarray []interface{},orm_ptr interface{})error
type Pdo_SelectAllHandleFunc func(sql string,bindarray []interface{})([]map[string]string,error)
type Pdo_SelectallObjectHandleFunc func(sql string,bindarray []interface{},orm_ptr interface{})error
type Pdo_CommitHandleFunc func(recover interface{})
type Pdo_RollbackHandleFunc func()
type Pdo_SelectVarHandleFunc func(sql string,bindarray []interface{})(string,error)
type Pdo_Commit_NewTXHandleFunc func()

/**
数据库执行，返回数据;
*/
type Pdo struct
{
    /*事务链接句柄*/
    TX *sql.Tx
    /*数据库配置，中间件*/
    Pdoconfig *PdoconfigMiddleware
}






func (this *Pdo) NewPdoMiddleware() *PdoMiddleware{
    return &PdoMiddleware{Pdo:this}
}
/**
* 中间件的扩展类middleware
*/
type PdoMiddleware struct{
    Insertindex int
    InsertHandleFuncs []Pdo_InsertHandleFunc
    Execindex int
    ExecHandleFuncs []Pdo_ExecHandleFunc
    SelectOneindex int
    SelectOneHandleFuncs []Pdo_SelectOneHandleFunc
    SelectOneObjectindex int
    SelectOneObjectHandleFuncs []Pdo_SelectOneObjectHandleFunc
    SelectAllindex int
    SelectAllHandleFuncs []Pdo_SelectAllHandleFunc
    SelectallObjectindex int
    SelectallObjectHandleFuncs []Pdo_SelectallObjectHandleFunc
    Commitindex int
    CommitHandleFuncs []Pdo_CommitHandleFunc
    Rollbackindex int
    RollbackHandleFuncs []Pdo_RollbackHandleFunc
    SelectVarindex int
    SelectVarHandleFuncs []Pdo_SelectVarHandleFunc
    Commit_NewTXindex int
    Commit_NewTXHandleFuncs []Pdo_Commit_NewTXHandleFunc
    Pdo *Pdo
    //日志记录的目标文件
    zloger *zerolog.Logger
}

/**
设置日志写入类。
*/
func (this *PdoMiddleware) SetZloger(l *zerolog.Logger) *PdoMiddleware {
    Zloger := l.With().Str("library","goorm").
                    Str("class","PdoMiddleware").
                    Logger()
    this.zloger = &Zloger
    return this
}


func (this *PdoMiddleware) Add_Insert(middlewares ...Pdo_InsertHandleFunc) Pdo_InsertHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.InsertHandleFuncs) == 0 {
        this.InsertHandleFuncs = append(this.InsertHandleFuncs, func(sql string,bindarray []interface{}) (int64,error) {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").Logger()
                    zloger.Debug().Msgf("耗时 - Pdo.Insert:%+v",tc)
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").Logger()
                zloger.Debug().Msgf("调起 - Pdo.Insert，参数：%#v ",[]interface{}{sql,bindarray})
            }
            return this.Next_CALL_Insert(sql,bindarray)
        })
    }

    //
	if this.InsertHandleFuncs == nil {
		this.InsertHandleFuncs = make([]Pdo_InsertHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.InsertHandleFuncs = append(this.InsertHandleFuncs, mid)
	}
	return this.Next_CALL_Insert
}
/**
* 中间件，替代函数入口
*/
func (this *PdoMiddleware) Insert(sql string,bindarray []interface{})(int64,error) {
    this.Insertindex = 0
    return this.Next_CALL_Insert(sql,bindarray)
}

/**
*/
func (this *PdoMiddleware) Next_CALL_Insert(sql string,bindarray []interface{})(int64,error){
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.InsertHandleFuncs) == 0 {
		this.Add_Insert(this.Pdo.Insert)
	} else if this.Insertindex == 0 {
        // 👇👇---- 原始函数入口
		this.InsertHandleFuncs = append(this.InsertHandleFuncs, this.Pdo.Insert)
	}
    index := this.Insertindex
	if this.Insertindex >= len(this.InsertHandleFuncs) {
		return 0,nil	}

	this.Insertindex++
	return this.InsertHandleFuncs[index](sql,bindarray)
}

func (this *PdoMiddleware) Add_Exec(middlewares ...Pdo_ExecHandleFunc) Pdo_ExecHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.ExecHandleFuncs) == 0 {
        this.ExecHandleFuncs = append(this.ExecHandleFuncs, func(sql string,bindarray []interface{}) (int64,error) {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").Logger()
                    zloger.Debug().Msgf("耗时 - Pdo.Exec:%+v",tc)
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").Logger()
                zloger.Debug().Msgf("调起 - Pdo.Exec，参数：%#v ",[]interface{}{sql,bindarray})
            }
            return this.Next_CALL_Exec(sql,bindarray)
        })
    }

    //
	if this.ExecHandleFuncs == nil {
		this.ExecHandleFuncs = make([]Pdo_ExecHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.ExecHandleFuncs = append(this.ExecHandleFuncs, mid)
	}
	return this.Next_CALL_Exec
}
/**
* 中间件，替代函数入口
*/
func (this *PdoMiddleware) Exec(sql string,bindarray []interface{})(int64,error) {
    this.Execindex = 0
    return this.Next_CALL_Exec(sql,bindarray)
}

/**
*/
func (this *PdoMiddleware) Next_CALL_Exec(sql string,bindarray []interface{})(int64,error){
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.ExecHandleFuncs) == 0 {
		this.Add_Exec(this.Pdo.Exec)
	} else if this.Execindex == 0 {
        // 👇👇---- 原始函数入口
		this.ExecHandleFuncs = append(this.ExecHandleFuncs, this.Pdo.Exec)
	}
    index := this.Execindex
	if this.Execindex >= len(this.ExecHandleFuncs) {
		return 0,nil	}

	this.Execindex++
	return this.ExecHandleFuncs[index](sql,bindarray)
}

func (this *PdoMiddleware) Add_SelectOne(middlewares ...Pdo_SelectOneHandleFunc) Pdo_SelectOneHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.SelectOneHandleFuncs) == 0 {
        this.SelectOneHandleFuncs = append(this.SelectOneHandleFuncs, func(sql string,bindarray []interface{}) (map[string]string,error) {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").Logger()
                    zloger.Debug().Msgf("耗时 - Pdo.SelectOne:%+v",tc)
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").Logger()
                zloger.Debug().Msgf("调起 - Pdo.SelectOne，参数：%#v ",[]interface{}{sql,bindarray})
            }
            return this.Next_CALL_SelectOne(sql,bindarray)
        })
    }

    //
	if this.SelectOneHandleFuncs == nil {
		this.SelectOneHandleFuncs = make([]Pdo_SelectOneHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.SelectOneHandleFuncs = append(this.SelectOneHandleFuncs, mid)
	}
	return this.Next_CALL_SelectOne
}
/**
* 中间件，替代函数入口
*/
func (this *PdoMiddleware) SelectOne(sql string,bindarray []interface{})(map[string]string,error) {
    this.SelectOneindex = 0
    return this.Next_CALL_SelectOne(sql,bindarray)
}

/**
*/
func (this *PdoMiddleware) Next_CALL_SelectOne(sql string,bindarray []interface{})(map[string]string,error){
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.SelectOneHandleFuncs) == 0 {
		this.Add_SelectOne(this.Pdo.SelectOne)
	} else if this.SelectOneindex == 0 {
        // 👇👇---- 原始函数入口
		this.SelectOneHandleFuncs = append(this.SelectOneHandleFuncs, this.Pdo.SelectOne)
	}
    index := this.SelectOneindex
	if this.SelectOneindex >= len(this.SelectOneHandleFuncs) {
		return nil,nil	}

	this.SelectOneindex++
	return this.SelectOneHandleFuncs[index](sql,bindarray)
}

func (this *PdoMiddleware) Add_SelectOneObject(middlewares ...Pdo_SelectOneObjectHandleFunc) Pdo_SelectOneObjectHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.SelectOneObjectHandleFuncs) == 0 {
        this.SelectOneObjectHandleFuncs = append(this.SelectOneObjectHandleFuncs, func(sql string,bindarray []interface{},orm_ptr interface{}) error {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").Logger()
                    zloger.Debug().Msgf("耗时 - Pdo.SelectOneObject:%+v",tc)
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").Logger()
                zloger.Debug().Msgf("调起 - Pdo.SelectOneObject，参数：%#v ",[]interface{}{sql,bindarray,orm_ptr})
            }
            return this.Next_CALL_SelectOneObject(sql,bindarray,orm_ptr)
        })
    }

    //
	if this.SelectOneObjectHandleFuncs == nil {
		this.SelectOneObjectHandleFuncs = make([]Pdo_SelectOneObjectHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.SelectOneObjectHandleFuncs = append(this.SelectOneObjectHandleFuncs, mid)
	}
	return this.Next_CALL_SelectOneObject
}
/**
* 中间件，替代函数入口
*/
func (this *PdoMiddleware) SelectOneObject(sql string,bindarray []interface{},orm_ptr interface{})error {
    this.SelectOneObjectindex = 0
    return this.Next_CALL_SelectOneObject(sql,bindarray,orm_ptr)
}

/**
*/
func (this *PdoMiddleware) Next_CALL_SelectOneObject(sql string,bindarray []interface{},orm_ptr interface{})error{
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.SelectOneObjectHandleFuncs) == 0 {
		this.Add_SelectOneObject(this.Pdo.SelectOneObject)
	} else if this.SelectOneObjectindex == 0 {
        // 👇👇---- 原始函数入口
		this.SelectOneObjectHandleFuncs = append(this.SelectOneObjectHandleFuncs, this.Pdo.SelectOneObject)
	}
    index := this.SelectOneObjectindex
	if this.SelectOneObjectindex >= len(this.SelectOneObjectHandleFuncs) {
		return nil	}

	this.SelectOneObjectindex++
	return this.SelectOneObjectHandleFuncs[index](sql,bindarray,orm_ptr)
}

func (this *PdoMiddleware) Add_SelectAll(middlewares ...Pdo_SelectAllHandleFunc) Pdo_SelectAllHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.SelectAllHandleFuncs) == 0 {
        this.SelectAllHandleFuncs = append(this.SelectAllHandleFuncs, func(sql string,bindarray []interface{}) ([]map[string]string,error) {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").Logger()
                    zloger.Debug().Msgf("耗时 - Pdo.SelectAll:%+v",tc)
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").Logger()
                zloger.Debug().Msgf("调起 - Pdo.SelectAll，参数：%#v ",[]interface{}{sql,bindarray})
            }
            return this.Next_CALL_SelectAll(sql,bindarray)
        })
    }

    //
	if this.SelectAllHandleFuncs == nil {
		this.SelectAllHandleFuncs = make([]Pdo_SelectAllHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.SelectAllHandleFuncs = append(this.SelectAllHandleFuncs, mid)
	}
	return this.Next_CALL_SelectAll
}
/**
* 中间件，替代函数入口
*/
func (this *PdoMiddleware) SelectAll(sql string,bindarray []interface{})([]map[string]string,error) {
    this.SelectAllindex = 0
    return this.Next_CALL_SelectAll(sql,bindarray)
}

/**
*/
func (this *PdoMiddleware) Next_CALL_SelectAll(sql string,bindarray []interface{})([]map[string]string,error){
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.SelectAllHandleFuncs) == 0 {
		this.Add_SelectAll(this.Pdo.SelectAll)
	} else if this.SelectAllindex == 0 {
        // 👇👇---- 原始函数入口
		this.SelectAllHandleFuncs = append(this.SelectAllHandleFuncs, this.Pdo.SelectAll)
	}
    index := this.SelectAllindex
	if this.SelectAllindex >= len(this.SelectAllHandleFuncs) {
		return nil,nil	}

	this.SelectAllindex++
	return this.SelectAllHandleFuncs[index](sql,bindarray)
}

func (this *PdoMiddleware) Add_SelectallObject(middlewares ...Pdo_SelectallObjectHandleFunc) Pdo_SelectallObjectHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.SelectallObjectHandleFuncs) == 0 {
        this.SelectallObjectHandleFuncs = append(this.SelectallObjectHandleFuncs, func(sql string,bindarray []interface{},orm_ptr interface{}) error {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").Logger()
                    zloger.Debug().Msgf("耗时 - Pdo.SelectallObject:%+v",tc)
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").Logger()
                zloger.Debug().Msgf("调起 - Pdo.SelectallObject，参数：%#v ",[]interface{}{sql,bindarray,orm_ptr})
            }
            return this.Next_CALL_SelectallObject(sql,bindarray,orm_ptr)
        })
    }

    //
	if this.SelectallObjectHandleFuncs == nil {
		this.SelectallObjectHandleFuncs = make([]Pdo_SelectallObjectHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.SelectallObjectHandleFuncs = append(this.SelectallObjectHandleFuncs, mid)
	}
	return this.Next_CALL_SelectallObject
}
/**
* 中间件，替代函数入口
*/
func (this *PdoMiddleware) SelectallObject(sql string,bindarray []interface{},orm_ptr interface{})error {
    this.SelectallObjectindex = 0
    return this.Next_CALL_SelectallObject(sql,bindarray,orm_ptr)
}

/**
*/
func (this *PdoMiddleware) Next_CALL_SelectallObject(sql string,bindarray []interface{},orm_ptr interface{})error{
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.SelectallObjectHandleFuncs) == 0 {
		this.Add_SelectallObject(this.Pdo.SelectallObject)
	} else if this.SelectallObjectindex == 0 {
        // 👇👇---- 原始函数入口
		this.SelectallObjectHandleFuncs = append(this.SelectallObjectHandleFuncs, this.Pdo.SelectallObject)
	}
    index := this.SelectallObjectindex
	if this.SelectallObjectindex >= len(this.SelectallObjectHandleFuncs) {
		return nil	}

	this.SelectallObjectindex++
	return this.SelectallObjectHandleFuncs[index](sql,bindarray,orm_ptr)
}

func (this *PdoMiddleware) Add_Commit(middlewares ...Pdo_CommitHandleFunc) Pdo_CommitHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.CommitHandleFuncs) == 0 {
        this.CommitHandleFuncs = append(this.CommitHandleFuncs, func(recover interface{})  {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").Logger()
                    zloger.Debug().Msgf("耗时 - Pdo.Commit:%+v",tc)
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").Logger()
                zloger.Debug().Msgf("调起 - Pdo.Commit，参数：%#v ",[]interface{}{recover})
            }
            this.Next_CALL_Commit(recover)
        })
    }

    //
	if this.CommitHandleFuncs == nil {
		this.CommitHandleFuncs = make([]Pdo_CommitHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.CommitHandleFuncs = append(this.CommitHandleFuncs, mid)
	}
	return this.Next_CALL_Commit
}
/**
* 中间件，替代函数入口
*/
func (this *PdoMiddleware) Commit(recover interface{}) {
    this.Commitindex = 0
    this.Next_CALL_Commit(recover)
}

/**
*/
func (this *PdoMiddleware) Next_CALL_Commit(recover interface{}){
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.CommitHandleFuncs) == 0 {
		this.Add_Commit(this.Pdo.Commit)
	} else if this.Commitindex == 0 {
        // 👇👇---- 原始函数入口
		this.CommitHandleFuncs = append(this.CommitHandleFuncs, this.Pdo.Commit)
	}
    index := this.Commitindex
	if this.Commitindex >= len(this.CommitHandleFuncs) {
        return
	}

	this.Commitindex++
    this.CommitHandleFuncs[index](recover)
}

func (this *PdoMiddleware) Add_Rollback(middlewares ...Pdo_RollbackHandleFunc) Pdo_RollbackHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.RollbackHandleFuncs) == 0 {
        this.RollbackHandleFuncs = append(this.RollbackHandleFuncs, func()  {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").Logger()
                    zloger.Debug().Msgf("耗时 - Pdo.Rollback:%+v",tc)
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").Logger()
                zloger.Debug().Msgf("调起 - Pdo.Rollback，参数：%#v ",[]interface{}{})
            }
            this.Next_CALL_Rollback()
        })
    }

    //
	if this.RollbackHandleFuncs == nil {
		this.RollbackHandleFuncs = make([]Pdo_RollbackHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.RollbackHandleFuncs = append(this.RollbackHandleFuncs, mid)
	}
	return this.Next_CALL_Rollback
}
/**
* 中间件，替代函数入口
*/
func (this *PdoMiddleware) Rollback() {
    this.Rollbackindex = 0
    this.Next_CALL_Rollback()
}

/**
*/
func (this *PdoMiddleware) Next_CALL_Rollback(){
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.RollbackHandleFuncs) == 0 {
		this.Add_Rollback(this.Pdo.Rollback)
	} else if this.Rollbackindex == 0 {
        // 👇👇---- 原始函数入口
		this.RollbackHandleFuncs = append(this.RollbackHandleFuncs, this.Pdo.Rollback)
	}
    index := this.Rollbackindex
	if this.Rollbackindex >= len(this.RollbackHandleFuncs) {
        return
	}

	this.Rollbackindex++
    this.RollbackHandleFuncs[index]()
}

func (this *PdoMiddleware) Add_SelectVar(middlewares ...Pdo_SelectVarHandleFunc) Pdo_SelectVarHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.SelectVarHandleFuncs) == 0 {
        this.SelectVarHandleFuncs = append(this.SelectVarHandleFuncs, func(sql string,bindarray []interface{}) (string,error) {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").Logger()
                    zloger.Debug().Msgf("耗时 - Pdo.SelectVar:%+v",tc)
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").Logger()
                zloger.Debug().Msgf("调起 - Pdo.SelectVar，参数：%#v ",[]interface{}{sql,bindarray})
            }
            return this.Next_CALL_SelectVar(sql,bindarray)
        })
    }

    //
	if this.SelectVarHandleFuncs == nil {
		this.SelectVarHandleFuncs = make([]Pdo_SelectVarHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.SelectVarHandleFuncs = append(this.SelectVarHandleFuncs, mid)
	}
	return this.Next_CALL_SelectVar
}
/**
* 中间件，替代函数入口
*/
func (this *PdoMiddleware) SelectVar(sql string,bindarray []interface{})(string,error) {
    this.SelectVarindex = 0
    return this.Next_CALL_SelectVar(sql,bindarray)
}

/**
*/
func (this *PdoMiddleware) Next_CALL_SelectVar(sql string,bindarray []interface{})(string,error){
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.SelectVarHandleFuncs) == 0 {
		this.Add_SelectVar(this.Pdo.SelectVar)
	} else if this.SelectVarindex == 0 {
        // 👇👇---- 原始函数入口
		this.SelectVarHandleFuncs = append(this.SelectVarHandleFuncs, this.Pdo.SelectVar)
	}
    index := this.SelectVarindex
	if this.SelectVarindex >= len(this.SelectVarHandleFuncs) {
		return "",nil	}

	this.SelectVarindex++
	return this.SelectVarHandleFuncs[index](sql,bindarray)
}

func (this *PdoMiddleware) Add_Commit_NewTX(middlewares ...Pdo_Commit_NewTXHandleFunc) Pdo_Commit_NewTXHandleFunc {
    // 第一个添加的是日志，如果设置了写出源的话，比如,os.Stdout
    if len(this.Commit_NewTXHandleFuncs) == 0 {
        this.Commit_NewTXHandleFuncs = append(this.Commit_NewTXHandleFuncs, func()  {
            defer func(start time.Time) {
                if this.zloger != nil {
                    tc := time.Since(start).String()
                    zloger := this.zloger.With().Str("fun_middle_type","timeuse").Logger()
                    zloger.Debug().Msgf("耗时 - Pdo.Commit_NewTX:%+v",tc)
                }
            }(time.Now())
            if this.zloger != nil {
                zloger := this.zloger.With().Str("fun_middle_type","call_args").Logger()
                zloger.Debug().Msgf("调起 - Pdo.Commit_NewTX，参数：%#v ",[]interface{}{})
            }
            this.Next_CALL_Commit_NewTX()
        })
    }

    //
	if this.Commit_NewTXHandleFuncs == nil {
		this.Commit_NewTXHandleFuncs = make([]Pdo_Commit_NewTXHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.Commit_NewTXHandleFuncs = append(this.Commit_NewTXHandleFuncs, mid)
	}
	return this.Next_CALL_Commit_NewTX
}
/**
* 中间件，替代函数入口
*/
func (this *PdoMiddleware) Commit_NewTX() {
    this.Commit_NewTXindex = 0
    this.Next_CALL_Commit_NewTX()
}

/**
*/
func (this *PdoMiddleware) Next_CALL_Commit_NewTX(){
    // 调起的时候，追加源功能函数。因为源功能函数没有调起NEXT，所以只有执行到它，必定阻断后面的所有中间件函数。
	if len(this.Commit_NewTXHandleFuncs) == 0 {
		this.Add_Commit_NewTX(this.Pdo.Commit_NewTX)
	} else if this.Commit_NewTXindex == 0 {
        // 👇👇---- 原始函数入口
		this.Commit_NewTXHandleFuncs = append(this.Commit_NewTXHandleFuncs, this.Pdo.Commit_NewTX)
	}
    index := this.Commit_NewTXindex
	if this.Commit_NewTXindex >= len(this.Commit_NewTXHandleFuncs) {
        return
	}

	this.Commit_NewTXindex++
    this.Commit_NewTXHandleFuncs[index]()
}

//检测接口是否被完整的实现了，如果没有实现，那么编译不通过
var _ PdoInterface =new(Pdo)
