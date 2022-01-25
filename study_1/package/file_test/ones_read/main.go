package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("d:/pingall.log")
	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	fmt.Println(string(file))

}
