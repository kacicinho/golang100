/*

this one is O(n) but the other one is O(log(n)) since it iterate over half of the array
func isPalindrome(x int) bool {
    y := strconv.Itoa(x)
    ok := true

    for i:=0; i<len(y); i++ {
        ok = ok && (y[i] == y[len(y)-i-1])
    }
    
    return ok
}
*/

func isPalindrome(x int) bool {
	
	if x < 0 || (x % 10 == 0 && x != 0) {
        return false
    }

	reverse := 0
	for x > reverse {
		reverse = reverse * 10 + x%10 
		x/=10
	}

	return reverse == x || x == reverse/10 
}