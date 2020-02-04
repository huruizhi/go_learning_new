package main

import "fmt"

/*
1. 数组：同一种数据类型，长度固定，值类型，传递给函数为副本
*/

func test1(list *[5]int) {
	(*list)[0] = 1000
}

func test2(list [5]int) {
	list[0] = 2000
}

func main() {
	var (
		list  [5]int
		list1 = [...]int{1, 2, 3}
		list2 = [5]string{3: "hello", 4: "world"}
		// 二维数组
		list3 = [...][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14}}
	)

	for k, v := range list {
		fmt.Println(k, v)
	}

	test1(&list)
	fmt.Println(list)
	test2(list)
	fmt.Println(list)

	fmt.Println(list1)
	fmt.Println(list2)

	for row, v := range list3 {
		for col, v1 := range v {
			fmt.Printf("(%d,%d):%d", row, col, v1)
		}
		fmt.Println("")
	}

}
