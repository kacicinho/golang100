/*
The complexity of this algorithm is O(n * log m), where n 
is the length of the input array and m is the length of the shortest string in the array.

The reason for this is that the algorithm performs a binary search on the shortest string in the array, 
with each iteration cutting the search range in half. This takes log m steps.

The binary search is performed n times, once for each string in the array. So the total time complexity is O(n * log m).

*/

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    prefix := strs[0]
    for i := 1; i < len(strs); i++ {
        for !strings.HasPrefix(strs[i], prefix) {
            prefix = prefix[:len(prefix)-1]
            if len(prefix) == 0 {
                return ""
            }
        }
    }
    return prefix
}