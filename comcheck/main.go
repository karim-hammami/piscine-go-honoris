package main

import (
	"fmt"
	"os"
)

func main() {
	argsLength := 0
	for range os.Args {
		argsLength++
	}
	if argsLength == 1 {
		return
	} else {
		arguments := os.Args[1:]
		for _, arg := range arguments {
			if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
				fmt.Println("Alert!!!")
				return
			}
		}

	}
}
