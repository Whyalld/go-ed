package main

import (
	"fmt"
	"math"
)

func Compose(
	f, g func(float64) float64,
	) func(float64) float64 {
	return func(x float64) float64 {
		return f(g(x))
	}
}

func main() {
	result := Compose(math.Sin, math.Cos)(0.5)
	fmt.Println(result)
}
