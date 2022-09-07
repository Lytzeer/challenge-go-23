package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '8'; j++ {
			for k := '0'; k <= '9'; k++ {
				for l := '0'; l <= '9'; l++ {
					nb1 := i*10 + j
					nb2 := k*10 + l
					if nb1 != nb2 && nb1 < nb2 {
						if i == '9' && j == '8' && k == '9' && l == '9' {
							z01.PrintRune(rune(i))
							z01.PrintRune(rune(j))
							z01.PrintRune(' ')
							z01.PrintRune(rune(k))
							z01.PrintRune(rune(l))
						} else {
							z01.PrintRune(rune(i))
							z01.PrintRune(rune(j))
							z01.PrintRune(' ')
							z01.PrintRune(rune(k))
							z01.PrintRune(rune(l))
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
