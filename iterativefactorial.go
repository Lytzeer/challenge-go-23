package piscine

func IterativeFactorial(nb int) int {
	if nb <= 0 {
		return 0
	} else {
		rep := nb
		for i := nb - 1; i > 0; i-- {
			rep *= i
		}
		return rep
	}
}
