package log

import "fmt"

const (
	StdOut = iota
	LogFile
	Kafka
)

type logLevel int8

type logging struct {
	level        logLevel
	outputMethod int
	logfile      string
}

func NewLogging(level string, flag int) logging {
	loglevel, err := traceLogLevel(level)
	if err != nil {
		panic("log level error!")
	}
	return logging{
		outputMethod: flag,
		level:        loglevel,
	}
}

func (l *logging) SetLogFile(filename string) {
	if l.outputMethod == LogFile {
		l.logfile = filename
	} else {
		fmt.Println("设置日志文件时，outputMethod 需要为 LogFile")
	}
}

func (l logging) Trace(format string, a ...interface{}) {
	if l.enable(TRACE) {
		l.writerLog(TRACE, format, a...)
	}
}

func (l logging) Debug(format string, a ...interface{}) {
	if l.enable(DEBUG) {
		l.writerLog(DEBUG, format, a...)
	}
}

func (l logging) Info(format string, a ...interface{}) {
	if l.enable(INFO) {
		l.writerLog(INFO, format, a...)
	}
}

func (l logging) Warning(format string, a ...interface{}) {
	if l.enable(WARNING) {
		l.writerLog(WARNING, format, a...)
	}
}

func (l logging) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		l.writerLog(ERROR, format, a...)
	}
}

func (l logging) Fatal(format string, a ...interface{}) {
	if l.enable(FATAL) {
		l.writerLog(FATAL, format, a...)
	}
}

func (l logging) writerLog(level logLevel, format string, a ...interface{}) {
	switch l.outputMethod {
	case StdOut:
		stdWriter(level, format, a...)
	case LogFile:
		fileWriter(level, l.logfile, format, a...)
	case Kafka:
		kafkaWriter(level, format)
	}
}

func (l logging) enable(lv logLevel) bool {
	return lv >= l.level
}
