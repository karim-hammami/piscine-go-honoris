package piscine

func ListSize(l *List) int {
	listSize := 0
	for l.Head != nil {
		listSize++
		l.Head = l.Head.Next
	}
	return listSize
}
