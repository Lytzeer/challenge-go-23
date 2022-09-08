package piscine

func SortIntegerTable(table []int) {
	for i := 1; i < len(table); i++ {
		for j := 0; j < len(table)-1; j++ {
			val_i := table[i]
			val_j := table[j]
			if table[i] < table[j] {
				table[i] = val_j
				table[j] = val_i
			}
		}
	}
}
