package piscine

func ListPushFront(l *List, data interface{}) {
	val := &NodeL{Data: data}
	if l.Tail == nil {
		l.Tail = l.Head
		l.Head = val
	} else {
		n := val
		n.Next, l.Head = l.Head, val
	}
}
