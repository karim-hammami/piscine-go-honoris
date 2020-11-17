package piscine

func Sqrt(nb int) int {
	if nb == 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	i := 0
	for ; i <= nb/2; i++ {
		if i*i > nb {
			return 0
		} else if i*i == nb {
			return i
		}
	}
	return 0
}
