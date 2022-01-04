package main

import "fmt"

func BinaryFind(arr *[]int, leftIndex int, rightIndex int, findVal int) {
	middle := (leftIndex + rightIndex) / 2
	if leftIndex > rightIndex {
		fmt.Println("没有找到，退出")
		return
	}
	if (*arr)[middle] > findVal {
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Println("找到了", findVal, "下标是", middle)
		return
	}
}

func main() {
	//二分查找的函数
	//思路：
	/*
			arr是个有序数组，并且是从小大大排序,需要查找的值为findVal
			先找到中间的下标middle = (leftIndex+rightIndex)/2
			1、如果arr[middle]>findVal，则在leftIndex和middle-1下标里面继续找
		    2、如果arr[middle]<findVal,则在middle+1和rightIndex里面找
		    3、如果arr[middle]==findVal，找到了，退出
		    4、没有找到的条件是leftIndex>rightIndex
	*/
	arr := []int{1, 2, 3, 4, 5, 6, 1232, 5555}
	fmt.Println(len(arr))
	BinaryFind(&arr, 0, len(arr), 2)
	BinaryFind(&arr, 0, len(arr), 22)
	BinaryFind(&arr, 0, len(arr), 6)
}
