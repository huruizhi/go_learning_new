package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.OpenFile("a.txt", os.O_RDONLY, 0644)
	//scanner := bufio.NewScanner(file)
	//for scanner.Scan() {
	//	txt := scanner.Text()
	//	fmt.Println(txt)
	//}
	var b [2048]byte
	reader := bufio.NewReader(file)
	n, _ := reader.Read(b[:])
	fmt.Println(string(b[:n]))

}
