func minSubArrayLen(target int, nums []int) int {
    ans := math.MaxInt32
    left, sum:= 0, 0
    for right:=0; right<len(nums); right++ {
        sum += nums[right]
        for sum >= target {
            if right-left+1 < ans {
                ans = right-left+1
            } 
            sum -= nums[left]
            left++ 
        } 
    }
    if ans == math.MaxInt32 {
        return 0
    } else {
        return ans
    }
}