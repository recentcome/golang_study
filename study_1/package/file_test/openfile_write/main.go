package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "d:/test.txt"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	str1 := "nihao baibai!\n"
	for i := 0; i < 5; i++ {
		writer.WriteString(str1)
	}
	writer.Flush()
}
