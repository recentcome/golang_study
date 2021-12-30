package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
func main() {
	nowTime := time.Now()
	test03()
	nowNewTime := time.Now()
	fmt.Printf("花费的时间为%ds", nowNewTime.Unix()-nowTime.Unix())
}
