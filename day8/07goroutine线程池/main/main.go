package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. 启动goroutine ，生成100个数 发送到ch1
// 2. 启动一个goroutine ch1 取值 计算平方 放到ch2
// 3. 在main中 取出ch2的值打印出来

var sw sync.WaitGroup

func worker(jobs <-chan int, result chan<- int, id int) {
	for j := range jobs {
		fmt.Printf("第 %d 个线程开始执行任务%d\n", id, j)
		time.Sleep(time.Duration(j) * time.Second)
		fmt.Printf("第 %d 个线程执行任务结束%d\n", id, j)
		result <- j * j
	}
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	for i := 0; i < 3; i++ {
		go worker(jobs, result, i)
	}

	for j := 1; j < 5; j++ {
		jobs <- j
	}

	close(jobs)
	for {
		num := <-result
		fmt.Println(num)
	}

}
