package main

import (
	"fmt"
	"runtime"
	"sync"
)

// GOMAXPROCS

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
	// runtime.GOMAXPROCS(2)
	wg.Add(2)
	go b()
	go a()
	wg.Wait()
}
