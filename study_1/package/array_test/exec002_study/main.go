package main

import (
	"fmt"
)

//已知有个排序好的数据，要求插入一个元素，最后打印该数组，顺序依然是升序
func insertArr(arr *[]int, num int, index int) []int {
	//num:插入的元素值  index：要插入的位置
	//思路：需要使用切片，
	slice1 := make([]int, len(*arr)+1)
	slice1 = *arr
	fmt.Println(slice1[index])
	slice1 = append(slice1[:index], append([]int{num}, slice1[index:]...)...)
	return slice1
}
func main() {
	var arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 := insertArr(&arr1, 10, 3)
	fmt.Println(arr2)
}
