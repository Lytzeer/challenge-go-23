package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	val := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = val
		l.Tail = val
	} else {
		l.Tail.Next = val
		l.Tail = val
	}
}
