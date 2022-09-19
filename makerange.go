package piscine

func MakeRange(min int, max int) []int {
	rep := []int(nil)
	if min < max {
		rep = make([]int, max-min)
		for i := min; i < max; i++ {
			rep[i-min] = i
		}
	}
	return rep
}
