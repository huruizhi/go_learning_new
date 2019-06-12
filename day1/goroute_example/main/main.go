package main

import (
	"fmt"
	"go_dev/day1/goroute_example/goroute_test"
)

func main()  {
	ch := make(chan int,1)
	go goroute_test.Add(10,5,ch)
	c :=<- ch
	fmt.Println(c)

}
