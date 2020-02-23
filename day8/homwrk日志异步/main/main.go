package main

import (
	"go_learning_new/day8/homwrk日志异步/log"
	"time"
)

/*
1. 日志中channel 怎么用
2. 什么时候起后台日志到文件中
*/

func main() {
	l := log.NewLogFileWriter("./", "test.log", "Info")
	for i := 0; i < 100; i++ {
		l.Error("第 %d 个错误日志", i)
		time.Sleep(200 * time.Millisecond)
	}
}
