package main

import (
	"fmt"
	utils "gocode/study_1/package/func_test/init"
)

func swap(n1, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

func main() {
	a := 20
	b := 10
	swap(&a, &b)
	fmt.Printf("a=%v,b=%v", a, b)
	utils.Name = "nihao"
	utils.Age = 10
	fmt.Println("age is ", utils.Age, "name is ", utils.Name)
}
