package piscine

func BasicJoin(elems []string) string {
	var word string
	for i := 0; i < len(elems); i++ {
		word += elems[i]
	}
	return word
}
