package main

import (
	"fmt"
	a "go_dev/day2/example2/add" // 别名
)

const (
	a1 = iota
	b1
    c1 = "aaa"
	d1 = iota
)

func main()  {
	a.Test()
	fmt.Println("name =",a.Name)
	fmt.Println("age =",a.Age)
	print(c1)
}
