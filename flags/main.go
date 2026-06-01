package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func SortString(s string) string {
	r := []rune(s)

	for i := 0; i < len(r)-1; i++ {
		for j := i + 1; j < len(r); j++ {
			if r[i] > r[j] {
				r[i], r[j] = r[j], r[i]
			}
		}
	}

	return string(r)
}

func RemoveSpaces(s string) string {
	result := ""

	for _, ch := range s {
		if ch != ' ' {
			result += string(ch)
		}
	}

	return result
}

func PrintHelp() {
	PrintStr("--insert\n")
	PrintStr("  -i\n")
	PrintStr("\t This flag inserts the string into the string passed as argument.\n")
	PrintStr("--order\n")
	PrintStr("  -o\n")
	PrintStr("\t This flag will behave like a boolean, if it is called it will order the argument.\n")
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		PrintHelp()
		return
	}

	insert := ""
	order := false
	word := ""

	for _, arg := range args {
		if len(arg) >= 10 && arg[:10] == "--insert=" {
			insert = arg[10:]
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insert = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else {
			word = arg
		}
	}

	result := word + insert
	result = RemoveSpaces(result)

	if order {
		result = SortString(result)
	}

	PrintStr(result)
	z01.PrintRune('\n')
}
