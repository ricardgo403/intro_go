package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Nombre string
	Edad   uint64
}

type ByNombre []Person

func (a ByNombre) Len() int           { return len(a) }
func (a ByNombre) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByNombre) Less(i, j int) bool { return a[i].Nombre < a[j].Nombre }

type ByEdad []Person

func (a ByEdad) Len() int           { return len(a) }
func (a ByEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

func main() {
	ps := []Person{
		Person{
			Nombre: "Joel",
			Edad:   21,
		},
		Person{
			Nombre: "Ricardo Joel",
			Edad:   17,
		},
		Person{
			Nombre: "Ricardo Gonzalez",
			Edad:   23,
		},
	}
	fmt.Println(ps)
	sort.Sort(ByNombre(ps))
	fmt.Println(ps)
	sort.Sort(ByEdad(ps))
	fmt.Println(ps)
}
