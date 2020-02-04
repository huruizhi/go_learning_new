package main

import "fmt"

// 结构体 用来定义复杂的数据结构
// 结构体是值类型
// go没有class类型， 只有struct类型

func main() {
	// 结构体申明
	type Student struct {
		Name  string
		Age   uint
		score float32
	}
	var stu Student
	stu.Name = "aa"
	stu.Age = 20
	stu.score = 90.6

	var stu1 *Student = &Student{
		Name: "hua",
		Age:  20,
	}

	var stu2 = Student{
		Name: "hh",
		Age:  20,
	}
	fmt.Println(stu1)
	fmt.Println(stu2)
	fmt.Println(stu)

	fmt.Printf("%T\n", stu1)
	fmt.Printf("%p,%p,%p", &stu.Name, &stu.Age, &stu.score)

}
