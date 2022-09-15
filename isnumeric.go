package piscine

func IsNumeric(s string) bool {
	for _, ch := range s {
		if rune(ch) < 48 || rune(ch) > 57 {
			return false
		}
	}
	return true
}
