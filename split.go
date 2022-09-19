package piscine

func Split(s, sep string) []string {
	rep := []string{}
	mot := ""
	a := 0
	for i := 0; i < len(s); i++ {
		if s[i] == sep[0] && len(sep) > 1 {
			mot_sep := string(sep[0])
			for j := 1; j < len(sep); j++ {
				mot_sep += string(s[i+j])
				a = j
			}
			if mot_sep == sep {
				i = i + a
				rep = append(rep, mot)
				mot = ""
			} else {
				mot += string(s[i])
			}
		} else {
			mot += string(s[i])
		}
	}
	rep = append(rep, mot)
	return rep
}
