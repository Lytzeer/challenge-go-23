package piscine

func IsAlpha(s string) bool {
	for _, ch := range s {
		if !(IsUpper(string(ch))) && !(IsLower(string(ch))) && (rune(ch) < 48 || rune(ch) > 57) {
			return false
		}
	}
	return true
}
