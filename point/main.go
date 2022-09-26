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

func main() {
	points := &point{}

	setPoint(points)
	x := string(points.x)
	y := string(points.y)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for _, ch := range x {
		z01.PrintRune(ch + 48)
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for _, ch := range y {
		z01.PrintRune(ch + 48)
	}
	z01.PrintRune('\n')
}
