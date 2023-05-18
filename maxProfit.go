func max(a int, b int) int{
    if a>b {
        return a
    }else{
        return b
    }
}



func maxProfit(prices []int) int {
    l, r := 0, 1 
    maxProfit := 0
    temp := 0
    for r<len(prices) {
        if prices[l] < prices[r]{
            temp = prices[r] - prices[l]
            maxProfit = max(maxProfit, temp)
        } else {
            l = r
        }
        r++
    }
    return maxProfit
}