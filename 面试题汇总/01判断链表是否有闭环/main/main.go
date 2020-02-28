package main

import (
	"errors"
	"fmt"
)

/*
判断链表是否有闭环
*/

type a struct {
	n    int
	next *a
}

func createCycleList() *a {
	// var cyclePoint *a
	var startA = &a{n: 0} // 定义链表开头
	var endA = startA     // 定义链表结尾，结尾初始指向开头
	for i := 1; i < 10; i++ {

		// 创建链表节点
		as := a{
			n: i,
		}
		// 链表最后节点指向新创建的节点
		endA.next = &as

		// 将新建节点置位最后节点
		endA = &as

		// 定义环路位置
		// if i == 3 {
		// 	cyclePoint = &as
		// }
	}
	// 创造一个环路
	// endA.next = cyclePoint
	return startA
}

func getNextListPoint(as a) (an *a, err error) {
	an = as.next
	if an == (*a)(nil) {
		err = errors.New("end of this Line!")
		fmt.Println(as)
		return
	}
	return
}

func reverseList(head *a) *a {
	var Node, nextNode, nextTmp *a
	if head.next == (*a)(nil) {
		return head
	}
	Node = head
	nextNode = Node.next
	for {
		// 1.将next_node的next 保存到tmp
		nextTmp = nextNode.next

		// 2. 将nextNode的next 指向Node节点
		nextNode.next = Node

		// 3.将nextNode 赋值给 Node
		Node = nextNode

		// 4.将nextTmp赋值给 Node
		nextNode = nextTmp

		// 5。判断nextNode 是不是空
		if nextNode == (*a)(nil) {
			// 将head的next 值为Nil
			head.next = nil
			return Node
		}
	}
}

func reverseList2(head *a) *a {
	Node := head
	nextNode := Node.next
	head.next = nil
	return f1(Node, nextNode)
}

func f1(node *a, nodeNext *a) *a {
	nextTmp := nodeNext.next
	nodeNext.next = node
	if nextTmp == (*a)(nil) {
		return nodeNext
	} else {
		f1(nodeNext, nextTmp)
	}
	return nodeNext
}

func main() {
	listStart := createCycleList() //获取开始位置 并赋给 步长1 与 步长2
	// stepONe := listStart
	// stepTwo := listStart
	// var err error
	// LOOP:
	// 	for {
	// 		stepONe, err = getNextListPoint(*stepONe)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			break
	// 		}
	// 		for i := 0; i < 2; i++ {
	// 			stepTwo, err = getNextListPoint(*stepTwo)
	// 			if err != nil {
	// 				fmt.Println(err)
	// 				break LOOP
	// 			}
	// 		}
	// 		if stepONe == stepTwo {
	// 			fmt.Println("链表存在闭环")
	// 			break
	// 		}
	// 	}
	// 	reverseList(listStart)

	list := reverseList2(listStart)
	for {
		fmt.Println(list.n)
		if list.next == (*a)(nil) {
			break
		}
		list = list.next
	}
}
