package main

import (
	"errors"
	"fmt"
)

type MyCustomError struct {
	Code int
}

func (e *MyCustomError) Error() string {
	return "Произошла ошибка"
}

func main() {
	err1 := errors.New("hello")
	err2 := fmt.Errorf("%w world", err1)
	fmt.Println(err2)
	fmt.Println(errors.Unwrap(err2))

	// erros.Is
	var e bool = errors.Is(err2, err1)
	fmt.Println(e)

	// erros.AsType
	if er, k := errors.AsType[*MyCustomError](err2); k {
		fmt.Println("Код ошибки:", er.Code)
	}
}
