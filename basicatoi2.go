package piscine

//func
func BasicAtoi2(s string) int {
	onumber := 0
	c := 0
	checker := true
	as := []rune(s)
	for _, word := range as {
		if byte(word) >= 48 && byte(word) <= 57 {
			for i := '0'; i < word; i++ {
				c++
			}
			onumber = onumber*10 + c
			c = 0
		} else {
			checker = false
		}
	}
	if checker {
		return onumber
	} else {
		return 0
	}
}
