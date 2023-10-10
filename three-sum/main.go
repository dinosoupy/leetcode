func threeSum(nums []int) [][]int {
    var res [][]int
    // sort the array
    sort.Ints(nums)
    
    for i:=0; i<len(nums);i++ {
        // skip duplicates
        if i>0 && nums[i] == nums[i-1]{
            continue
        }

        // binary search for 0-nums[i]
        l, r:= i+1, len(nums)-1
        for l <= r {
            sum := nums[i] + nums[l] + nums[r]
            if sum < 0 {
                l++
            } else if sum > 0 {
                r--
            } else {
                res = append(res, []int{nums[i], nums[l], nums[r]})
                // skip duplicates in l and r
                for l<r && nums[l]==nums[l+1] { l++ }
                for l<r && nums[r]==nums[r-1] { r-- }
                // move one away from current l, r since those are the last duplicates
                l++
                r--
            }
        }
    }
}