func maxSubArray(nums []int) int {
    maxSum, currSum:= nums[0], nums[0]
    for right:=1; right<len(nums); right++{
        if currSum<0{
            currSum=0
        }
        currSum+=nums[right]
        if currSum>maxSum {
            maxSum=currSum
        }
    }
    return maxSum
}