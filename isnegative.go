package piscine

func IsNegative(nb int) string {
	nb_str := string(nb)
	if nb_str[0] == '-' {
		return "T"
	} else {
		return "F"
	}
}
