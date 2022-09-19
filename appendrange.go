package piscine

func AppendRange(min, max int) []int {
	rep := []int(nil)
	if min < max {
		for i := min; i < max; i++ {
			rep = append(rep, i)
		}
	}
	return rep
}
