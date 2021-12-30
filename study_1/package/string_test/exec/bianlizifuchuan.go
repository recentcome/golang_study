package main

import (
	"fmt"
	"strconv"
)

func main() {
	str2 := "hello北京"
	str3 := []rune(str2) //切片 字符串遍历同时处理中文的问题
	for i := 0; i < len(str3); i++ {
		fmt.Printf("字符=%c\n", str3[i])
	}

	//字符串转整数
	n, error := strconv.Atoi("1245")
	if error != nil {
		fmt.Println("转换错误", error)
	} else {
		fmt.Println("转换后的值为", n)
	}

	//整数转字符串
	str1 := strconv.Itoa(123)
	fmt.Println(str1)

	//字符串转 []byte
	var bytes = []byte("hello go")
	fmt.Printf("str=%v\n", bytes)

	//[]byte转字符串 str=string([]byte(97,98,99))
	str1 = string([]byte{97, 98, 99})
	fmt.Printf("str is %v", str1)
}
