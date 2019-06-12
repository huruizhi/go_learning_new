package main

import "fmt"

func list(num int) {
	for i :=0; i<= num; i++ {
		fmt.Printf("%d+%d=%d\n",i,num-i,num)
	}

}

func main(){
	list(10)
}