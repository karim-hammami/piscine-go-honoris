package piscine

// func strRev
func StrRev(s string) string {
	var reverse = ""
	for _, a := range s {
		reverse = string(a) + reverse
	}
	return reverse
}
