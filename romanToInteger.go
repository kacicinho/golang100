func romanToInt(s string) int {
    romanToIntMap := map[byte]int {
        'I' : 1,
        'V' : 5, 
        'X' : 10,
        'L' : 50,
        'C' : 100,
        'D' : 500,
        'M' : 1000,
    }

    result := 0

    for i:=0; i<len(s); i++ {
        currentValue := romanToIntMap[s[i]]
        // Check if we need to subtract the previous value
        if i > 0 && romanToIntMap[s[i-1]] < currentValue {
            result -= romanToIntMap[s[i-1]]
            result += currentValue - romanToIntMap[s[i-1]]
        } else {
            result += currentValue
        }
    }

    return result
}