package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Sachinó"

	fmt.Printf("str: %v\n", str)
	fmt.Printf("len: %d\n", len(str))
	fmt.Printf("rune: %d\n", []rune(str))
	fmt.Printf("byte: %d\n", []byte(str))

	fmt.Println(strings.Contains(str, "c"))

	str = strings.ToUpper(str)
	fmt.Println(str)

}
