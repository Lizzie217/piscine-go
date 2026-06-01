package main

import (
	"os"
)

func atoi(s string) (int64, bool) {
	if s == "" {
		return 0, false
	}

	sign := int64(1)
	i := 0

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	var n int64

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		n = n*10 + int64(s[i]-'0')
	}

	return n * sign, true
}

func itoa(n int64) string {
	if n == 0 {
		return "0"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	var s string
	for n > 0 {
		s = string('0'+(n%10)) + s
		n /= 10
	}

	return sign + s
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := atoi(os.Args[1])
	op := os.Args[2]
	b, ok2 := atoi(os.Args[3])

	if !ok1 || !ok2 {
		return
	}

	var res int64

	switch op {
	case "+":
		// overflow-safe add
		if (b > 0 && a > 9223372036854775807-b) ||
			(b < 0 && a < -9223372036854775808-b) {
			return
		}
		res = a + b

	case "-":
		if (b < 0 && a > 9223372036854775807+b) ||
			(b > 0 && a < -9223372036854775808+b) {
			return
		}
		res = a - b

	case "*":
		if a != 0 && b != 0 {
			if a > 9223372036854775807/b || a < -9223372036854775808/b {
				return
			}
		}
		res = a * b

	case "/":
		if b == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		res = a / b

	case "%":
		if b == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		res = a % b

	default:
		return
	}

	os.Stdout.WriteString(itoa(res) + "\n")
}
