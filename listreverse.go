package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	current := l.Head
	tempList := current
	next := &NodeL{}
	previous := &NodeL{}
	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	l.Head = previous
	l.Tail = tempList
	l.Tail.Next = nil
}
