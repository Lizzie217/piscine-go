package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	n := 0

	if s == "" {
		return -1
	}

	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return -1
		}
		n = n*10 + int(ch-'0')
	}

	return n
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	upper := false
	start := 0

	if args[0] == "--upper" {
		upper = true
		start = 1
	}

	for i := start; i < len(args); i++ {
		n := Atoi(args[i])

		if n >= 1 && n <= 26 {
			letter := rune('a' + n - 1)

			if upper {
				letter = rune('A' + n - 1)
			}

			z01.PrintRune(letter)
		} else {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
}
