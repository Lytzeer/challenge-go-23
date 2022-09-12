package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb > 10 {
		return 0
	} else {
		rep := 1
		for i := nb - 1; i > 0; i-- {
			rep *= i
		}
		return rep
	}
}
