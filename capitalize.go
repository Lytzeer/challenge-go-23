package piscine

func Capitalize(s string) string {
	rep := ""
	for i := 0; i < len(s); i++ {
		if i == 0 && IsLower(string(s[i])) {
			rep += string(rune(s[i] - 32))
		} else if IsLower(string(s[i])) && (!IsAlpha(string(s[i-1]))) {
			rep += string(rune(s[i] - 32))
		} else if i == 0 && IsUpper(string(s[i])) {
			rep += string(rune(s[i]))
		} else if IsUpper(string(s[i])) && IsUpper(string(rep[i-1])) {
			rep += string(rune(s[i] + 32))
		} else if IsUpper(string(s[i])) && IsLower(string(rep[i-1])) {
			rep += string(rune(s[i] + 32))
		} else if IsUpper(string(s[i])) && IsNumeric(string(rep[i-1])) {
			rep += string(rune(s[i] + 32))
		} else {
			rep += string(s[i])
		}
	}
	return rep
}
