package piscine

func StrRev(s string) string {
	var word string
	slice := []string{}
	for _, ch := range s {
		slice = append(slice, string(ch))
	}
	for i := len(slice) - 1; len(slice) != 0; i-- {
		word += slice[i]
		slice = slice[:len(slice)-1]
	}
	return word
}
