package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]
	z01.PrintRune('.')
	z01.PrintRune('/')
	for _, ch := range arguments {
		z01.PrintRune(ch)
	}
}
