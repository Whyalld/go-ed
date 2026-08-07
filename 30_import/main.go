package main

import (
	"fmt"
	"myproject/geometry"
)

func main() {
	circle, _ := geometry.NewCircle(10)
	fmt.Println(geometry.CalculateArea(circle))
}
