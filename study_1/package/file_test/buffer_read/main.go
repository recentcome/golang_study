package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("d:/pingall.log")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	//当文件使用完成后，需要及时关闭file
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到一个换行就继续
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Print("读取文件完成")
}
