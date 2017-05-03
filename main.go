package main

import (
	"fmt"
	"math"
)

type Demo struct {
	X, Y float64
}

func (v Demo) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func MethodsEx() {
	fmt.Println("+++++++++++++++++++++MethodsEx+++++++++++++++++++++")
	v := Demo{3, 4}
	fmt.Println(v.Abs())
}

func main() {
	MethodsEx()
}
