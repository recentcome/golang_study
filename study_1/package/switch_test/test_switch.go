package main

import (
	"fmt"
	"time"
)

func switch_test() {
	t := time.Now()
	// t1 := time.Now().Format("1970-01-01 00:00:00")
	// fmt.Println("now time is ", t1)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	defer fmt.Println("first?") //将函数推迟到外层函数返回之后执行 参数会先计算，但是最后才调用。
	fmt.Println("when's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	switch_test()

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) //推迟的函数会压栈，后进先出。 先打印9，最后打印0

	}
	fmt.Println("done")

}
