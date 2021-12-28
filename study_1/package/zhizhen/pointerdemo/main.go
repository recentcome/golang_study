package main

import "fmt"

func main() {
	// var i int = 10
	// fmt.Println("i的地址为", &i)
	// //1. ptr是一个指针int型变量
	// //2. ptr类型*int
	// //3. ptr本身的值&i
	// var ptr *int = &i
	// fmt.Printf("i的值为%v", *ptr)
	// fmt.Printf("prt的地址为%v", &ptr)

	//var a int = 300
	//var b int = 400
	//var ptr *int = &a
	//*ptr = 100 //a=100
	//ptr = &b
	//*ptr = 200 //b=200
	//fmt.Printf("a=%d,b=%d,*ptr=%d\n", a, b, *ptr)
	////a=100 b=200 *ptr=200
	//fmt.Printf("hero is %s", identifier.HeroName)

	i := 100
	fmt.Println("i的地址值为", &i)
	var ptr *int = &i
	fmt.Println("ptr的值为", *ptr)
}
