package main

import "time"

type Car struct {
	name string
	age  int
}

type train struct {
	Car
	Start time.Time
	int
	age int // 首先使用本结构体内的
}

func main() {
	var t train
	t.name = "train"
	t.int = 6
	t.age = 10

}
