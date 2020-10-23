package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	radio float64
}

func (c *Circulo) area() float64 {
	return c.radio * c.radio * math.Pi
}

func main() {
	c01 := Circulo{5}
	// c02 := Circulo{radio: 10}
	// c03 := Circulo{15}
	// c04 := new(Circulo)
	// c05 := &Circulo{200}
	fmt.Println(c01.area())
	// fmt.Println(c02)
	// fmt.Println(c03)
	// fmt.Println(c04)
	// fmt.Println(c05)
}
