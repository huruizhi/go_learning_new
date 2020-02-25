package main

import (
	"fmt"
	"sync"
)

var x int
var lock sync.RWMutex
var sw sync.WaitGroup

func read() {
	lock.RLock()
	fmt.Println(x)
	lock.RUnlock()
	sw.Done()
}

func write() {
	lock.Lock()
	x++
	lock.Unlock()
	sw.Done()
}

func main() {
	for i := 0; i < 100; i++ {
		sw.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		sw.Add(1)
		go read()
	}

	sw.Wait()

}
