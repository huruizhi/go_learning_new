package main

import (
	"fmt"
)

/*
使用 goroutine 和channel 实现一个计算int64随机数各位数的程序
1. 开启一个goroutine 循环生成int64的随机数 发送到jobChan
2. 开24 个 goroutine 从jobChan 中取出并计算各位之和保存到resultChan
3.  主goroutine 从 resultChan 中取出结果打印到终端
*/

func main() {
	chan1 := make(chan int64, 1)
	//select基本用法
	for i := 1; i < 20; i++ {
		select {
		case n := <-chan1:
			fmt.Println(n)

			// 如果chan1成功读到数据，则进行该case处理语句
		case chan1 <- 1:
			// 如果成功向chan1写入数据，则进行该case处理语句
		default:
			// 如果上面都没有成功，则进入default处理流程
		}
	}
}
