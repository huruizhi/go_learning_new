package main

import (
	"fmt"
	"sort"
)

/*
1. slice
初始化方式
	1. 切片
	2. make

slice 内置函数：
	1. append 追加会扩容
	2. copy  拷贝不会扩容

2. string string底层就是一个byte 数组，也可以进行切片
字符串 byte数组 不能直接进行修改，可以赋值给一个byte数组后进行修改。

3. 排序 使用 sort 包
*/

func test() {
	// 方式一 由数组创建
	var list [5]int = [...]int{1, 2, 3, 4, 5}
	var slice []int

	slice = list[1:4]
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))

	slice = slice[0:1]
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))

}

func test1() {
	var slice = make([]int, 0, 2)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))

	for i := 0; i < 17; i++ {
		slice = append(slice, i) //append
		fmt.Printf("%d cap:%d len:%d \n", i, cap(slice), len(slice))
	}

	numbers1 := make([]int, len(slice), (cap(slice))*2)
	copy(numbers1, slice)

	fmt.Println(numbers1)
	fmt.Println(cap(numbers1))
	fmt.Println(len(numbers1))

}

//
func test2() {
	var str = "我 hello world!"
	substr := str[4:8]
	fmt.Println(substr)

	// 不要用这个 str_byte_list := []byte(str)
	str_byte_list := []rune(str)
	str_byte_list[0] = '你'
	str_byte_list[2] = 'n'
	fmt.Println(string(str_byte_list))

}

func sortTest() {
	num_list := []int{1, 2, 3, 8, 5, 2, 0, 6, 10}
	sort.Ints(num_list)
	fmt.Println(num_list)

	str_list := []string{"abc", "bhj", "asdf", "onu"}
	sort.Strings(str_list)
	fmt.Println(str_list)

	index := sort.SearchInts(num_list, 2)
	fmt.Println(index)

}

func main() {
	// test()
	// test1()
	// test2()
	sortTest()
}
