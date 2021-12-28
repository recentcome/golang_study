package utils

import "fmt"

var (
	Age  int
	Name string
)

func InitTest() {
	Age = 10
	Name = "nicko"
	fmt.Println("in init init()")
}
