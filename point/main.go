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
	x := string(string(points.x))
	y := string(string(points.y))
	word := "x = " + x + ", y = " + y
	printStr(word)
}
