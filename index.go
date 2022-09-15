package piscine

func Index(s string, toFind string) int {
	rep := -1
	if len(toFind) < len(s) {
		return -1
	} else if len(toFind) == 0 {
		return 0
	} else {
		for i := 0; i < len(s); i++ {
			for j := 0; j < len(toFind); j++ {
				if s[i] == toFind[j] {
					if len(toFind) == 1 {
						return i
					} else {
						var mot string
						for k := 0; k < len(toFind); k++ {
							mot += string(string(s[i+k]))
						}
						if mot == toFind {
							return i
						}
					}
				}
			}
		}
	}
	return rep
}
