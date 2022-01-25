package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "d:/test.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		//当读取到EOF最后时，退出读取
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	writer := bufio.NewWriter(file)
	str := "你好呀，北京\n"
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//不加这句，不会刷到硬盘上
	writer.Flush()
}
