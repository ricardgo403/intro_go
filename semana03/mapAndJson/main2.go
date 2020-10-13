package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]int{
		"Programación": 100,
		"Algoritmia":   90,
	}
	fmt.Println(m["Programación"])
	m["Programación"] = 95
	if v, err := m["eda1"]; !err {
		fmt.Println("Tiene algo:", err)
	} else {
		fmt.Println("Valor: ", v)
	}
	j, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)
	fmt.Println(j)
	fmt.Println(string(j))
}
