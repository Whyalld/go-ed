package main

import "fmt"

// any равнозначем interface{}
// эти 2 функции равнозначны
func printValue(value interface{}) {
	fmt.Println(value)
}

func printValue2(value any) {
	fmt.Println(value)
}

func main() {
	var value any = 5
	// fmt.Println(value + 5) // Ошибка - value имеет тип any (interface{})

	number, ok := value.(int)

	if ok {
		fmt.Println(number + 5) // 10
	}
	// это называется type assertion
}

// функции double и double2 равнозначны, во второй показана инициализация переменный прямо в условии
func double(value any) {
	number, ok := value.(int)

	if ok {
		fmt.Println(number * 2)
	} else {
		fmt.Println("value не является числом")
	}
}

func double2(value any) {
	if number, ok := value.(int); ok {
		fmt.Println(number * 2)
	} else {
		fmt.Println("value не является числом")
	}
}

// Проверка нескольких типов
// для этого обычно применяют type switch
func describe(value any) {
	switch v := value.(type) {
	case int:
		fmt.Println("int: ", v)
	case string:
		fmt.Println("string: ", v)
	case bool:
		fmt.Println("bool: ", v)
	default:
		fmt.Println("неизвестный тип: ", v)
	}
}

// Конструкция .(type) используется только внутри type switch. Отдельно написать так нельзя:
// v := value.(type) // ошбика
