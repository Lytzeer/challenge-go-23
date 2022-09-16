package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	p := []rune(arguments[0])
	for i := 2; i < len(p); i++ {
		z01.PrintRune(rune(p[i]))
	}
	z01.PrintRune('\n')
}
