package main

import (
	"fmt"
	"math/rand"
	"time"
)

//随机生成10个整数（1_100的范围）保存到数组，并倒序打印以及求平均值、最大值、最大值的下标、并查找里面是否含有55
func main() {
	var arr1 []int
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 10; i++ {
		if rand.Intn(100) == 0 {
			rand.Intn(100)
		}
		arr1 = append(arr1, rand.Intn(100))
	}
	fmt.Println(arr1)
	//这里使用冒泡排序
	//原理，有一个中间数，如果第一个数比第二个数大，则第一个数和中间数交换(temp=arr[0])，
	//第二个数和第一个数交换，arr[0]=arr[1]，中间数和第二个数交换arr[1]=temp
	//第一轮完成后，第二轮从第二个数开始和第三个数进行比较，按照规则
	for j := 0; j < len(arr1); j++ {
		fmt.Printf("")
	}

}
