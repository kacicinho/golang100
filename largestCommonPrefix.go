func longestCommonPrefix(strs []string) string {
	pref := ""
	i := 0
	j := 1
	for i < len(strs[0]) {
		pref += string(strs[0][i])
		for j < len(strs) {
			if strs[j][:i+1] != pref {
				return pref[:len(pref)-1]
			}
			j++
		}
		j = 1
		i++
	}
	return ""
}