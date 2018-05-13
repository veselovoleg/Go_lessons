package main

import (
	"fmt"
	"math"
	)

func main(){
	
	c := Circle{0, 0, 5}//initialization
	

	fmt.Println(c.x, c.y, c.r)//Get acces with . operator
	//0 0 5
	c.x = 10
	c.y = 5
	c.r = 6
	fmt.Println(c.x, c.y, c.r) // 10 5 6

	fmt.Println(circleArea(c)) // (6^2)*Pi = 113.09

	fmt.Println(c.area())//We can call func with . (c.area())

	r := Rectangle{0, 0, 10, 10} 
	fmt.Println(r.area()) //100
	
}

type Circle struct { //Creating of new struct
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r*c.r
}

func (c *Circle) area() float64 { // area() is getter.
    return math.Pi * c.r*c.r
}
//rectangle example

func distance(m1, t1, m2, t2 float64) float64 {
    a := m2 - m1
    b := t2 - t1
    return math.Sqrt(a*a + b*b)
}

type Rectangle struct {
    z1, p1, z2, p2 float64
}
func (r *Rectangle) area() float64 {
    l := distance(r.z1, r.p1, r.z1, r.p2)
    w := distance(r.z1, r.p1, r.z2, r.p1)
    return l * w
} 