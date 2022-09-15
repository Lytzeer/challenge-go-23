package piscine

func IsUpper(s string) bool {
	for _, ch := range s {
		if rune(ch) < 65 || rune(ch) > 90 {
			return false
		}
	}
	return true
}
