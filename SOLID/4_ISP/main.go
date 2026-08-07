/* Четвертый принцип SOLID: ISP (Принцип разделения интерфейса)
Формулировка принципа: «Клиенты не должны зависеть от методов, которые они не используют».
Если перевести это на простой язык: лучше создать много маленьких, узкоспециализированных
интерфейсов, чем один огромный «всемогущий» интерфейс.В Go этот принцип возведен в абсолют.
Если вы откроете исходный код стандартной библиотеки Go, вы увидите, что большинство интерфейсов
там содержат всего по одному методу (например, io.Reader, io.Writer, fmt.Stringer). */

package main

import "fmt"

type Flyer interface {
	Fly() string
}

type Runner interface {
	Run() string
}

type Duck struct{}

func (d *Duck) Fly() string {
	return "Утка летит в небе!"
}

type Ostrich struct{}

func (o *Ostrich) Run() string {
	return "Страус бежит со скоростью 50 км/ч!"
}

func MakeItFly(f Flyer) {
	fmt.Println(f.Fly())
}

func MakeItRun(r Runner) {
	fmt.Println(r.Run())
}

func main() {
	duck := &Duck{}
	ostrich := &Ostrich{}

	MakeItFly(duck)
	MakeItRun(ostrich)
}
