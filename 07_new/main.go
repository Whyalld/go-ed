package main

import "fmt"

func main() {
	str := new(string)
	*str = "Monday"
	fmt.Println(*str)
}
