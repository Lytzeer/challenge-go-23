package piscine

func StrLen(s string) int {
	cpt := 0
	slice := []rune{}
	for _, ch := range s {
		slice = append(slice, rune(ch))
		cpt = len(slice)
	}
	return cpt
}
