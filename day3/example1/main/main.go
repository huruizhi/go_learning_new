package main

import (
	"fmt"
	"strconv"
	"strings"
)

func urlProccess(url string) string {
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}
	return url
}

func pathProccess(path string) string {
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}
	return path
}

func string_test(str string, char string) {
	index := strings.Index(str, char)
	last_index := strings.LastIndex(str, char)

	fmt.Println(index)
	fmt.Println(last_index)

}

func main() {
	//var (
	//	url string
	//	path string
	//)
	//fmt.Scanf("%s%s", &url, &path)
	//
	//url = urlProccess(url)
	//path = pathProccess(path)
	//
	//fmt.Println(url)
	//fmt.Println(path)

	index := strings.Index("baba", "a")
	fmt.Println(index)

	last_index := strings.LastIndex("baba", "a")
	fmt.Println(last_index)

	str_rep1 := strings.Replace("aabbaadd", "aa", "hu", -1)
	fmt.Println(str_rep1)
	str_rep2 := strings.Replace("aabbaadd", "aa", "hu", 1)
	fmt.Println(str_rep2)

	str_repe := strings.Repeat("abc", 3)
	fmt.Println(str_repe)

	str_trim := strings.Trim("abababa", "a")
	fmt.Println(str_trim)

	str_sl := strings.Fields("abc def ghi")
	fmt.Println(str_sl)

	str_sl2 := strings.Split("abc:def:ghi", ":")
	fmt.Println(str_sl2)

	list := []string{"abc", "def", "ghi"}
	str_join := strings.Join(list, ":")
	fmt.Println(str_join)

	str2 := strconv.Itoa(123)
	fmt.Println(str2)

	int2, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(int2)
	}

}
