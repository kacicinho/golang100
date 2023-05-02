/*

There is a function signFunc(x) that returns:

1 if x is positive.
-1 if x is negative.
0 if x is equal to 0.
You are given an integer array nums. Let product be the product of all values in the array nums.

Return signFunc(product).

*/

func arraySign(nums []int) int {

    product := 1

    for i:=0; i<len(nums); i++ {
        if (nums[i] == 0) {
            return 0
        }
        if (nums[i] < 0){
            product *=-1
        }
    }

    if (product > 0) {
        return 1
    } else  {
        return -1
    }
}
