package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	// map
	id := [5]byte{3, 4, 1, 5, 3}
	fmt.Println(id)

	pops := id[0:4]
	pops = append(pops, 1)
	fmt.Println(pops)
	fmt.Println(id)

	rikby := id[3:5]
	rikby[0] = 9
	fmt.Println(id)


	pops = append(pops, 8)
	fmt.Println(id)
	fmt.Println(pops)

	fmt.Println("\n==========================\n")

	var list []int
	fmt.Println(list == nil)
	fmt.Println(len(list) == 0, "\n")

	list = []int{}
	fmt.Println(list == nil)
	fmt.Println(len(list) == 0)

	fmt.Println("\n===========================\n")
	
	// Аллокации памяти
	mordekai := make([]int, 0, 10)
	mordekai = append(mordekai, 1, 4, 5, 6, 1)
	fmt.Println(len(mordekai), cap(mordekai), mordekai)
	println()
	coolerMordekai := Double(mordekai)
	fmt.Println(coolerMordekai)
	mordekai = DoubleButDifferent(mordekai)
	fmt.Println(mordekai)

	fmt.Println("\n============================\n")
	
	// Передача слайсов по значению и по ссылке
	trash := []int{1, 5, 2, 90}
	fmt.Println("before:", trash)
	Do(trash[:1])
	Handle(trash)
	fmt.Println("after:", trash)

	fmt.Println("\n============================\n")

	banson := make([]int, 4, 5)
	banson = append(banson, 1)
	banson[0] = 5
	banson[2] = 38
	fmt.Println(banson)

	fmt.Println("\n============================\n")

	ls := make([]int, 4, 4)
	ls = Append(ls, 2)
	fmt.Println(ls, len(ls), cap(ls))
}

func Double(nums []int) []int {
	res := make([]int, 0, len(nums))
	for _, value := range nums {
		res = append(res, value*2)
	}
	return res
}


func DoubleButDifferent(nums []int) []int {
	res := make([]int, len(nums))
	for i, value := range nums {
		res[i] = value * 2
	}
	return res
}


func Handle(list []int) {
	list[2] = 484928
}

func Do(list []int) {
	newList := make([]int, len(list))
	copy(newList, list)
	newList = append(newList, 42)
	fmt.Println("append:", newList)
}


var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) ([]byte, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	b = digitRegexp.Find(b)
	res := make([]byte, len(b))
	copy(res, b)
	return res, nil
}

func findDigitsWithAppend(fileName string) ([]byte, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	b = digitRegexp.Find(b)
	res := append([]byte(nil), b...)
	return res, nil
}

func Append(list []int, el int) []int {
	var res []int
	resLen := len(list) + 1
	if resLen <= cap(list) {
		res = list[:resLen]
	} else {
		resCap := resLen
		if resCap < 2 * len(list) {
			resCap = 2 * len(list)
		}
		res = make([]int, resLen, resCap)
		copy(res, list)
	}

	res[len(list)] = el

	return res
}
