package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	ite := l
	i := 0
	for ite != nil {
		if i == pos {
			return ite
		} else {
			i++
			ite = ite.Next
		}
	}
	return nil
}
