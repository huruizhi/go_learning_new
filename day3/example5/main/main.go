package main

import "fmt"

/*
golang 函数特点
1. 不支持重载，一个包内函数不能重名
2. 函数是一等公民， 函数也是一种类型， 一个函数可以赋值给变量
3. 匿名函数
4. 多返回值

变量传递方式：
1. 值传递
2. 引用传递

无论值传递还是引用传递 都是传递变量的副本。
一般来说引用传递 是拷贝的变量的地址，所以引用传递的效率更高。

*/

type add_func func(int, int) int

func add(a int, b int) int {
	return a + b

}

func operator(op add_func, a int, b int) int {
	a = a + 1
	return op(a, b)
}

// 命名返回值
func calc(a int, b int) (sum int, avg int) {
	sum = a + b
	avg = (a + b) / 2
	return
}

// 可变参数
func add2(a int, arg ...int) int {
	sum := a
	for _, v := range arg {
		sum = sum + v
	}
	return sum
}

/*
defer ：
1. 在函数返回时，执行defer 语句。 因此可以做资源清理。
2. 多个defer 语句 按先进后出的方式执行
3. defer 语句的变量在defer 语句声明时就确定了。

defer 用途
1. 关闭文件句柄
file : = open(file)
defer file.close()
2. 锁资源释放
mc.Lock()
defer mc.Unlock()
*/

func defer_ex1() {
	var i = 0
	defer fmt.Println(i)
	i = 10
	fmt.Println(i)
	return
}

func defer_ex2() {
	var num = 0
	for i := 0; i < 10; i++ {
		defer fmt.Println(num)
		num++
	}
	return
}

//匿名函数
func test(a, b int) int {
	//定义匿名函数 并调用
	result := func(a, b int) int {
		return a + b
	}(a, b)
	return result
}

func main() {
	f := add
	fmt.Println(f(2, 3))
	fmt.Println(operator(f, 3, 3))
	fmt.Println(calc(2, 4))

	// 可变参数
	sum1 := add2(1, 2, 3, 4, 5, 6)
	fmt.Println(sum1)
	num_list := []int{1, 2, 3, 4, 5, 7}
	sum2 := add2(6, num_list...)
	fmt.Println(sum2)

	//defer
	defer_ex1()
	defer_ex2()

}
