package log

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	Trace   = "Trace"
	debug   = "Debuge"
	Info    = "Info"
	warning = "Warning"
	error   = "Error"
)

func fileWriter(level string, filename string, msg string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed ERROR :\n %v", err)
	}
	defer file.Close()
	now := time.Now()
	time := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	writer := bufio.NewWriter(file)
	writer.WriteString(time + " " + level + " " + msg + "\n")
	writer.Flush()
}

func stdWriter(level string, msg string) {
	writer := bufio.NewWriter(os.Stdout)
	now := time.Now()
	time := fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	writer.WriteString(time + " " + level + " " + msg + "\n")
	writer.Flush()
}

func kafkaWriter(level string, msg string) {

}

func logFormat() {

}
