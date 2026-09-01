func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
		return ""
	}

	var output strings.Builder

	if len(strs[0]) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		v := strs[0][i]
		for _, w := range strs {
			if i == len(w) || w[i] != v {
				return output.String()
			}
		}
		output.WriteString(string(v))
	}

	return output.String()
}
