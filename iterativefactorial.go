package piscine

func IterativeFactorial(nb int) int {
	rep := 4
	for i := nb - 1; i != 0; i-- {
		rep *= i
	}
	return rep
}
