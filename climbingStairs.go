/*
Sure, let me explain why we use the recurrence relation dp[i] = dp[i-1] + dp[i-2] to compute the number of ways to climb the staircase with i steps.

Let's start with the base cases: if there is only one step, there is only one way to climb it, which is by taking one step. If there are two steps, there are two ways to climb it: either take one step at a time twice, or take two steps at once. Therefore, dp[1] = 1 and dp[2] = 2.

Now, let's consider the general case where we have i steps. We can reach the ith step either by taking one step from the i-1th step, or by taking two steps from the i-2th step. If we take one step from the i-1th step, then the number of ways to climb to the ith step is the same as the number of ways to climb to the i-1th step, which is dp[i-1]. If we take two steps from the i-2th step, then the number of ways to climb to the ith step is the same as the number of ways to climb to the i-2th step, which is dp[i-2]. Therefore, the total number of ways to climb to the ith step is the sum of these two possibilities, which gives us dp[i] = dp[i-1] + dp[i-2].

By using this recurrence relation, we can compute the number of ways to climb the staircase for any number of steps. The time complexity of this algorithm is O(n), where n is the number of steps, since we need to compute dp[i] for every i from 3 to n.
*/



func climbStairs(n int) int {
    if n <= 2 {
        return n
    }

    way := make([]int, n+1)
    way[1], way[2] = 1, 2

    for i:=3; i<=n; i++ {
        way[i] = way[i-1] + way[i-2] 
    }

    return way[n]
}