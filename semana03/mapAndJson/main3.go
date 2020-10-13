package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]map[string]int{
		"Primero": {
			"Programación": 100,
		},
		"Tercero": {
			"Algoritmia": 90,
		},
	}
	fmt.Println(m["Primero"])
	m["Primero"]["Programación"] = 95
	// if v, err := m["eda1"]; !err {
	// 	fmt.Println("Tiene algo:", err)
	// } else {
	// 	fmt.Println("Valor: ", v)
	// }
	j, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)
	fmt.Println(j)
	fmt.Println(string(j))
}
