package piscine

func IsPrintable(s string) bool {
	for _, ch := range s {
		if rune(ch) < 33 || rune(ch) > 126 {
			return false
		}
	}
	return true
}
