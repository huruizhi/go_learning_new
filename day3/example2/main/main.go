package main

import (
	"fmt"
	"time"
)

func f() {
	time.Sleep(time.Millisecond * 1)
}

func main() {
	// 时间格式化
	now := time.Now()
	fmt.Println(now.Format("02/1/2006 15:04"))
	fmt.Println(now.Format("2006/1/02 15:04"))
	fmt.Println(now.Format("2006/1/02"))

	begin := time.Now().UnixNano()
	f()
	end := time.Now().UnixNano()

	fmt.Println(end - begin)

}
