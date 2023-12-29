package main

import (
	"fmt"
	"math"
)

func main() {
	cirlce := Circle{
		Redius: 5,
	}

	PrintAreas(cirlce)
}

type Square struct {
	Length float64
}
type Circle struct {
	Redius float64
}

type areaCalculate interface {
	area() float64
}

func (s Square) area() float64 {
	return s.Length * s.Length
}

func (c Circle) area() float64 {
	return math.Pi * c.Redius * c.Redius
}

func PrintAreas(a areaCalculate) {
	fmt.Println(a.area())
}
