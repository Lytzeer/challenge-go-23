package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]
	for i := 2; i < len(arguments); i++ {
		z01.PrintRune(rune(arguments[i]))
	}
}
