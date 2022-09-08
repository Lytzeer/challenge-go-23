package piscine

func Atoi(s string) int {
	atoi := 0
	n := 1
	var isnegative bool
	if s == "" {
		return 0
	}
	if s[0] == '-' {
		isnegative = true
	} else {
		isnegative = false
	}
	slice := []rune{}
	for _, ch := range s {
		if ch != '0' && ch != '1' && ch != '2' && ch != '3' && ch != '4' && ch != '5' && ch != '6' && ch != '7' && ch != '8' && ch != '9' && ch != '+' && ch != '-' {
			return 0
		} else {
			slice = append(slice, rune(ch))
		}
	}
	if slice[0] == '-' || slice[0] == '+' {
		slice = slice[1:]
	}
	for i := len(slice) - 1; len(slice) != 0; i-- {
		if slice[i] == '-' || slice[i] == '+' {
			return 0
		}
		atoi += int(slice[i]-48) * n
		slice = slice[:len(slice)-1]
		n *= 10
	}
	if isnegative == true {
		atoi *= -1
	}
	return atoi
}
