package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}
func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func printShapeInfo(s Shape) {
	fmt.Printf("Shape: %T\n", s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}
	t := Triangle{A: 3, B: 4, C: 5}

	s := []Shape{c, r, t}
	for _, shape := range s {
		printShapeInfo(shape)

		switch v := shape.(type) {
		case Circle:
			fmt.Printf("Circle with radius %.2f\n", v.Radius)
		case Rectangle:
			fmt.Printf("Rectangle with width %.2f and height %.2f\n", v.Width, v.Height)
		case Triangle:
			fmt.Printf("Triangle with sides %.2f, %.2f, and %.2f\n", v.A, v.B, v.C)
		}
	}

}
