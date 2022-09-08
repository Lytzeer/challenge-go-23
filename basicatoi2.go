package piscine

func BasicAtoi2(s string) int {
	atoi := 0
	n := 1
	slice := []rune{}
	for _, ch := range s {
		if ch != '0' && ch != '1' && ch != '2' && ch != '3' && ch != '4' && ch != '5' && ch != '6' && ch != '7' && ch != '8' && ch != '9' {
			return 0
		} else {
			slice = append(slice, rune(ch))
		}
	}
	for i := len(slice) - 1; len(slice) != 0; i-- {
		atoi += int(slice[i]-48) * n
		slice = slice[:len(slice)-1]
		n *= 10
	}
	return atoi
}
