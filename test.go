package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"time"
)

var i, j int = 1, 2

var (
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func add(x, y float64) float64 {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func Sqrt(x float64) float64 {
	z := 1.0
	i := 0
	for int(z*z) != int(x) {
		z -= (z*z - x) / (2 * z)
		i++
	}
	return z
}
func main() {
	v := 123.124312
	fmt.Println("Hello, World")
	fmt.Println("Time now is", time.Now())
	fmt.Println(math.Pi)
	fmt.Println(add(100.122421412, 100))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println(split(7))
	var c, java, python = true, false, "no"
	fmt.Println(i, j, c, java, python)
	fmt.Println(z)
	fmt.Printf("v is of type %T\n", v)
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i

	}
	fmt.Println(sum)
	s := 1
	for s < 1000 {
		s += s
	}
	fmt.Println(s)
	fmt.Println(pow(3, 2, 10))
	fmt.Println(Sqrt(2))
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good evening")
	}
	defer fmt.Println("world")
	fmt.Println("Hello")

}
