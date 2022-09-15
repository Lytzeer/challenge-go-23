package piscine

func IsAlpha(s string) bool {
	for _, ch := range s {
		if rune(ch) < 65 || rune(ch) > 90 && rune(ch) < 97 || rune(ch) > 122 && rune(ch) < 48 || rune(ch) > 57 {
			return false
		}
	}
	return true
}
