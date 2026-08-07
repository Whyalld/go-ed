package geometry

import (
	"fmt"
	"math"
)

type Figure interface {
	Area() float64
	Perimeter() float64
}

func CalculateArea(f Figure) float64 {
	return f.Area()
}

func CalculatePerimеter(f Figure) float64 {
	return f.Perimeter()
}

type Circle struct {
	radius float64
}

func NewCircle(radius float64) (*Circle, error) {
	if !validDimension(radius) {
		return nil, fmt.Errorf("некорректный радиус %v: ожидается конечное неотрицательное число", radius)
	}

	return &Circle{radius: radius}, nil
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Rectangle struct {
	length float64
	width  float64
}

func NewRectangle(length float64, width float64) (*Rectangle, error) {
	if !validDimension(length) {
		return nil, fmt.Errorf("некорректная длина %v: ожидается конечное неотрицательное число", length)
	}

	if !validDimension(width) {
		return nil, fmt.Errorf("некорректная ширина %v: ожидается конечное неотрицальеное число", width)
	}

	return &Rectangle{length: length, width: width}, nil
}

func (r *Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func validDimension(value float64) bool {
	return value >= 0 &&
		!math.IsNaN(value) &&
		!math.IsInf(value, 0)
}
