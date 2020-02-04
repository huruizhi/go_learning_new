package main

import (
	"fmt"
	"strconv"
)

/*
if condition1 {
}

if condition1 {
} else {

}

if condition1 {
} else if condition2 {
} else if condition3 {
}
*/

func ifEx() {
	var input string
	fmt.Scanf("%s", &input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}

/*

switch a {
	case 0:
	case 1:
	default:
}
*/

func switchEx(a int) {
	switch a {
	case 0:
		fmt.Println("hu")
	case 1:
		fmt.Println("rui")
	default:
		fmt.Println("zhi")
	}
}

/*
for range
遍历 数组 slice map chan
*/

func forEx() {
	var str string = "hello world ,中国"
	for i, v := range str {
		fmt.Printf("Index:%d, value:%s, length:%d \n", i, string(v), len([]byte(string(v))))
	}
}

/*
goto
continue
*/
func main() {
	// switchEx(1)
	forEx()

}
