package goorm
//对象必须实现的接口方法
type Logger interface {
    /**    输出调试信息    */
    Debug(format string,v ...interface{})
    /**    输出错误信息    */
    Error(format string,v ...interface{})
}