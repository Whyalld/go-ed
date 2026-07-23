package main

import "fmt"

func main() {
	printType(3)
	printType("интерфейсы это легко")
	printType([]string{"привет"})	

	var value any = 3
	number, ok := value.(int)
	fmt.Println(number, ok)
}

func printType1(value interface{}) {
	if _, ok := value.(int); ok {
		fmt.Println("тип аргумента int")
	} else if _, ok := value.(string); ok {
		fmt.Println("тип аргумента string")
	} else {
		fmt.Println("тип аргумента не int и не string")
	}
}

func printType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("тип аргумента int")
	case string:
		fmt.Println("тип аргумента string")
	default:
		fmt.Println("неизвестный тип аргумента")
	}
}