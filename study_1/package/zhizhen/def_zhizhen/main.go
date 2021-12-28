package main

import (
	"fmt"
	"math"
)

type Strcut1 struct {
	X, Y float64
}

func Abs(v Strcut1) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Strcut1) NewAbs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Strcut1) ScaleOld(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//对原值乘以f的倍数
func Scale(v *Strcut1, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Strcut1{3, 4}
	// Scale(&v, 10)
	v.ScaleOld(10)
	var v1 = Strcut1{3, 4}
	// Scale(v1, 5) //编译错误！
	Scale(&v1, 10)
	fmt.Println(Abs(v))
	fmt.Println(Abs(v1))

	v3 := Strcut1{3, 4}
	fmt.Printf("Before scaling: %+v, abs: %v\n", v3, v3.NewAbs())
	v3.ScaleOld(5)
	fmt.Printf("After scaling: %+v, abs: %v\n", v3, v3.NewAbs())
}
