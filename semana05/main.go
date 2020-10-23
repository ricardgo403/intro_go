package main

import (
	"fmt"

	"./figuras"
)

func sumArea(figs ...figuras.Figura) float64 {
	sum := 0.0
	for _, f := range figs {
		sum += f.Area()
	}
	return sum
}

func main() {

	c01 := figuras.Circulo{Radio: 10}
	r01 := figuras.Rectangulo{Altura: 10, Base: 20}
	// c02 := Circulo{radio: 10}
	// c03 := Circulo{15}
	// c04 := new(Circulo)
	// c05 := &Circulo{200}
	fmt.Println(c01.Area())
	fmt.Println(r01.Area())

	fmt.Println(sumArea(&c01, &r01))

	figMulti := figuras.FiguraMulti{
		Figuras: []figuras.Figura{
			&c01,
			&r01,
			// &figuras.Circulo{2},
		},
	}

	fmt.Println(figMulti.Area())
	// fmt.Println(c02)
	// fmt.Println(c03)
	// fmt.Println(c04)
	// fmt.Println(c05)
}
