package main

import (
	"fmt"
)

func main() {
	var classNum = 2
	var studentNum = 5
	var score = 0.0
	var classTotalScore = 0.0
	var scorePass = 0
	for i := 1; i <= classNum; i++ {
		var singleTotal = 0.0
		for j := 1; j <= studentNum; j++ {
			fmt.Printf("请输入第%d个班的第%d个学生的成绩\n", i, j)
			fmt.Scanln(&score)
			singleTotal += score
			if score >= 60 {
				scorePass++
			}
		}
		fmt.Printf("第%d个班级的平均分为%f\n", i, singleTotal/float64(studentNum))
		classTotalScore += singleTotal
	}
	fmt.Printf("所有班级的平均分为%f\n", classTotalScore/float64(classNum*studentNum))
	fmt.Printf("及格人数有%d个", scorePass)
}
