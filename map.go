package piscine

func Map(f func(int) bool, a []int) []bool {
	rep := []bool{}
	for i := 0; i < len(a); i++ {
		rep = append(rep, f(a[i]))
	}
	return rep
}
