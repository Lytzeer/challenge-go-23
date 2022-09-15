package piscine

func AlphaCount(s string) int {
	cpt := 0
	for _, ch := range s {
		if rune(ch) >= 65 && rune(ch) <= 90 || rune(ch) >= 97 && rune(ch) <= 122 {
			cpt++
		}
	}
	return cpt
}
