package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["Programación"] = 100
	m["Algoritmia"] = 90

	j, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)
	fmt.Println(j)
	fmt.Println(string(j))
}
