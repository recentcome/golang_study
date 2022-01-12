package main

import "fmt"

//小学生、大学生
type Student struct {
	Name  string
	Age   int
	Score float64
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名为%v,年龄为%v，成绩为%v", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score float64) {
	//业务判断
	stu.Score = score
}

type Pupil struct {
	Student
}

func (pupil *Pupil) testing() {
	fmt.Println("小学生正在考试。。。")
}

type Graduate struct {
	Student
}

func (gra *Graduate) testing() {
	fmt.Println("大学生正在考试。。。")
}
