package piscine

func NRune(s string, n int) rune {
	str := []rune(s)
	if n > 0 && n <= StrLen(s) {
		return str[n-1]
	} else {
		return 0
	}
}
