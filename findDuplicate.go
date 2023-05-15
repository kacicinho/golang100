func containsDuplicate(nums []int) bool { 
    seen := make(map[int]int)

    for _, num := range nums {
        seen[num]++
        if(seen[num] > 1){
            return true
        }
    }

    return false
}


/* another approche using space complexity instead of time complexity 
import "sort"

func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
    
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            return true
        }
    }
    
    return false
}
*/