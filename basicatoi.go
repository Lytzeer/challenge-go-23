package piscine

func BasicAtoi(s string) int {
	atoi := 0
	n := 1
	slice := []rune{}
	for _, ch := range s {
		slice = append(slice, rune(ch))
	}
	for i := len(slice) - 1; len(slice) != 0; i-- {
		atoi += int(slice[i]-48) * n
		slice = slice[:len(slice)-1]
		n *= 10
	}
	return atoi
}
