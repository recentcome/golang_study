package main

import "fmt"

func main() {
	//试保存1 3 5 7 9五个奇数到数组，并倒序打印
	var str2 [5]int
	for i := 0; i < len(str2); i++ {
		fmt.Printf("请输入数组的第%d个元素的数值", i+1)
		fmt.Scanln(&str2[i])
	}

	//使用冒泡排序。
	//思路：第一轮：第一个数和第二个数比较，若第一个数大于第二个数，则交换位置。
	// 第二个数和第三个数比较，若第二个数大于第三个数，则交换位置。比到最后一个位置。将最大的数找到，比较次数：数组个数-1
	//第二轮：
	for j := 0; j < len(str2)-1; j++ {
		temp := 0
		for i := 0; i < len(str2)-j-1; i++ {
			if str2[i] < str2[i+1] {
				temp = str2[i]
				str2[i] = str2[i+1]
				str2[i+1] = temp
			}
		}
	}
	fmt.Println(str2)

}
