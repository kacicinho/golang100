func majorityElement(nums []int) int {
    myMap := make(map[int]int)
    size := len(nums)/2

    for _, num := range nums {
        value, exists := myMap[num]
        if exists {
            value +=1
            myMap[num] = value
        } else {
            myMap[num] = 1
        }

        if myMap[num] > size {
            return num
        }
    }

    return -1
}