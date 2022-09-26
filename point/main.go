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
	x := string(points.x)
	y := string(points.y)
	tab_x := []rune{}
	tab_y := []rune{}
	for _, ch := range x {
		tab_x = append(tab_x, ch)
	}
	for _, ch := range y {
		tab_y = append(tab_y, rune(ch))
	}
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for i := 0; i < len(tab_x); i++ {
		z01.PrintRune(tab_x[i])
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for i := 0; i < len(tab_y); i++ {
		z01.PrintRune(tab_y[i])
	}
	z01.PrintRune('\n')
}
