package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrDivByZero = errors.New("division by zero")

func do(x float64, y float64) (float64, error) {
	if y == 0 { return 0, ErrDivByZero } 
	return x / y, nil
}

func main() {
	var err error
	fmt.Println(err)
	
	v, err := do(1, 2)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(v)
}