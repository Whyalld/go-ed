package main

import "fmt"

type BaseTransformer struct {
	Name string
}

func (t *BaseTransformer) run() {
	fmt.Printf("%s бежит!\n", t.Name)
}

func (t *BaseTransformer) fire() {
	fmt.Printf("%s стреляет\n", t.Name)
}

type Autobot struct {
	BaseTransformer
}

func (a *Autobot) transform() {
	fmt.Printf("%s трансформировался в автомобиль\n", a.Name)
}

type Deception struct {
	BaseTransformer
}

func (d *Deception) transform() {
	fmt.Printf("%s рансформировался в самолет\n", d.Name)
}

func main() {
	robot1 := &Autobot{
		BaseTransformer: BaseTransformer{Name: "Бамблби"},
	}

	robot2 := new(Deception)
	robot2.Name = "Старскрим"

	robot1.run()
	robot1.fire()
	robot1.transform()

	fmt.Println()

	robot2.run()
	robot2.fire()
	robot2.transform()
}
