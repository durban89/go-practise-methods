package main

import (
	"fmt"
	"math"
)

type Demo struct {
	X, Y float64
}

type MyFloat float64

func (v Demo) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Fabs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func MethodsEx() {
	fmt.Println("+++++++++++++++++++++MethodsEx+++++++++++++++++++++")
	v := Demo{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Fabs())
}

func main() {
	MethodsEx()
}
