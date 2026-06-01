package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ProgramPath := os.Args[0]
	ProgramName := ""
	for i := len(ProgramPath) - 1; i >= 0; i-- {
		if ProgramPath[i] == '/' {
			ProgramName = ProgramPath[i+1:]
			break
		}
	}
	if ProgramName == "" {
		ProgramName = ProgramPath
	}
	for _, char := range ProgramName {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
