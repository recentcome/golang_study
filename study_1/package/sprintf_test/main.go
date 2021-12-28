package main

import "fmt"

//基本数据类型转成string
func main() {
	var (
		num1   int     = 99
		num2   float64 = 23.24
		b      bool    = true
		mychar byte    = 'h'
		str1   string  = ""
	)
	str1 = fmt.Sprintf("%d", num1)
	fmt.Printf("%T, str1=%q", str1, str1)

	str1 = fmt.Sprintf("%f", num2)
	fmt.Printf("%T, str1=%q", str1, str1)

	str1 = fmt.Sprintf("%t", b)
	fmt.Printf("%T, str1=%q", str1, str1)

	str1 = fmt.Sprintf("%c", mychar)
	fmt.Printf("%T, str1=%q", str1, str1)
}
