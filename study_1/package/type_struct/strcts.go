package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	var (
		v1 = Vertex{X: 1}
		v2 = Vertex{}
		p  = &Vertex{1, 2}
	)
	fmt.Println(v1, v2, *p)
}
