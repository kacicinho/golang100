func rotateString(s string, goal string) bool {
    if len(s) != len(goal) {
        return false
    }

    rotated := s[1:] + string(s[0])

    return rotated == goal
}
