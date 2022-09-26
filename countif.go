package piscine

rep := []bool{}
	for i := 0; i < len(a); i++ {
		rep = append(rep, f(a[i]))
	}
	for i := 0; i < len(rep); i++ {
		if rep[i] {
			return true
		}
	}
	return false