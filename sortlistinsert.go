package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		return nil
	}
	curr := l
	node := l
	n := &NodeI{Data: data_ref}
	var previous *NodeI = nil
	for curr.Next != nil {
		if curr.Data > data_ref {
			if previous == nil {
				node = n
				n.Next = curr
			} else {
				previous.Next = n
				n.Next = curr
			}
			return node
		}
		previous = curr
		curr = curr.Next
	}
	if curr.Data > data_ref {
		previous.Next = n
		n.Next = curr
	} else {
		curr.Next = n
	}
	return node
}
