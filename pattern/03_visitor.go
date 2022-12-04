package pattern

import (
	"fmt"
	"math"
)

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Shape interface {
	GetPosition() (float64, float64)
	PrintInfo()
	Accept(Visitor)
}

type Rectangle struct {
	X, Y                  float64
	Width, Height         float64
	halfWidth, halfHeight float64
}

func (r Rectangle) GetPosition() (float64, float64) {
	return r.X, r.Y
}

func (r Rectangle) PrintInfo() {
	fmt.Printf("rectangle (%f, %f) %f x %f\n", r.X, r.Y, r.Width, r.Height)
}

func (r Rectangle) Accept(v Visitor) {
	v.visitRectangle(r)
}

type Circle struct {
	X, Y, Radius float64
}

func NewCircle(x, y, radius float64) *Circle {
	return &Circle{
		X: x, Y: y, Radius: radius}
}

func (c Circle) GetPosition() (float64, float64) {
	return c.X, c.Y
}

func (c Circle) PrintInfo() {
	fmt.Printf("circle (%f, %f) r=%f\n", c.X, c.Y, c.Radius)
}

func (c Circle) Accept(v Visitor) {
	v.visitCircle(c)
}

type Visitor interface {
	visitRectangle(Rectangle)
	visitCircle(Circle)
}

type AreaCalculator struct {
	area float64
}

func (a AreaCalculator) visitRectangle(r Rectangle) {
	fmt.Println("area of rectangle:", r.Width*r.Height)
}

func (a AreaCalculator) visitCircle(c Circle) {
	fmt.Println("area of circle:", math.Pi*c.Radius*c.Radius)
}

func VisitorExample() {
	shapes := make([]Shape, 0, 5)
	shapes = append(shapes, Circle{X: 5, Y: 4, Radius: 3})
	shapes = append(shapes, Circle{X: 7, Y: 40, Radius: 28})
	shapes = append(shapes, Rectangle{X: 30, Y: -4, Width: 5, Height: 10})
	shapes = append(shapes, Rectangle{X: 15, Y: 90, Width: 120, Height: 75})
	shapes = append(shapes, Circle{X: -18, Y: 64, Radius: 7.5})

	ac := AreaCalculator{}

	for _, shape := range shapes {
		shape.PrintInfo()
		shape.Accept(ac)
	}
}
