func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    // Pointer for tracking the position of unique elements
    uniqueIndex := 0

    // Iterate through the array starting from the second element
    for i := 1; i < len(nums); i++ {
        // If the current element is different from the previous unique element,
        // update the unique index and copy the current element to the unique position
        if nums[i] != nums[uniqueIndex] {
            uniqueIndex++
            nums[uniqueIndex] = nums[i]
        }
    }

    // Return the number of unique elements (unique index + 1)
    return uniqueIndex + 1
}
