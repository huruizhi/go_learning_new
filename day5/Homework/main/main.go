package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func (p *Student) init(name string, age int, score int) { //需要使用指针
	p.Name = name
	p.Age = age
	p.Score = score

}

func (p Student) get() {
	fmt.Printf("name :%s", p.Name)
}

func main() {
	var stu Student
	stu.init("aaa", 16, 90) //调用方法的时候可以简写指针
	stu.get()
}
