func lengthOfLastWord(s string) int {

    count := 0 

    
    for i := len(s) - 1; i >= 0 && s[i] == ' '; i-- {
        s = s[:i]
    }
    
    for i := len(s); i<len(s); i-- {
        if s[i] == ' ' {
            count = i
            break
        }
    }

    return count
}