package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
1. golang的包
	1. 目前有150 多个标准包， 覆盖了几乎所有的基础库
	2. golang。org有所有包文档，没事要多看
2. 线程同步
	1. import("sync")
	2. 互斥所 var mu sync.Mutex
	3. 读写锁 var mu sync.RWMutex 读锁可以同时读
	读写锁在读多写少情况下性能高很多
3. go get 安装包
*/

var lock sync.Mutex
var rwlock sync.RWMutex

func test() {
	var a = map[string]int{"a": 0, "b": 2}
	var count int32

	for i := 0; i < 3; i++ {
		go func(b map[string]int) {
			lock.Lock()
			b["a"] = i
			lock.Unlock()
		}(a)
	}
	fmt.Println(a)

	for i := 0; i < 100; i++ {
		go func(b map[string]int) {
			rwlock.RLock()
			// fmt.Println(a)
			rwlock.RUnlock()
			atomic.AddInt32(&count, 1)
		}(a)
		fmt.Println(atomic.LoadInt32(&count))
	}

}

func main() {
	for i := 0; i < 10; i++ {
		test()
	}
}
