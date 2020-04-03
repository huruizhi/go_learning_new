package main

import (
	"fmt"
)

type animal interface {
	speak()
}

type dog struct{}

type cat struct{}

func (d dog) speak() {
	fmt.Println("wangwangwang")
}

func (c cat) speak() {
	fmt.Println("miaomiaomiao")
}

func hit(a animal) { // 实现了接口animal 中speak方法的结构体 都可以作为参数传入
	a.speak()
}

func main() {
	d := dog{}
	hit(d)
	c := cat{}
	hit(c)
}
