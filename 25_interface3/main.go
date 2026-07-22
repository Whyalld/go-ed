package main

import "fmt"

func main() {
	printType(3)
	printType("интерфейсы это легко")
	printType([]string{"привет"})	

	value := 3
	number, ok := value.(int)
	fmt.Println(number, ok)
}

func printType(value interface{}) {
	if _, ok := value.(int); ok {
		fmt.Println("тип аргумента int")
	} else if _, ok := value.(string); ok {
		fmt.Println("тип аргумента string")
	} else {
		fmt.Println("тип аргумента не int и не string")
	}
}