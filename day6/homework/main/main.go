package main

import (
	"go_learning_new/day6/homework/log"
	"time"
)

func f1() {
	l := log.NewLogging("info", log.StdOut)
	// l.SetLogFile("logfile")
	msg := `adafasddafadfbcdefg abcdefg`
	for {
		l.Error(msg)
		l.Info("name:%s, age:%d", "DD", 18)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	f1()
}

//需求 日志
//1. 打印出日期 文件函数信息
//2. 分不同的日志级别
//3. 设置日志开关 打印高于该级别的日志
//4. 输出到控制台或者 日志文件
