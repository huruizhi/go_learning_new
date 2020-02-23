package main

import (
	"fmt"
	"sync"
)

var x int = 0
var sw sync.WaitGroup
var lock sync.Mutex

func add() {
	defer sw.Done()
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
}

func main() {
	sw.Add(2)
	go add()
	go add()
	sw.Wait()
	fmt.Println(x)
}
