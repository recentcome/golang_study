package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Stu struct {
	Name  string
	Age   int
	Score float32
}

//接口编程，最佳实践，切片排序
/*
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
*/
//定义一个结构体的切片
type Student []Stu

func (stu Student) Len() int {
	return len(stu)
}

//使用升序还是降序
func (stu Student) Less(i, j int) bool {
	return stu[i].Name < stu[j].Name
}

//交换
func (stu Student) Swap(i, j int) {
	stu[i].Name, stu[j].Name = stu[j].Name, stu[i].Name
}

func main() {
	fmt.Println("开始")
	//新增10个学生
	var stu = Stu{}
	var stuSlice = Student{}
	for i := 0; i < 10; i++ {
		stu.Name = fmt.Sprintf("Hero%d\n", rand.Intn(100))
		stu.Age = rand.Intn(100)
		stu.Score = float32(rand.Intn(100))
		stuSlice = append(stuSlice, stu)
	}
	fmt.Println("排序前的学生为\n", stuSlice)

	sort.Sort(stuSlice)
	fmt.Println("排序后的学生为\n", stuSlice)
	//给一个结构体进行排序，使用sort包提供的sort.Sort()，需要实现Interface的三个方法

}
