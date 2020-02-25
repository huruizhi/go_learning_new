package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64
var sw sync.WaitGroup

func add() {
	defer sw.Done()
	atomic.AddInt64(&x, 1) //原子性的操作，省略加锁
}

func main() {
	for i := 0; i < 10000; i++ {
		sw.Add(1)
		go add()
	}

	sw.Wait()
	fmt.Println(x)
}
