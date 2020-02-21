package main

import (
	"fmt"
	"sync"
)

func send(c chan int) {
	// 通道的操作
	// 1. 发送
	a1 := 1
	c <- a1
	defer sw.Done()
}

func receive(c chan int) {
	// 2.接收值
	a2 := <-c
	fmt.Println(a2)
	defer sw.Done()
}

var sw sync.WaitGroup

func main() {
	var c chan int
	c = make(chan int)      //不含缓冲区的初始化
	b := make(chan int, 16) //带缓冲区的初始化

	// 不带缓冲区的 需要将接收放置于后台
	sw.Add(1)
	go send(c)
	sw.Add(1)
	go receive(c)

	//带缓冲区
	a1 := 2
	b <- a1
	a2 := <-b
	fmt.Println(a2)

	// 3. 关闭
	defer close(c)
	defer close(b)
	sw.Wait()

}
