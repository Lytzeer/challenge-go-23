package piscine

func Join(strs []string, sep string) string {
	var word string
	for i := 0; i < len(strs); i++ {
		if i != len(strs)-1 {
			word += strs[i]
		} else {
			word += strs[i] + sep
		}
	}
	return word
}
