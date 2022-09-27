package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			min := a[i]
			if f(a[i], a[j]) == 1 {
				a[j] = a[i]
				a[i] = min
			}
		}
	}
}
