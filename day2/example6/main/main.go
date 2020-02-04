package main

import "fmt"

func reverse(str string) string {
	var str_reverse string
	for i := 0; i < len(str); i++ {
		str_reverse = str_reverse + fmt.Sprintf("%c", str[len(str)-1-i])
	}
	return str_reverse
}

func reverse1(str string) string {
	var result []byte
	tmp := []byte(str)
	for i := 0; i < len(str); i++ {
		result = append(result, tmp[len(str)-1-i])
	}
	return string(result)

}

func main() {
	str1 := "hello"
	str2 := "world"
	str3 := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str3)

	var substr = str3[5:]
	fmt.Println(substr)

	str_r1 := reverse(str3)
	fmt.Println(str_r1)

	str_r2 := reverse1(str3)
	fmt.Println(str_r2)

}
