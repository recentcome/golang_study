package main

import "fmt"

func main() {
	var slice3 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice3 = append(slice3, 10, 20, 30)
	fmt.Println(slice3)
	//切片追加自己的切片，切片必须跟...
	slice3 = append(slice3, slice3[1:3]...)
	fmt.Println(slice3)

	var slice4 = []int{1, 2, 3, 4, 5}
	var a = make([]int, 1)
	copy(a, slice4) //拷贝到一个长度不够的切片，不会报错，也不会增加切片的长度
	fmt.Println("slice4 is", slice4)
	fmt.Println("a is", a)

	str := "hello@nihao"
	//将hello改变为zello
	str2 := []byte(str)
	str2[0] = 'z'
	str = string(str2)
	fmt.Println(str)

	//将hello改变为‘中ello’
	str3 := []rune(str) //
	str3[0] = '中'
	str = string(str3)
	fmt.Println(str)

}
