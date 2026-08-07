package main

import "fmt"

func main() {
	defer fmt.Println("Выпольнится в конце")

	fmt.Println("выполнится сейчас")
}
