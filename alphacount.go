package piscine

func AlphaCount(str string) int {
	count := 0
	for _, letter := range str {
		if letter > 64 && 91 > letter || letter > 96 && 123 > letter {
			count++
		}
	}
	return count
}
