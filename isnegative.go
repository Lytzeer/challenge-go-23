package piscine

import (
	"github.com/01-edu/z01"
)

func IsNegative(nb int) {
	nb_int := int(string(string(nb))[0])
	if nb_int >= 0 && nb_int <= 9 {
		z01.PrintRune('F')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('T')
		z01.PrintRune('\n')
	}
}
