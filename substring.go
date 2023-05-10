func strStr(haystack string, needle string) int {
    n := len(haystack)
    m := len(needle)

    // Handle special cases
    if m == 0 {
        return 0
    }
    if m > n {
        return -1
    }

    // Precompute the bad character table
    badChar := make(map[byte]int)
    for i := 0; i < m-1; i++ {
        badChar[needle[i]] = m - i - 1
    }

    // Precompute the good suffix table
    goodSuffix := make([]int, m)
    suffix := make([]int, m)
    lastPrefixPos := m
    for i := m - 1; i >= 0; i-- {
        if isPrefix(needle, i+1) {
            lastPrefixPos = i + 1
        }
        suffix[m-1-i] = lastPrefixPos - i + m - 1
    }
    for i := 0; i < m-1; i++ {
        len := suffixLen(needle, i)
        goodSuffix[len] = m - i - 1 + len
    }

    // Search for the needle in the haystack
    i := m - 1
    j := m - 1
    for i < n && j >= 0 {
        if haystack[i] == needle[j] {
            i--
            j--
        } else {
            badCharShift := badChar[haystack[i]]
            goodSuffixShift := goodSuffix[j]
            if badCharShift > goodSuffixShift {
                i += badCharShift
            } else {
                i += goodSuffixShift
            }
            j = m - 1
        }
    }
    if j == -1 {
        return i + 1
    }
    return -1
}

// Check if the suffix of needle starting from pos is also a prefix
func isPrefix(needle string, pos int) bool {
    n := len(needle)
    for i := pos; i < n; i++ {
        if needle[i-pos] != needle[i] {
            return false
        }
    }
    return true
}

// Compute the length of the longest suffix of needle that is also a prefix
func suffixLen(needle string, pos int) int {
    n := len(needle)
    len := 0
    for i := pos; i >= 0 && needle[n-1-i] == needle[n-1-pos+len]; i-- {
        len += 1
    }
    return len
}