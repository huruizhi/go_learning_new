package main

import "fmt"

/*完数 */
func wanshu(num int) bool {
	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum = sum + i
		}
	}
	if num == sum {
		return true
	} else {
		return false
	}
}

/* 回文 */
func huiwen(str string) bool {
	sli := []byte(str)
	for i := 0; i <= len(sli)/2; i++ {
		if sli[i] != sli[len(sli)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	for i := 1; i <= 1000; i++ {
		if wanshu(i) {
			fmt.Println(i)
		}
	}

	if huiwen("abcba") {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
