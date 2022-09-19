package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	if len(arguments) > 1 && arguments[1] == "--upper" {
		for i := 2; i < len(arguments); i++ {
			a := Atoi(arguments[i])
			if a == 0 {
				z01.PrintRune(rune(32))
			} else if a <= 26 {
				z01.PrintRune(rune(a + 64))
			} else {
				z01.PrintRune(rune(32))
			}
		}
		z01.PrintRune(rune('\n'))
	} else if len(arguments) > 1 {
		for i := 1; i < len(arguments); i++ {
			a := Atoi(arguments[i])
			if a == 0 {
				z01.PrintRune(rune(32))
			} else if a <= 26 {
				z01.PrintRune(rune(a + 96))
			} else {
				z01.PrintRune(rune(32))
			}
		}
		z01.PrintRune(rune('\n'))
	}
}

func Atoi(s string) int {
	atoi := 0
	n := 1
	var isnegative bool
	if s == "" {
		return 0
	}
	if s[0] == '-' {
		isnegative = true
	} else {
		isnegative = false
	}
	slice := []rune{}
	for _, ch := range s {
		if ch != '0' && ch != '1' && ch != '2' && ch != '3' && ch != '4' && ch != '5' && ch != '6' && ch != '7' && ch != '8' && ch != '9' && ch != '+' && ch != '-' {
			return 0
		} else {
			slice = append(slice, rune(ch))
		}
	}
	if slice[0] == '-' || slice[0] == '+' {
		slice = slice[1:]
	}
	for i := len(slice) - 1; len(slice) != 0; i-- {
		if slice[i] == '-' || slice[i] == '+' {
			return 0
		}
		atoi += int(slice[i]-48) * n
		slice = slice[:len(slice)-1]
		n *= 10
	}
	if isnegative == true {
		atoi *= -1
	}
	return atoi
}
