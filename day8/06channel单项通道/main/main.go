package main

import (
	"fmt"
	"sync"
)

// 1. 启动goroutine ，生成100个数 发送到ch1
// 2. 启动一个goroutine ch1 取值 计算平方 放到ch2
// 3. 在main中 取出ch2的值打印出来

var sw sync.WaitGroup

func f1(c chan<- int) { //chan<-  只能向channel 传值
	defer sw.Done()
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c) //关闭channel 不会影响channel 的读取
}

func f2(c1 <-chan int, c2 chan<- int) { // <-chan 只能从channel 取值
	defer sw.Done()
	for {
		num, ok := <-c1
		if !ok {
			break
		}
		c2 <- num * num
	}
	close(c2) //关闭channel 不会影响channel 的读取
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	sw.Add(1)
	go f1(ch1)
	sw.Add(1)
	go f2(ch1, ch2)

	for i := range ch2 {
		fmt.Println(i)
	}
	sw.Wait()
}
