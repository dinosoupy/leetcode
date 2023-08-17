func productExceptSelf(nums []int) []int {
    ans:=make([]int, len(nums))
    //forward pass
    buffer := 1 
    for i, _:=range nums {
        ans[i] = buffer
        buffer*= nums[i]
    }
    //backward pass
    buffer = 1
    for j:=len(nums)-1; j>=0; j-- {
        ans[j]*=buffer
        buffer*=nums[j]
    }
    return ans
}