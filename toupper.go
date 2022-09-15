package piscine

func ToUpper(s string) string {
	var word_upper string
	for _, ch := range s {
		if IsLower(string(ch)) {
			word_upper += string(rune(ch - 32))
		} else {
			word_upper += string(ch)
		}
	}
	return word_upper
}
