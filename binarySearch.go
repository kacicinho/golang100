func search(nums []int, target int) int {
    l := 0
    r := len(nums)-1
    m := (l+r)/2

    for l <= r {
        if nums[m] > target {
            r-=1
            m = (l+r)/2
        } else if nums[m] < target{
            l+=1
            m = (l+r)/2
        } else {
            return m
        }
    }

    return -1

}