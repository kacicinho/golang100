func reverseString(s []byte)  []byte{
    i := 0 
    j := len(s)-1

    for i<j {
        temp := s[j]
        s[j] = s[i]
        s[i] = temp
        i++
        j--
    }

    return s
}


//optimized version 

func reverseString(s []byte)  {
	for i, j := 0, len(s)-1; i<j; i, j = i+1, j-1 {
	   s[i], s[j] = s[j], s[i]
   }
}