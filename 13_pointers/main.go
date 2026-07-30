package main

import "fmt"

func main8() {
	x := 10
	p := &x
	j := new(x)

	*p += 5
	increment(j)

	fmt.Println(*j)
	fmt.Println(x)
	fmt.Println(*p)
}

func increment(value *int) {
	*value += 1
}
