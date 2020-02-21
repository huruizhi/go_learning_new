package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

/*
使用 goroutine 和channel 实现一个计算int64随机数各位数的程序
1. 开启一个goroutine 循环生成int64的随机数 发送到jobChan
2. 开24 个 goroutine 从jobChan 中取出并计算各位之和保存到resultChan
3.  主goroutine 从 resultChan 中取出结果打印到终端
*/

var wg sync.WaitGroup
var once sync.Once

func randNum(jobChan chan<- int64) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		jobChan <- rand.Int63()
	}
	close(jobChan)
}

func sunNum(jobChan <-chan int64, resultChan chan<- string) {
	defer wg.Done()
	for {
		var s string
		var n int
		num, ok := <-jobChan
		if !ok {
			break
		}
		s = fmt.Sprintf("%d: ", num)
		for {
			num2 := int(num % 10)
			n = n + int(num2)
			num = num / 10
			s = s + "+" + strconv.Itoa(num2)
			if num == 0 {
				s = fmt.Sprintf("%s=%d", s, n)
				resultChan <- s
				time.Sleep(time.Second)
				break
			}
		}
	}
}

func client(resultChan <-chan string) {
	defer wg.Done()
	num := 1
	for {
		s, ok := <-resultChan
		if !ok {
			break
		}
		fmt.Printf("%d,%s\n", num, s)
		num += 1
	}
}

func main() {
	jobChan := make(chan int64, 100)
	resultChan := make(chan string, 100)
	go randNum(jobChan)

	for i := 0; i < 24; i++ {
		wg.Add(1)
		go sunNum(jobChan, resultChan)
	}

	wg.Add(1)
	go client(resultChan)

	wg.Wait()

}
