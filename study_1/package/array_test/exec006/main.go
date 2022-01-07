package main

import "fmt"

//已知数组arr[10]string里面保存了10个元素，现要查找"AA"在其中是否存在，打印提示。如果有多个"AA"，找到对应的下标
func main() {
	var arr = [10]string{"123", "AA", "1234", "AAAA", "AA", "000", "SDGG", "sdfsdfd", "123213", "sdfds"}
	for i := 0; i < len(arr); i++ {
		if arr[i] == "AA" {
			fmt.Println("有AA,下标为", i)
		}
	}
}
