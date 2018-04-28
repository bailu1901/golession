package lesson

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (re rect) area() float64 {
	return re.width * re.height
}

func (re rect) perim() float64 {
	return 2*re.width + 2*re.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func InterfaceTest() {
	r := rect{width: 10, height: 20}
	c := circle{radius: 10}

	measure(r)
	measure(c)
}
