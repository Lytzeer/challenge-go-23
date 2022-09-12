package piscine

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if power < 0 || nb < 0 {
		return 1
	} else {
		return RecursivePower(nb, power-1) * nb
	}
}
