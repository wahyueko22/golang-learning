package main

import (
	"fmt"
	"math"
)

type geometry interface {
    area() float64
}

type rect struct {
    width, height float64
}

type circle struct {
    radius float64
}

type triangle struct{
	base  float64
	height float64
} 

func (r rect) area() float64 {
    return r.width * r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (t triangle) area() float64{
	return t.base * t.height * 0.5
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
	measure(c)
	
	var geo geometry
	geo = triangle{base:10, height: 5}

	fmt.Println(geo)
    fmt.Println(geo.area())

}