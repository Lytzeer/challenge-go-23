package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb > 20 {
		return 0
	} else {
		rep := 1
		for i := nb; i > 0; i-- {
			rep *= i
		}
		return rep
	}
}
