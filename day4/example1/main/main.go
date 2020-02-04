package main

import (
	"errors"
	"fmt"
)

/*
1. close
2. len
3. new 分配值类型，返回指针
4. make 分配引用类型 chan slice map
5. append
6. panic和 recover

recover 捕获错误，保证程序继续执行
panic 定义错误并终止程序
*/

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	b := 0
	a := 100 / b
	fmt.Println(a)
}

func err_test() error {
	return errors.New("function err_test error")
}

func new_make_test() {
	s1 := new([]int) //new 创建了空slice

	s2 := make([]int, 10) //make 进行了初始化

	*s1 = make([]int, 10) // new 之后  还需要make 进行了初始化 。。。

	(*s1)[0] = 100
	fmt.Println(*s1)
	fmt.Println(s2)

}

func main() {

	p := new(int)
	fmt.Println(p)
	fmt.Println(*p)

	var sli []int
	sli = append(sli, 1, 2, 3)
	sli = append(sli, sli...)
	fmt.Println(sli)

	test()

	//err := err_test()
	//if err != nil{
	//	panic(err)
	//}

	new_make_test()

}
