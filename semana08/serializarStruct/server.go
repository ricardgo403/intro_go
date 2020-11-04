package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"./person"
)

// type Person struct {
// 	Nombre string
// 	Email  []string
// }

func server() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {
	var person person.Person
	err := gob.NewDecoder(c).Decode(&person)
	// b := make([]byte, 100)
	// bs, err := c.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Mensaje: ", person)
		// fmt.Println("Bytes: ", bs)
	}
}

func main() {
	go server()
	var input string
	fmt.Scanln(&input)
}
