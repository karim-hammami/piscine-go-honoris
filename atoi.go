package piscine

//func
func Atoi(s string) int {
	onumber := 0
	c := 0
	checker := true
	as := []rune(s)
	pl := 1

	if s != "" {
		if byte(as[0]) == 45 {
			pl = -1
			as[0] = '0'
		} else if byte(as[0]) == 43 {
			as[0] = '0'
		}
	}
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
		return onumber * pl
	} else {
		return 0
	}
}
