package main

import (
	"fmt"
	"math"
)

type Calculate interface {
	area() float64
}
type Rect struct {
	Length float64
	Width  float64
}
type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (r Rect) area() float64 {
	return r.Length * r.Width
}
func cacularea(a Calculate) {
	fmt.Println("Shape: ", a)
	fmt.Println("Area: ", a.area())

}
func main() {
	rectangle := Rect{Length: 2, Width: 3}
	cacularea(rectangle)
	circle := Circle{Radius: 10}
	cacularea(circle)
}
