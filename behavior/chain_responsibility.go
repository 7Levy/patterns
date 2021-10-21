package behavior

import "fmt"

const (
	DebugLogLevel = iota
	InfoLogLevel
	ErrorLogLevel
)

type BaseLogger interface {
	PrintLog(level int, message string)
	Write(message string)
}

type DebugLogger struct {
	Level      int
	NextLogger BaseLogger
}

func NewDebugLogger() *DebugLogger {
	return &DebugLogger{
		Level:      0,
		NextLogger: nil,
	}
}
func (dl *DebugLogger) Write(message string) {
	fmt.Printf("Debug Logger out : %s.\n", message)
}
func (dl *DebugLogger) PrintLog(level int, message string) {
	if dl.Level == level {
		dl.Write(message)
	}
	if dl.NextLogger != nil {
		dl.NextLogger.PrintLog(level, message)
	}
}

//SetNextLogger 标准日志类设置下一个对象方法
func (dl *DebugLogger) SetNextLogger(logger BaseLogger) {
	dl.NextLogger = logger
}

//InfoLogger 提示日志类
type InfoLogger struct {
	Level      int
	NextLogger BaseLogger
}

//NewInfoLogger 实例化提示日志类
func NewInfoLogger() *InfoLogger {
	return &InfoLogger{
		Level:      InfoLogLevel,
		NextLogger: nil,
	}
}

//Write 提示日志类的写方法
func (infoL *InfoLogger) Write(message string) {
	fmt.Printf("Info Logger out: %s.\n", message)
}

//PrintLog 提示日志类的输入日志方法
func (infoL *InfoLogger) PrintLog(level int, message string) {
	if infoL.Level == level {
		infoL.Write(message)
	}
	if infoL.NextLogger != nil {
		infoL.NextLogger.PrintLog(level, message)
	}
}

//SetNextLogger 提示日志类设置下一个对象
func (infoL *InfoLogger) SetNextLogger(logger BaseLogger) {
	infoL.NextLogger = logger
}

//ErrorLogger 错误日志类
type ErrorLogger struct {
	Level      int
	NextLogger BaseLogger
}

//NewErrorLogger 实例化错误日志类
func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{
		Level:      ErrorLogLevel,
		NextLogger: nil,
	}
}

//Write 错误日志类写方法
func (el *ErrorLogger) Write(message string) {
	fmt.Printf("Error logger out: %s.\n", message)
}

//PrintLog 错误日志类输入日志方法
func (el *ErrorLogger) PrintLog(level int, message string) {
	if el.Level == level {
		el.Write(message)
	}
	if el.NextLogger != nil {
		el.NextLogger.PrintLog(level, message)
	}
}

//SetNextLogger 错误日志类设置下一个对象
func (el *ErrorLogger) SetNextLogger(logger BaseLogger) {
	el.NextLogger = logger
}

//GetChainOfLoggers 获取日志对象链
func GetChainOfLoggers() BaseLogger {
	errLog := NewErrorLogger()
	infoLog := NewInfoLogger()
	debugLog := NewDebugLogger()

	errLog.SetNextLogger(infoLog)
	infoLog.SetNextLogger(debugLog)

	return errLog
}
