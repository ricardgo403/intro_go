package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var line string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your full name: ")
	scanner.Scan()
	line = scanner.Text()

	fmt.Println("Your name is: ", line)
}
