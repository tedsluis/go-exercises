package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (s *Rectangle) area() float64 {
	x := math.Abs(s.x2 - s.x1)
	y := math.Abs(s.y2 - s.y1)
	return x * y
}

func (s *Rectangle) distance() float64 {
	x := s.x2 - s.x1
	y := s.y2 - s.y1
	return math.Sqrt(x*x + y*y)
}

type Circle struct {
	x, y, r float64
}

func (s *Circle) area() float64 {
	fmt.Println("area:", s.r, ", x:", s.x, " ,y:", s.y)
	return math.Pi * s.r * s.r
}

func main() {
	r1 := Rectangle{1, 1, 4, 4}
	fmt.Println("Rectangle1: ", r1.area())
	r2 := Rectangle{2, 4, 3, 6}
	fmt.Println("Rectangle2: ", r2.area())
	c1 := Circle{8, 12, 3}
	fmt.Println("Circle1: ", c1.area())
	c2 := Circle{10, 10, 2}
	fmt.Println("Circle2: ", c2.area())
	m := MultiShape{[]Shape{&r1, &r2, &c1, &c2}}
	fmt.Println(m.area())
}
