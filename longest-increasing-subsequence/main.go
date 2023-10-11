func lengthOfLIS(nums []int) int {
    dp:=make([]int, len(nums))

    for idx:= range dp { dp[idx] = 1 } 

    for i:=1; i<len(nums); i++ { 
        for j:=0; j<i; j++ {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
    }

    max:= 0
    for q:=0; q<len(dp); q++ {
        if dp[q] > max {
            max = dp[q]
        }
    }
    return max

}
func max(x, y int) int {
    if x>y {
        return x
    } else {
        return y
    }
}