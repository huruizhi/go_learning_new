package log

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

const (
	TRACE = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

func traceLogLevel(level string) (logLevel, error) {
	level = strings.ToLower(level)
	switch level {
	case "trace":
		return TRACE, nil
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	}
	err := errors.New("Get log level failed!")
	return -1, err

}

func parseLogLevel(lv logLevel) string {
	switch lv {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}

func fileWriter(level logLevel, filename string, format string, a ...interface{}) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed ERROR :\n %v", err)
	}
	defer file.Close()
	now := time.Now()
	writer := bufio.NewWriter(file)
	funcName, fileName, lineNo := getInfo(4)
	msg := fmt.Sprintf(format, a...)
	logString := fmt.Sprintf("%s [%s] %s:%s:%d %s\n", now.Format("2006-01-02 15:04:05"), parseLogLevel(level), fileName, funcName, lineNo, msg)
	writer.WriteString(logString)
	writer.Flush()
}

func stdWriter(level logLevel, format string, a ...interface{}) {
	writer := bufio.NewWriter(os.Stdout)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(4)

	msg := fmt.Sprintf(format, a...)
	logString := fmt.Sprintf("%s [%s] %s:%s:%d %s\n", now.Format("2006-01-02 15:04:05"), parseLogLevel(level), fileName, funcName, lineNo, msg)
	writer.WriteString(logString)
	writer.Flush()
}

func kafkaWriter(level logLevel, msg string) {

}

func getInfo(skip int) (string, string, int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		panic("call runtime failed!")
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1] + "()"
	_, fileName := path.Split(path.Base(file))
	lineNo := line
	return funcName, fileName, lineNo

}
