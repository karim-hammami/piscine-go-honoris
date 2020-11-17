package piscine

func LastRune(s string) rune {
	str := []rune(s)
	return str[StrLen(s)-1]
}
