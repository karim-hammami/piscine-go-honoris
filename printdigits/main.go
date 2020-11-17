package main

import "github.com/01-edu/z01"

func main() {
	var i rune
	for i = 48; i < 58; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}
