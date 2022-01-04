package main

import "fmt"

func main() {
	var str1 [26]byte
	str1[0] = 'A'
	for i := 1; i < 26; i++ {
		str1[i] = str1[i-1] + 1
	}
	for i, v := range str1 {
		fmt.Printf("index is %v, value is %c\n", i, v)
	}
}
