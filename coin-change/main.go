
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    
    // top down: go from 1 to amount
    for i := 1; i <= amount; i++ {
        best := math.MaxInt32
        for _, coin := range coins {
            if i-coin >= 0 && dp[i-coin] != -1 {
                best = min(best, dp[i-coin]+1)
            }
        }
        
        if best == math.MaxInt32 {
            dp[i] = -1
        } else {
            dp[i] = best
        }
    }
    
    return dp[amount]
}

func min(x, y int) int {
    if x < y {
        return x
    } else {
        return y
    }
}