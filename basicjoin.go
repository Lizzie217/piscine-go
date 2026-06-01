package piscine

func BasicJoin(elems []string) string {
	if len(elems) == 0 {
		return ""
	}

	result := ""
	for _, s := range elems {
		result += s
	}

	return result
}
