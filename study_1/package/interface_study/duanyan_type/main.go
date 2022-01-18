package main

import "fmt"

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case float32:
			fmt.Printf("第%v个参数类型为flout32,值为%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数类型为flout64,值为%v\n", index, x)
		case bool:
			fmt.Printf("第%v个参数类型为bool,值为%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数类型为string,值为%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数类型为整型,值为%v\n", index, x)
		case Student:
			fmt.Printf("第%v个参数类型为Student,值为%v\n", index, x)
		case *Student:
			fmt.Printf("第%v个参数类型为*Student,值为%v\n", index, x)
		default:
			fmt.Printf("第%v个参数类型为未知,值为%v\n", index, x)

		}
	}
}

type Student struct {
}

func main() {
	var (
		n1 = 1
		n2 = 1.1
		n3 = "string"
		n4 = true
		n5 = 3000000000000000
	)
	student := Student{}
	student1 := &Student{}
	TypeJudge(n1, n2, n3, n4, n5, student, student1)
}
