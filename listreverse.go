package piscine

func ListReverse(l *List) {
	noeud_courant := l.Head
	r := l.Head
	r = nil
	for noeud_courant != nil {
		noeud_suivant := noeud_courant.Next
		noeud_courant.Next = r
		r = noeud_courant
		noeud_courant = noeud_suivant
	}
	l.Head, l.Tail = l.Tail, l.Head
}
