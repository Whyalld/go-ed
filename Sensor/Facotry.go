package main

import (
	"fmt"
	"math/rand"
	// "strconv"
)

type ISensorVisitor interface {
	VisitHumidity(*HumiditySensor)
	VisitTemp(*TempSensor)
	VisitPressure(*PressureSensor)
}

type ISensor interface {
	GetName() string
	GetValue() int
	Accept(ISensorVisitor)
}

type Sensor struct {
	name  string
	price int
}
func (s *Sensor) GetName() string { return s.name }
// func (s *Sensor) GetPrice() string {return "$ "+strconv.Itoa(s.price)}

type HumiditySensor struct { Sensor }
func (h *HumiditySensor) GetValue() int { return rand.Intn(100) }
func (h *HumiditySensor) Accept(v ISensorVisitor) {	v.VisitHumidity(h) }

type TempSensor struct { Sensor }
func (t *TempSensor) GetValue() int { return rand.Intn(40) }
func (h *TempSensor) Accept(v ISensorVisitor) {	v.VisitTemp(h) }

type PressureSensor struct { Sensor }
func (p *PressureSensor) GetValue() int { return rand.Intn(50) + 950 }
func (h *PressureSensor) Accept(v ISensorVisitor) {	v.VisitPressure(h) }

func CreateSensor(sensorType string) (ISensor, error) {
	switch sensorType {
	case "humidity":
		return &HumiditySensor{Sensor{name: "HSsrx788", price: rand.Intn(1000)}}, nil
	case "temperature":
		return &TempSensor{Sensor{name: "TSlkj112", price: rand.Intn(1000)}}, nil
	case "pressure":
		return &PressureSensor{Sensor{name: "PSbn2", price: rand.Intn(1000)}}, nil
	default:
		return nil, fmt.Errorf("unknown sensor type: %s", sensorType)
	}
}


