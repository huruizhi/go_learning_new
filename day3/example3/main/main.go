package main

import "fmt"

/*
指针保存的是地址， 指针指向的地址保存着变量的值
使用& 获取变量的指针
使用* 获取指针存放的值
*/

func changValue(p *string) {
	*p = "changValue"
}

func main() {
	var a int = 2
	fmt.Println(&a)

	var s string = "hurz"
	changValue(&s)
	fmt.Println(s)

}
