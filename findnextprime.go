package piscine

func FindNextPrime(nb int) int {
	for i := nb; i < 20000000; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}
