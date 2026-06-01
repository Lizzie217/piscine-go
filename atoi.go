package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}

	sign := 1
	result := 0
	start := 0

	// Handle sign
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	// Invalid if only sign exists
	if start == len(s) {
		return 0
	}

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		result = result*10 + int(s[i]-'0')
	}

	return result * sign
}
