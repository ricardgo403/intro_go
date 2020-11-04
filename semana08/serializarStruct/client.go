package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"./person"
)

func client(person person.Person) {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	// msg := "Hello world! :)"
	// fmt.Println(msg)

	err2 := gob.NewEncoder(c).Encode(person)
	if err != nil {
		fmt.Println(err2)
		return
	}
	// c.Write([]byte(msg))
	// c.Write()
}

func main() {
	person := person.Person{
		Nombre: "Ricardo",
		Email: []string{
			"ricardojoel.gonzalezv@gmail.com",
			"ricardo.gonzalez5243@alumnos.udg.mx",
		},
	}
	go client(person)
	var input string
	fmt.Scanln(&input)
}
