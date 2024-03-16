package logging

/*

Этот файл определяет интерфейс Logger, который определяет методы регистрации 
сообщений с различными уровнями серьезности, которые задаются 
с использованием значения LogLevel в диапазоне от Trace до Fatal

*/

type LogLevel int 

const (
	Trace LogLevel = iota
	Debug
	Information
	Warning
	Fatal
	None
)

type Logger interface{
	Trace(string)
	Tracef(string, ...interface{})

	Debug(string)
	Debugf(string, ...interface{})

	Info(string)
	Infof(string, ...interface{})

	Warn(string)
	Warnf(string, ...interface{})

	Panic(string)
	Panicf(string, ...interface{})
}