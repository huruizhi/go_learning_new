package main

import (
	"fmt"
)

func swap_1(a *int, b *int){
	temp := *a
	*a = *b
	*b = temp
}

func swap_2(a *int, b *int) (*int,*int){
	temp := a
	a = b
	b = temp
	return a,b
}

func main() {
	a := 10
	b := 20
	swap_1(&a, &b)
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	a1 := 100
	b1 := 200
	a_n, b_n := swap_2(&a1, &b1)
	fmt.Println("a =", *a_n)
	fmt.Println("b =", *b_n)

}