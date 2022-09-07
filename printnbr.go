package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('3')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('5')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('7')
		z01.PrintRune('5')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('8')
		return
	}
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
