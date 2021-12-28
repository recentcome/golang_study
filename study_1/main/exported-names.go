package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	Tobe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	it1    int
	f      float64
	b      bool
	s      string
	i42    int     = 42
	f64    float64 = float64(i42)
	uint1  uint    = uint(f64)
	x, y   int     = 3, 4
	sqrtf  float64 = math.Sqrt(float64(x*x + y*y))
	zint   uint    = uint(sqrtf)
)

const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

//简洁赋值语句 :=  无法在函数外面使用，必须使用var,func等
//k := 3
var i int = 0 //声明一个变量列表 和函数的参数列表一样，类型在最后
var c, python, java = true, false, "no!"

func main() {
	const World = "世界"
	fmt.Println(Pi, World)
	Type_test := 0.867 + 0.5i
	fmt.Printf("type_test is of type %T \n", Type_test)
	fmt.Printf("i42 = %v, f64 = %v, uint1 = %v\n", i42, f64, uint1)
	fmt.Println(sqrtf, zint)
	fmt.Printf("%v %v %v %q\n", it1, f, b, s)
	//简洁赋值语句 :=
	k := 3
	fmt.Println(c, python, java, k)
	fmt.Println(math.Pi)
	fmt.Println(add(1, 2))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(18))
	fmt.Println("i = ", i)
	fmt.Printf("Type: %T Value: %v\n", Tobe, Tobe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(y float64) float64 {
	return y * 0.1
}

func swap(x, y string) (string, string) {
	return y, x
}

func add(x, y int) int {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return //没有参数的return语句返回已经命名的返回值
	//较短的函数可以这么使用
}
