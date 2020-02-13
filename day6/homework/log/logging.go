package log

import "fmt"

const (
	StdOut = iota
	LogFile
	Kafka
)

type logging struct {
	outputMethod int
	logfile      string
}

func NewLogging(flag int) logging {
	return logging{
		outputMethod: flag,
	}
}

func (l *logging) SetLogFile(filename string) {
	if l.outputMethod == LogFile {
		l.logfile = filename
	} else {
		fmt.Println("设置日志文件时，outputMethod 需要为 LogFile")
	}
}

func (l logging) Info(msg string) {
	switch l.outputMethod {
	case StdOut:
		stdWriter(Info, msg)
	case LogFile:
		fileWriter(Info, l.logfile, msg)
	case Kafka:
		kafkaWriter(Info, msg)
	}
}

func (l logging) Trace(msg string) {
	switch l.outputMethod {
	case StdOut:
		stdWriter(Trace, msg)
	case LogFile:
		fileWriter(Trace, l.logfile, msg)
	case Kafka:
		kafkaWriter(Trace, msg)
	}
}
