package main

import (
	"fmt"
	"math/rand"
	"time"
)

//生成10个随机数
func randNum(num int) []int {
	var arr1 []int
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 10; i++ {
		if rand.Intn(100) == 0 {
			rand.Intn(100)
		}
		arr1 = append(arr1, rand.Intn(100))
	}
	fmt.Println(arr1)
	return arr1
}

func maoPao(arr1 []int) []int {
	//这里使用冒泡排序
	//原理，有一个中间数，如果第一个数比第二个数大，则第一个数和中间数交换(temp=arr[0])，
	//第二个数和第一个数交换，arr[0]=arr[1]，中间数和第二个数交换arr[1]=temp
	//第一轮完成后，第二轮从第二个数开始和第三个数进行比较，按照规则
	//排序
	arr2 := arr1
	temp := 0
	for i := 0; i < len(arr2)-1; i++ {
		for j := 0; j < len(arr2)-1-i; j++ {
			if arr2[j] < arr2[j+1] {
				temp = arr2[j]
				arr2[j] = arr2[j+1]
				arr2[j+1] = temp
			}
		}

	}
	return arr2
}

//随机生成10个整数（1_100的范围）保存到数组，并倒序打印以及求平均值、最大值、最大值的下标、并查找里面是否含有55
func main() {
	//随机生成10个整数，并保存到数组
	var arr1 []int
	arr1 = randNum(10)

	////使用递归查找最大值
	//max:=0
	//for i:=0;i<len(arr1)-1;i++{
	//	if arr1[i]>arr1[i+1] {
	//		max = arr1[i]
	//	}
	//}
	//fmt.Println("最大值为",max)

	arr2 := maoPao(arr1)

	fmt.Println("倒序打印的值为：\n", arr2)
	//求平均值
	sum := 0.0
	for i := 0; i < len(arr2); i++ {
		sum += float64(arr2[i])
		//if arr2[i] == 55{
		//	fmt.Println("有55")
		//}
	}
	fmt.Println("总值为", sum)
	fmt.Println("平均值为", sum/(float64(len(arr2))))
}
