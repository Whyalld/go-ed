package main

import "fmt"

func main() {
	str := new(string)
	// var str string
	*str = "Monday"
	fmt.Println(*str == "Monday")
}
