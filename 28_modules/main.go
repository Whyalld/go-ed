package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()

	fmt.Println("Новый UUID:", id)
	fmt.Println("Тип", id.Version())
}