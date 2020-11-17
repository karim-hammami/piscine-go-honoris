package piscine

func Compact(ptr *[]string) int {
	arrLength := 0
	for _, word := range *ptr {
		if word != "" {
			arrLength++
		}
	}
	s := make([]string, arrLength)
	i := 0
	for _, word := range *ptr {
		if word != "" {
			s[i] = word
			i++
		}
	}
	*ptr = s
	return arrLength
}
