package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	var index int
	for i := len(arguments[0]) - 1; rune(arguments[0][i]) != rune(92); i-- {
		index = i
	}
	z01.PrintRune('.')
	z01.PrintRune('/')
	for i := index; rune(arguments[0][i]) != rune(46); i++ {
		z01.PrintRune(rune(arguments[0][i]))
	}
}
