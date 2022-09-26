package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune(rune('x'))
	z01.PrintRune(rune(32))
	z01.PrintRune(rune('='))
	z01.PrintRune(rune(32))
	listee := []rune{}
	for points.x > 0 {
		nb := points.x % 10
		listee = append(listee, rune(nb+48))
		points.x /= 10

	}
	for i := len(listee) - 1; i >= 0; i-- {
		z01.PrintRune(listee[i])
	}
	z01.PrintRune(rune(points.x))

	z01.PrintRune(rune(','))
	z01.PrintRune(rune(32))
	z01.PrintRune(rune('y'))
	z01.PrintRune(rune(32))
	z01.PrintRune(rune('='))
	z01.PrintRune(rune(32))
	liste := []rune{}
	for points.y > 0 {
		nb := points.y % 10
		liste = append(liste, rune(nb+48))
		points.y /= 10

	}
	for i := len(liste) - 1; i >= 0; i-- {
		z01.PrintRune(liste[i])
	}
	z01.PrintRune(rune(points.y))
	z01.PrintRune('\n')
}
