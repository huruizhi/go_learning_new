package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var sw sync.WaitGroup

func sender(c chan int, i int)  {
	defer sw.Done()
	for {
		rand.Seed(time.Now().UnixNano())
		second := rand.Intn(5)
		fmt.Printf("生产者%d : 生产了 %d\n",i,second )
		time.Sleep(time.Duration(second) * time.Second)
		c <- second
	}

}

func reciver(c chan int,i int)  {
	defer sw.Done()
	for {
		second :=<- c
		fmt.Printf("消费者%d : 消费了 %d\n",i,second )
		time.Sleep(time.Duration(second) * time.Second)


	}

}

func main()  {
	c := make(chan int, 10)
	for i:=0 ;i<2;i++ {
		sw.Add(1)
		go sender(c,i)
		sw.Add(1)
		go reciver(c,i)
	}
	sw.Wait()

}

