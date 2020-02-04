package main

import "fmt"

/*
map 声明 不消耗内存，需要使用make 初始化

注意 golang 中的 map是无序的。
https://blog.csdn.net/yzf279533105/article/details/81087848
*/

func test() {
	var b = map[string]string{
		"key": "value",
	}
	fmt.Println(b)

	a := make(map[string]string)
	a["k1"] = "efg"
	a["k2"] = "efg"
	a["k1"] = "efg"

	fmt.Println(a)
}

func test2() {
	var map1 = make(map[string]map[string]string)

	map1["r1"] = make(map[string]string)
	map1["r1"]["c1"] = "a"
	map1["r1"]["c2"] = "b"
	map1["r1"]["c3"] = "c"
	map1["r1"]["c4"] = "d"
	map1["r1"]["c5"] = "e"

	map1["r2"] = nil

	fmt.Println(map1)

}

func main() {
	//test()
	test2()
}
