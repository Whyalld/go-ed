package main

import "fmt"
import "os"

func main() {
	// Sprint
	var verb string = "hello"
	fmt.Printf("%v my friend\n", verb)
	age := 18
	name := "Jack"
	var data string = fmt.Sprintf("Name: %s, age: %d", name, age)
	data2 := fmt.Sprint(name, age)
	print(data, "\n", data2, "\n")

	// Fprint
	file, _ := os.Create("06_fmt/log.txt")
	defer file.Close()

	fmt.Fprint(file, "baby")
	fmt.Fprintln(os.Stderr, "chiken burger")

}
