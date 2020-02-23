package main

import "fmt"

func main() {
	ch1 := make(chan int, 100)

	for i := 0; i < 10; i++ {
		ch1 <- i
	}

	close(ch1)

	for x := range ch1 {
		fmt.Println(x)
	}
}
