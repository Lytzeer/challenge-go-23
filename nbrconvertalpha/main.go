package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	if len(arguments) > 1 && arguments[1] == "--upper" {
		tab := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
		for i := 2; i < len(arguments); i++ {
			a := Atoi(arguments[i])
			if a == 0 {
				z01.PrintRune(' ')
			} else if a <= 26 {
				z01.PrintRune(tab[a-1])
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune(rune('\n'))
	} else if len(arguments) > 1 {
		tab := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
		for i := 1; i < len(arguments); i++ {
			a := Atoi(arguments[i])
			if a == 0 {
				z01.PrintRune(' ')
			} else if a <= 26 {
				z01.PrintRune(tab[a-1])
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune(rune('\n'))
	}
}

func Atoi(s string) int {
	atoi := 0
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			atoi = atoi*10 + int(ch-'0')
		} else {
			return 0
		}
	}
	return atoi
}
