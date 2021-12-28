package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string //定义string类型数组
	a[0] = "hello"
	a[1] = "world"
	fmt.Println("dfg", a)

	primes := [6]int{2, 3, 4, 6, 7, 8} //定义int型数组
	//对数组切片
	var s []int = primes[1:3]
	fmt.Println(primes)
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	ac := names[0:2]
	b := names[1:3]
	b[0] = "xxx" //更改切片的元素会修改底层数组对应的元素
	fmt.Println(ac, b, names)

	q := []int{2, 3, 3, 4, 5, 6, 76}
	fmt.Println(q)
	r := []bool{true, false, true, false, false}
	fmt.Println(r)

	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s2)
	fmt.Println(s2[2].b)
	// 截取切片使其长度为0
	q = q[:0]
	printSlice(q)
	//扩展其长度
	q = q[:4]
	printSlice(q)
	//舍弃前两个值
	q = q[2:]
	printSlice(q)
	//扩展到超过原有值，会报错panic slice bounds out of range [:10] with capacity 5
	// q = q[:10]
	// printSlice(q)
	var s3 []int
	fmt.Println(s3, len(s3), cap(s3))
	if s3 == nil {
		fmt.Println("nil!")
	}
	//使用make创建切片，动态创建数组
	make_1 := make([]int, 5) // len(make_1) = 5
	printSlice2("make_1", make_1)

	//要指定容量，需要传入第三个参数
	make_2 := make([]int, 0, 5) //len(make_2)=0, cap(b)=5
	printSlice2("make_2", make_2)

	make_3 := make_2[:2]
	printSlice2("make_3", make_3)

	make_4 := make_3[2:5]
	printSlice2("make_4", make_4)

	//创建一个井字板
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	//两个玩家轮流打上X和O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	//添加一个空切片
	var null_slice []int
	printSlice(null_slice)
	null_slice = append(null_slice, 1)
	printSlice(null_slice)
	null_slice = append(null_slice, 3, 4, 6)
	printSlice(null_slice)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
