func reverse(s []byte, start int, end int) {
    for start < end {
        s[start], s[end] = s[end], s[start]
        start++
        end--
    }
}

func reverseWords(s string) string {
    result := []byte(s)
    start := 0
    for i := 0; i < len(s); i++ {
        if s[i] == ' ' {
            reverse(result, start, i-1)
            start=i+1
        }
    }
    reverse(result, start, len(s)-1)
    return string(result)
}