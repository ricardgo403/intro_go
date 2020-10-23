package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello world", "Helo"))
	fmt.Println(strings.Count("Hello world", "l"))
	fmt.Println(strings.HasPrefix("Hello world", "Hell"))
	fmt.Println(strings.HasSuffix("Hello world", "world"))
	fmt.Println(strings.Index("Hello world", "world"))
	fmt.Println(strings.Index("Hello world", "qas"))
	fmt.Println(strings.Join([]string{"Hello", "world", "!"}, " "))
	fmt.Println(strings.Join([]string{"Hello", "world", "!"}, ""))
	fmt.Println(strings.Repeat("Hello", 2))
	fmt.Println(strings.Replace("Hello", "l", "b", 2))
	fmt.Println(strings.Split("aaaa bbbb cccc dddd ee f", " "))
	fmt.Println(strings.ToLower("Hello World!"))
	fmt.Println(strings.ToUpper("Hello World!"))

	b := []byte("Hello world!")
	fmt.Println(b)

	s := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(s)
}
