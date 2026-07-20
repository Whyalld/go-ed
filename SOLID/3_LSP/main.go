package main

import "fmt"

/* Код нарушающий LSP
type Bird interface {
	Fly() string
}

type Duck struct{}
func (d *Duck) Fly() string { return "Утка летит в небе!"}

type Ostrich struct{}
func (o *Ostrich) Fly() string {
	panic("Страус не умеет летать")
}

func MakeBirdFly(b Bird) {
	fmt.Println(b.Fly())
}

func main() {
	var bird Bird
	bird = &Duck{}
	MakeBirdFly(bird)
} */

// Правильный код по LSP
type Walker interface {
	Walk() string
}

type Flyer interface {
	Fly() string
}

type Duck struct{}
func (d Duck) Walk() string { return "Утка идет" }
func (d Duck) Fly() string { return "Утка летит" }

type Ostrich struct{}
func (o Ostrich) Walk() string { return "Страус бежит со скоростью 50 км/ч" }

func MakeItFly(f Flyer) {
	fmt.Println(f.Fly())
}

func MakeItWalk(w Walker) {
	fmt.Println(w.Walk())
}

func main() {
	duck := Duck{}
	ostrich := Ostrich{}

	MakeItFly(duck)
	// MakeItFly(ostrich) // Компилятор Go не скомпилируем этот код, тем самым защитит нас от ошибки на этапе сборки, потому что у Ostrich нет метода Fly!
	MakeItWalk(duck)
	MakeItWalk(ostrich)
}