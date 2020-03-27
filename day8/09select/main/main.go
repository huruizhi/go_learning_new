package main

import (
	"fmt"
)


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
