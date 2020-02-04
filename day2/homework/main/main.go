package main

import (
	"fmt"
	"strconv"
)

/*
100-999之间的水仙花数
*/

func narcissistic(num int) bool {
	str := fmt.Sprintf("%d", num)
	var sum int
	for i := 0; i < len(str); i++ {
		j := string(str[len(str)-1-i])
		n, _ := strconv.Atoi(j)
		sum = sum + n*n*n
	}
	// fmt.Println(num," ",sum)

	if num == sum {
		return true
	} else {
		return false
	}

}

func order(num int) int {
	total := 0
	for i := 1; i <= num; i++ {
		product := 1
		for j := 1; j <= i; j++ {
			product = product * j
		}
		total = total + product
	}
	return total
}

func main() {
	for i := 100; i < 1000; i++ {
		if narcissistic(i) {
			fmt.Println(i)
		}
	}

	fmt.Println(order(3))
}
