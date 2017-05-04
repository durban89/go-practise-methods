package main

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"
)

type Demo struct {
	X, Y float64
}

type MyFloat float64

func (v Demo) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Demo) Scale(x float64) {
	v.X = v.X * x
	v.Y = v.Y * x
}

func ScaleFunc(v *Demo, x float64) {
	v.X = v.X * x
	v.Y = v.Y * x
}

func (f MyFloat) Fabs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func MethodsEx() {
	fmt.Println("+++++++++++++++++++++MethodsEx+++++++++++++++++++++")
	v := Demo{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Fabs())
}

func MethodsAndPointerIndirectionEx() {
	fmt.Println("+++++++++++++++++++++MethodsAndPointerIndirectionEx+++++++++++++++++++++")
	v := &Demo{3, 4}
	ScaleFunc(v, 10)
	fmt.Println(v)
	fmt.Println(v.Abs())
}

type Abser interface {
	Abs() float64
}

func InterfacesEx() {
	fmt.Println("+++++++++++++++++++++InterfacesEx+++++++++++++++++++++")
	var a Abser

	f := MyFloat(-math.Sqrt2)
	v := Demo{3, 4}

	a = f
	a = &v

	fmt.Println(a.Abs())
}

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func (f F) M() {
	fmt.Println(f)
}
func InterfaceValueEx() {
	fmt.Println("+++++++++++++++++++++InterfaceValueEx+++++++++++++++++++++")
	var i, y I

	i = &T{"Hello"}
	describe(i)
	i.M()

	var t *T
	y = t
	describe(y)
	y.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T) \n", i, i)
}

func InterfaceEmptyEx() {
	fmt.Println("+++++++++++++++++++++InterfaceEmptyEx+++++++++++++++++++++")
	var i interface{}
	describeWithEmptyInterface(i)

	i = 42
	describeWithEmptyInterface(i)

	i = "hello"
	describeWithEmptyInterface(i)
}

func describeWithEmptyInterface(i interface{}) {
	fmt.Printf("(%v, %T) \n", i, i)
}

func doInterfaceSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func InterfaceTypeSwitchesEx() {
	fmt.Println("+++++++++++++++++++++InterfaceTypeSwitchesEx+++++++++++++++++++++")
	doInterfaceSwitch(21)
	doInterfaceSwitch("string")
	doInterfaceSwitch(true)
}

type Person struct {
	Age  int
	Name string
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func InterfaceStringersEx() {
	fmt.Println("+++++++++++++++++++++InterfaceStringersEx+++++++++++++++++++++")
	x := Person{21, "durban1"}
	y := Person{34, "durban2"}
	fmt.Println(x, y)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"It's Work",
	}
}

func ErrorsEx() {
	fmt.Println("+++++++++++++++++++++ErrorsEx+++++++++++++++++++++")
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func ReaderEx() {
	fmt.Println("+++++++++++++++++++++ReaderEx+++++++++++++++++++++")
	r := strings.NewReader("Hello World")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v \n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func main() {
	MethodsEx()
	MethodsAndPointerIndirectionEx()
	InterfacesEx()
	InterfaceValueEx()
	InterfaceEmptyEx()
	InterfaceTypeSwitchesEx()
	InterfaceStringersEx()
	ErrorsEx()
	ReaderEx()
}
