package log

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
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

type logLevel int

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

func parseLogLevel(level logLevel) (string, error) {
	switch level {
	case TRACE:
		return "TRACE", nil
	case DEBUG:
		return "DEBUG", nil
	case INFO:
		return "INFO", nil
	case WARNING:
		return "WARNING", nil
	case ERROR:
		return "ERROR", nil
	case FATAL:
		return "FATAL", nil
	}
	err := errors.New("Get log level failed!")
	return "", err
}

var lock sync.Mutex

type logFileWriter struct {
	fileName string
	filePath string
	level    logLevel
	fileObj  *os.File
	logChan  chan string
}

func NewLogFileWriter(filePath string, fileName string, level string) logFileWriter {
	file := path.Join(filePath, fileName)
	fileObj, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	levelInfo, err := traceLogLevel(level)
	if err != nil {
		panic(err)
	}

	l := logFileWriter{
		fileName: fileName,
		filePath: filePath,
		level:    levelInfo,
		fileObj:  fileObj,
		logChan:  make(chan string, 1000),
	}
	go l.logFileCuter()
	go l.writeLog()
	return l
}

func funcInfo(skip int) (file string, fcName string, line int) {
	pc, file, line, _ := runtime.Caller(skip)
	fcInfo := runtime.FuncForPC(pc)
	fcName = strings.Split(fcInfo.Name(), ".")[1]
	return
}

func (l logFileWriter) logToChan(level logLevel, format string, a ...interface{}) {

	msg := fmt.Sprintf(format, a...)
	levelstr, err := parseLogLevel(level)

	if err != nil {
		return
	}

	file, fcName, line := funcInfo(3)

	_, fileName := path.Split(file)

	datetime := time.Now().Format("2006-01-02 15:04:05")

	logMsg := fmt.Sprintf("%s [%s] %s:%s:%d %s\n", datetime, levelstr, fileName, fcName, line, msg)
	l.logChan <- logMsg
}

func (l *logFileWriter) writeLog() {
	for {
		lock.Lock()
		logMsg, ok := <-l.logChan
		if !ok {
			continue
		}
		_, err := l.fileObj.WriteString(logMsg)
		if err != nil {
			fmt.Println(err)
			continue
		}
		lock.Unlock()
	}
}

func (l logFileWriter) Trace(format string, a ...interface{}) {
	if l.enable(TRACE) {
		l.logToChan(TRACE, format, a...)
	}
}

func (l logFileWriter) Debug(format string, a ...interface{}) {
	if l.enable(DEBUG) {
		l.logToChan(DEBUG, format, a...)
	}
}

func (l logFileWriter) Info(format string, a ...interface{}) {
	if l.enable(INFO) {
		l.logToChan(INFO, format, a...)
	}
}

func (l logFileWriter) Warning(format string, a ...interface{}) {
	if l.enable(WARNING) {
		l.logToChan(WARNING, format, a...)
	}
}

func (l logFileWriter) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		l.logToChan(ERROR, format, a...)
	}
}

func (l logFileWriter) Fatal(format string, a ...interface{}) {
	if l.enable(FATAL) {
		l.logToChan(FATAL, format, a...)
	}
}

func (l logFileWriter) enable(lv logLevel) bool {
	return lv >= l.level
}

func (l *logFileWriter) logFileCuter() {
	min := time.Now().Minute()
	for {

		minNow := time.Now().Minute()
		if time.Now().Second() == 0 && minNow != min {
			min = minNow
			lock.Lock()
			l.fileObj.Close()
			file := path.Join(l.filePath, time.Now().Format("20060102150405.log"))
			os.Rename(path.Join(l.filePath, l.fileName), file)
			fmt.Printf("切割日志到%s\n", file)

			file = path.Join(l.filePath, l.fileName)
			l.fileObj, _ = os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			lock.Unlock()

		}
	}
}
