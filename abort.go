package piscine

func Sort(table []int) {
	i := 1
	for i < 5 {
		if table[i-1] > table[i] {
			tmp := table[i]
			table[i] = table[i-1]
			table[i-1] = tmp
			i = 1
		} else {
			i++
		}
	}
}

func Abort(a, b, c, d, e int) int {
	tmp := []int{a, b, c, d, e}

	Sort(tmp)

	return tmp[2]
}
