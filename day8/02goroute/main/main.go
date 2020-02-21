package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var sw sync.WaitGroup

func f1(i int) {
	sw.Add(1)       //添加计数器
	defer sw.Done() // 任务结束 计数器减1
	rand.Seed(time.Now().UnixNano())
	second := rand.Intn(i)
	time.Sleep(time.Duration(second) * time.Second)
	fmt.Println(i, ":", second)
}

func main() {
	for i := 1; i < 10; i++ {
		go f1(i)
	}
	sw.Wait() // 等待所有的计数器退出
}

// M结构是Machine，系统线程，它由操作系统管理，goroutine就是跑在M之上的；M是一个很大的结构，里面维护小对象内存cache（mcache）、当前执行的goroutine、随机数发生器等等非常多的信息
// P结构是Processor，处理器，它的主要用途就是用来执行goroutine，它维护了一个goroutine队列，即runqueue。Processor的让我们从N:1调度到M:N调度的重要部分。
// G是goroutine实现的核心结构，它包含了栈，指令指针，以及其他对调度goroutine很重要的信息，例如其阻塞的channel。
