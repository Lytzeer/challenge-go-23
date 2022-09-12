package piscine

func IterativePower(nb int, power int) int {
	if nb < 0 || power < 0 {
		return 0
	} else if nb == 0 && power == 0 {
		return 1
	} else {
		rep := nb
		for i := power; i != 1; i-- {
			rep *= nb
		}
		return rep
	}
}
