package piscine

func RecursiveFactorial(nb int) int {
	if nb < 20 {
		if nb == 1 {
			return nb
		} else if nb == 0 {
			return 1
		} else {
			return RecursiveFactorial((nb - 1)) * nb
		}
	} else {
		return 0
	}
}
