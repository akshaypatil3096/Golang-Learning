package interfacesexample

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}
type circle struct {
	redius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.redius * c.redius
}

func MySquare() square {
	return square{20}
}
func PrintAreas(s square) {
	fmt.Println(s.area())
}
