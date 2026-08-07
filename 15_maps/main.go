package main

import "fmt"

func main() {
	var ages map[string]int
	fmt.Println(ages, ages == nil)
	// ages["Maxim"] = 20 -> ошибка будет

	// Сначала нужно make сделать, потом вносить значения
	ages = make(map[string]int)
	fmt.Println(ages, ages == nil)

	ages["Maxim"] = 20
	ages["Deni"] = 31
	ages["Isa"] = 26

	for key, value := range ages {
		fmt.Printf("%s - %d\n", key, value)
	}

	fmt.Println("Возраст Максима:", ages["Maxim"])

	// Можно инициализировать мапу сразу со значениями
	var data = map[string]int{
		"Саид": 100,
		"Дени": 20,
	}
	fmt.Printf("У Саида %d баллов, а у Дуни %d\n", data["Саид"], data["Пися"])

	// Проверка на сущ-е в мапе
	_, exists := ages["Deni"]
	if !exists {
		fmt.Println("Не сущ")
	} else {
		fmt.Println("Сущ")
	}

	// delete
	delete(data, "Дени")
	fmt.Println(data)

	// Связанность
	agesCopy := make(map[string]int)
	agesCopy = ages
	delete(agesCopy, "Maxim")
	delete(agesCopy, "Deni")
	fmt.Println(ages)
	fmt.Println(agesCopy)
}
