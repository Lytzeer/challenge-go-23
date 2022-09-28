package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	courant := l.Head
	precedent := l.Head
	for courant != nil && courant.Data == data_ref {
		l.Head = courant.Next
		courant = l.Head
	}
	for courant != nil {
		for courant != nil && courant.Data != data_ref {
			precedent = courant
			courant = courant.Next
		}
		if courant == nil {
			return
		}
		precedent.Next = courant.Next
		courant = precedent.Next
	}
}
