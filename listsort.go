package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}

	node := l
	notSorted := true
	for notSorted {
		curr := node
		var previous *NodeI = nil
		notSorted = false
		for curr.Next != nil {
			if curr.Data > curr.Next.Data {
				notSorted = true
				next := curr.Next
				if previous == nil {
					node = curr.Next
				} else {
					previous.Next = next
				}
				if curr.Next.Next == nil {
					curr.Next = nil
				} else {
					curr.Next = curr.Next.Next
				}
				next.Next = curr
				previous = next
			} else {
				previous = curr
				curr = curr.Next
			}
		}
	}
	return node
}
