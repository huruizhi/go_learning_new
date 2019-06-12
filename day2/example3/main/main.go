package main

import (
	"fmt"
	"os"
)
// 传入变量
func func1(a int)  {
	a = 10
}
// 传入指针
func func2(a *int)  {
	*a = 10
}

func main() {
	/*
	值类型： int float bool string
	引用类型： 指针 slice map chan
	 */
	var (
		a = 5
		goos string = os.Getenv("GOPATH")
	)
	fmt.Printf("path=%s\n",goos)
	func1(a)
	fmt.Println("a =",a)
	func2(&a)
	fmt.Println("a =",a)
}