package piscine

func Join(strs []string, sep string) (s string) {
	length := 0
	for range strs {
		length++
	}

	for i, word := range strs {
		s += word
		if i != length-1 {
			s += sep
		}
	}
	return
}
