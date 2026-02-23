package functions

import (
	"fmt"
	"math"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func compute(fn func(float64, float64) float64, x, y float64) float64 {
	return fn(x, y)
}

type Vertex struct {
	x, y float64
}

// value reciver
func (v Vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// we can declare methods with pointer receivers
// but wat is this ?
// if we use value reciever all the change we do in method will not reflect on instance of struct
// it will on the copy

// in pointer reciver whatever changes we make will reflects on instance of that type struct

// pointer reciver
func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Run() {
	fmt.Println("--------Go Functions--------")
	// functions are the first class citizens , so we can assign them to any variables

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 5))
	fmt.Println(compute(hypot, 6, 6))

	fmt.Println("--------Go Closures--------")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	fmt.Println("--------Go Methods--------")
	// Go does not have classes. However, we can define methods on types.
	// A method is a function with a special receiver argument.
	v := Vertex{4, 4}
	fmt.Println(v.abs())

	// we can declare a method on non-struct types, too.
	// You can only declare a method with a receiver whose type is defined in the same package as the method.

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(5)
	fmt.Println(v)
}
