package piscine

func Sqrt(nb int) int {
	rep := 0
	if nb <= 0 {
		return rep
	} else if nb == 1 {
		return 1
	} else {
		for i := 1; i < nb; i++ {
			if i*i == nb {
				rep = i
			}
		}
		return rep
	}
}
