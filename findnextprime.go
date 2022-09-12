package piscine

func FindNextPrime(nb int) int {
	rep := 0
	if IsPrime(nb) == true {
		return nb
	} else {
		for i := nb + 1; IsPrime(i) == true; i++ {
			rep = i
		}
		return rep
	}
}
