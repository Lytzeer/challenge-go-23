package piscine

func ListPushFront(l *List, data interface{}) {
	val := &NodeL{Data: data}
	if l.Tail == nil {
		l.Head = val
		l.Head = l.Head
	} else {
		n := val
		n.Next, l.Head = l.Head, n
	}
}
