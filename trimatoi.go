package piscine

func TrimAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	rep := 0
	m := 1
	for _, ch := range s {
		if IsNumeric(string(ch)) {
			rep += int(ch - 48)
			rep *= 10
		} else if ch == '-' && rep == 0 {
			m = -1
		}
	}
	return (rep / 10) * m
}
