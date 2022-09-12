package piscine

func RecursiveFactorial(nb int) int {
	if nb == 1 {
		return nb
	} else {
		return RecursiveFactorial((nb - 1)) * nb
	}
}
