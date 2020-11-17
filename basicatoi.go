package piscine

//func
func BasicAtoi(s string) int {
	on := 0
	c := 0
	as := []rune(s)
	for _, st := range as {
		for i := '0'; i < st; i++ {
			c++
		}
		on = on*10 + c
		c = 0
	}
	return on
}
