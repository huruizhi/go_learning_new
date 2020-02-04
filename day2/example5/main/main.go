package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("随机整数：")
	for i := 0; i < 10; i++ {
		num := rand.Int()
		fmt.Println(num)
	}
	fmt.Println("0~100 随机整数：")
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		fmt.Println(num)
	}
	fmt.Println("随机浮点数：")
	for i := 0; i < 10; i++ {
		num := rand.Float32()
		fmt.Println(num)
	}
}
