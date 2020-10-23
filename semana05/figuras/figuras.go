package figuras

import "math"

type FiguraMulti struct {
	Figuras []Figura
}

func (fm *FiguraMulti) Area() float64 {
	sum := 0.0
	for _, f := range fm.Figuras {
		sum += f.Area()
	}
	return sum
}

type Figura interface {
	Area() float64
}

type Circulo struct {
	Radio float64
}

func (c *Circulo) Area() float64 {
	return c.Radio * c.Radio * math.Pi
}

type Rectangulo struct {
	Base   float64
	Altura float64
}

func (r *Rectangulo) Area() float64 {
	return r.Altura * r.Base
}
