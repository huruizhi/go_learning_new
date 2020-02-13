package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeFile() {
	file, err := os.OpenFile("./NewFile.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		fmt.Printf("Write file error! %v \n", err)
	}
	s := `123
456
789
`
	file.Write([]byte(s))
	file.WriteString("aaabbbvvvccc\n")

}

func writeFileByBufIo() {
	file, err := os.OpenFile("./NewFile.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		fmt.Printf("Write file error! %v \n", err)
	}
	wr := bufio.NewWriter(file)
	wr.WriteString("qweqwe123123")
	// 将缓存内容写入磁盘
	wr.Flush()
}

func writeFileByIoUtil() {
	err := ioutil.WriteFile("NewFile.txt", []byte("123123123"), 0644)
	if err != nil {
		fmt.Printf("Write file error! %v \n", err)
	}
}

func main() {
	writeFileByIoUtil()
}
