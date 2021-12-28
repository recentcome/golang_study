package main

import (
	"fmt"
	"math"
)

type Struct1 struct {
	X, Y float64
}

func (v Struct1) Abc() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//上下两个方法是同样的
func Acd(v Struct1) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (m MyFloat) AAC() float64 {
	if m < 0 {
		return float64(-m)
	}
	return float64(m)
}

func main() {
	v := Struct1{3, 4}
	fmt.Println(v.Abc())
	fmt.Println(Acd(v))
	m := MyFloat(-math.Sqrt2)
	fmt.Println(m.AAC())
}
