func canSplitArray(nums []int, m int) bool {
    if len(nums) <= 2 {
        return true
    } else {
        dp := make(map[string]bool, len(nums))
        return check(0, len(nums)-1, m, nums, dp)
    }
}

func check(start int , end int, m int, nums []int, dp map[string]bool) bool {
    sub := string(start)+string(end)
    if start == end {
        return true
    } else if ans, ok := dp[sub]; ok {
        return ans
    } else if ok := isValid(start, end, m, nums); !ok {
        return false
    } else {
        ans := false
        for i:=start; i<end; i++ {
            subLeft:=check(start, i, m, nums, dp)
            subRight:=check(i+1, end, m, nums, dp)

            if subLeft && subRight {
            ans = true
            }
        }

        dp[string(start)+string(end)] = ans
        return ans
    }
}

func isValid(start int, end int, m int, nums []int) bool {
    sum:=0
    for i:=start; i<=end; i++ {
        sum+=nums[i]
    }
    return sum>=m
}
