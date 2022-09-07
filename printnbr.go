package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = n * -1
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	numberSlice := []rune{}
	for n > 0 {
		nb := n % 10
		numberSlice = append(numberSlice, rune(nb+48))
		n /= 10
	}
	for i := len(numberSlice) - 1; len(numberSlice) != 0; i-- {
		z01.PrintRune(numberSlice[i])
		numberSlice = numberSlice[:len(numberSlice)-1]
	}
}
