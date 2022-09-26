package piscine

func CountIf(f func(string) bool, tab []string) int {
	rep := []bool{}
	for i := 0; i < len(tab); i++ {
		rep = append(rep, f(tab[i]))
	}
	cpt := 0
	for i := 0; i < len(rep); i++ {
		if rep[i] {
			cpt += 1
		}
	}
	return cpt
}
