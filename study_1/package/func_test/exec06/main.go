package main

import "fmt"

//打印a-z 以及Z-A

func printAscii() {
	for i := 90; i >= 65; i++ {
		fmt.Printf("%c\n", i)
	}

	for j := 97; j <= 122; j++ {
		fmt.Printf("%c\n", j)
	}
}

func main() {
	fmt.Printf("开始打印\n")
	printAscii()
}
