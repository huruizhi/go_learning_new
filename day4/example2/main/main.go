package main

import (
	"fmt"
	"time"
)

func recusive(n int) {
	fmt.Println(n, " ", "hello")
	time.Sleep(time.Second)
	n++
	if n > 10 {
		return
	} else {
		recusive(n)
	}

}

func factor(n int) int {
	if n == 1 {
		return n
	} else {
		n = n * factor(n-1)
		return n
	}

}

func main() {
	//recusive(1)
	fmt.Println(factor(4))
}
