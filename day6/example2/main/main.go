package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 通过循环读取文件
func readFileByFor() {
	// 打开文件
	file, err := os.Open("./file")
	if err != nil {
		fmt.Printf("open file filed err:%v", err)
		return
	}

	// 关闭文件
	defer file.Close()
	// 读取文件 方式1
	var tmp [120]byte
	for {
		n, err := file.Read(tmp[:])
		if err != nil {

		}
		fmt.Printf("读取了%d个字节:", n)
		fmt.Println(string(tmp[:n]))
		if n < 120 {
			return
		}
	}
}

// 方式2 按行读
func readFileByBuf() {
	file, err := os.Open("./file")
	file.Read()
	defer file.Close()
	if err != nil {
		fmt.Printf("read file filed err:%v", err)
		return
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			println("read file error: %s", err)
			return
		}
		fmt.Print(line)
	}

}

// 方式3 读取整个文件
func readFileByIoUtil() {
	content, err := ioutil.ReadFile("./file")
	if err != nil {
		fmt.Printf("read file filed err:%v", err)
		return
	}
	fmt.Println(string(content[:]))
}

func main() {
	readFileByIoUtil()

}
