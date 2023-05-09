/*
In this version, we use a hash table (map) to store each number in the input array nums along with its index.
We loop through the array once and for each number, we calculate its complement (the target minus the current number)
 and check if it exists in the hash table. 
 If the complement exists in the hash table, we return the indices of the two numbers that add up to the target. 
 If the complement does not exist in the hash table, we add the current number and its index to the hash table.

By using a hash table, we can achieve an average time complexity of O(n)
*/

func twoSum(nums []int, target int) []int {
    hashTable := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if index, ok := hashTable[complement]; ok {
            return []int{index, i}
        }
        hashTable[num] = i
    }
    return nil
}
