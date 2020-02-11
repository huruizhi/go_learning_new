package main

import (
	"fmt"
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

func main() {
	writeFile()
}
