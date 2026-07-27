package main

import "fmt"

// Глобальные переменные (нельзя через :=)
var King string = "me"

// Константы (могут быть как локальными, та и глобальными)
const pi float32 = 3.1415

func main() {
	// Локальные переменные
	var name string = "Said"
	var count int = 10
	isReady := true
	count += 1
	name = "Said"

	fmt.Println(count)
	fmt.Println(name)
	fmt.Println(isReady)

	var b, k, j, p = true, 238, "Cool", 3.14
	fmt.Println(b, k, j, p)

	var pool, loop, cool int
	pool, loop, cool = 32, 123, 903292310
	fmt.Println(pool, loop, cool)

	distance, John, nether := 91832.231, "Smith", "Biom"
	fmt.Println(distance, John, nether)

	// Naming
	// С большой буквы публичные, с маленькой приватные
	// локальные переменные не могут быть публичными
	// публичные / привтные - это понятие отвечает за видимость переменных для других пактов
	// локальные / глобальные - объявлены в функции, объявлены вне функции

}
