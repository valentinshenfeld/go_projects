package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	var number uint
	fmt.Println("Ведите основание(Пр: 2, 10, 16): ")
	fmt.Scanln(&number)

	var input string
	for {
		fmt.Println("Enter hex number or 'stop' to exite: ")
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input == "stop" {
			break
		}
		i := new(big.Int)
		if _, ok := i.SetString(processHex(input), int(number)); !ok {
			fmt.Println("Invalid hexadecimal number!")
			continue
		}
		fmt.Println(i)
	}
}

func processHex(hexStr string) string {
	return strings.TrimPrefix(hexStr, "0x")
}
