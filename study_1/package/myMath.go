package main

import (
	"fmt"
	"math"
	"runtime"
	// "time"
)

func add(x, y int) int {
	return x + y
}

func sqrt(x float64) float64 {
	// 	if x < 0 {
	// 		return sqrt(-x) + "i"
	// 	}
	// 	return fmt.Sprint(math.Sqrt(x))
	// }
	z := float64(1)
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	// count := 10
	// sum := 20
	// for jae := 0; jae < count; jae++ {
	// 	sum += jae
	// }
	// fmt.Println(sum)
	// fmt.Println(sqrt(2), sqrt(-4))
	// fmt.Println(pow(3, 2, 10))
	// fmt.Println(pow(3, 3, 20))
	// for {

	// 	fmt.Println(time.Now())
	// }
	fmt.Println(sqrt(344))
	fmt.Println("MQ", math.Sqrt(344))
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}

}
