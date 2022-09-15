package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	n_tab := []rune{}
	if n <= 0 {
		z01.PrintRune('0')
	} else {
		for n > 0 {
			div := n % 10
			n /= 10
			n_tab = append(n_tab, rune(div))
		}
		for i := 0; i < len(n_tab); i++ {
			for j := 0; j < len(n_tab); j++ {
				tmp := n_tab[i]
				if n_tab[i] < n_tab[j] {
					n_tab[i] = n_tab[j]
					n_tab[j] = tmp
				}
			}
		}
		for _, ch := range n_tab {
			z01.PrintRune(rune(ch + 48))
		}
	}
}
