package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rectangle struct {
	wight  float64
	height float64
}

type geometry interface {
	area() float64
	perim() float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perim() float64 {
	return c.radius * 2 * math.Pi
}

func (r rectangle) area() float64 {
	return r.height * r.wight
}

func (r rectangle) perim() float64 {
	return (r.height + r.wight) * 2
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rectangle{wight: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
