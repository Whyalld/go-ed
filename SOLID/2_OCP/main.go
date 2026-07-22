package main

import "fmt"

type DeliveryService interface {
	Deliver(order Order)
}

type Order struct {
	TotalCost int
}

type CourierDelivery struct{}

func (c *CourierDelivery) Deliver(order Order) {
	fmt.Println("Доставка курьером: Заказ доставлен.")
}

type PostDelivery struct{}

func (p *PostDelivery) Deliver(order Order) {
	fmt.Println("Почта России: Посылка отправлена.")
}

type CdekDelivery struct{}

func (c *CdekDelivery) Deliver(order Order) {
	fmt.Println("СДЭК: Заказ доставлен в пункт выдачи")
}

func ProcessOrder(order Order, ds DeliveryService) {
	ds.Deliver(order)
}

func main() {
	order := Order{TotalCost: 1000}

	var myDelivery DeliveryService

	myDelivery = &CourierDelivery{}
	ProcessOrder(order, myDelivery)

	myDelivery = &PostDelivery{}
	ProcessOrder(order, myDelivery)
}
