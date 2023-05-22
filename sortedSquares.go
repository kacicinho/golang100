func sortedSquares(nums []int) []int {
    left := 0
    n :=  len(nums)
    right := len(nums) - 1 
    index := len(nums) - 1  
    result := make([]int, n)

    if len(nums)<=0 {
        return result
    }

    for left <= right {
        leftValue := nums[left] * nums[left]
        rightValue := nums[right] * nums[right]

        if leftValue > rightValue {
            result[index] = leftValue
            left+=1
        } else {
            result[index] = rightValue
            right-=1
        }

        index-=1
    }

    return result 
}