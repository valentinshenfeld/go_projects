package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	phonebook := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("phonebook")
	fmt.Println("")

	fmt.Println("Hello, World!", phonebook)
}
