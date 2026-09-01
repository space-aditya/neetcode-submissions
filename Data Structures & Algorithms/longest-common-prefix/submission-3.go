func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
		return ""
	}

	prefix := []rune{}

	minWord := strs[0]

	for _, v := range strs {
		if len(v) < len(minWord) {
			minWord = v
		}
	}

	for i, v := range minWord {
		for _, w := range strs {
			if w[i] != byte(v) {
				return string(prefix)
			}
		}
		prefix = append(prefix, v)
	}

	return string(prefix)
}
