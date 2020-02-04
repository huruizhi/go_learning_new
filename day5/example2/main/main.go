package main

import "fmt"

type student struct {
	Name string
	Age  int
	next *student
}

func listStruct(stu student) {
	var p *student = &stu
	println()
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func tailINsert() {
	var headStu *student
	var preStu *student = &student{}
	for i := 0; i < 10; i++ {
		age := 18 + i
		var stu = student{
			Age:  age,
			Name: fmt.Sprintf("stu%d", i),
		}
		preStu.next = &stu
		if i == 0 {
			headStu = &stu
			preStu = headStu
		} else {
			preStu = &stu
		}
	}
	listStruct(*headStu)
}

func headInsert(headStu **student) {

	//fmt.Printf("%p\n", &(**headStu).Name)

	for i := 0; i < 10; i++ {
		age := 18 + i
		var newStu = student{
			Age:  age,
			Name: fmt.Sprintf("stu%d", i),
		}
		newStu.next = *headStu
		*headStu = &newStu
	}
}

// 这个函数不会改变外面headStu 的值
// 因为 赋值 headStu = &newStu 是直接修改指针形参的值，而非改变指针指向的数据。
func headInsert2(headStu *student) {
	//fmt.Printf("%p\n", &(*headStu).Name)

	for i := 0; i < 10; i++ {
		age := 18 + i
		var newStu = student{
			Age:  age,
			Name: fmt.Sprintf("stu%d", i),
		}
		newStu.next = headStu
		headStu = &newStu
	}
}

func test(i *int) {
	fmt.Printf("%p\n", i)
	*i = 10
	fmt.Printf("%p\n", i)
}

func main() {

	//tailINsert()
	var headStu *student = &student{
		Name: "tail",
		Age:  10,
	}
	//fmt.Printf("%p\n", &(*headStu).Name)
	headInsert(&headStu)
	headInsert2(headStu)

	listStruct(*headStu)

	// 指针测试
	//i := 0
	//fmt.Printf("%p\n", &i)
	//test(&i)
	//
	//fmt.Println(i)
}
