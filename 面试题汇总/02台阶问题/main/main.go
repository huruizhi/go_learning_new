package main

import "fmt"

/*
问题： 有n个台阶，1次可以迈 1步或者 2 步，
问到n个台阶 一共有几种走法

思路：
第N个阶梯的走法 = 第N-2个阶梯的走法+ 1(走两步) + 第N-1个阶梯的走法+ 1(走两步)

方式一： 递归 使用递归的内存消耗太大

*/

func f(n int) int64 {
	if n <= 2 {
		return int64(n)
	}

	var step, tmp, n_1, n_2 int64
	n_1 = 2
	n_2 = 1
	for i := 2; i < n; i++ {
		step = n_1 + n_2 + 2
		tmp = n_1
		n_1 = step
		n_2 = tmp
	}
	return step
}

func main() {
	for i := 1; i < 30; i++ {
		sum := f(i)
		fmt.Println(sum)
	}
}
