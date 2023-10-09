func sortArray(nums []int) []int {

    var merge func(left, right []int) []int
    merge = func(left, right []int) []int {
        size, i, j:= len(left)+len(right), 0, 0
        slice:= make([]int, size)

        for k:=0; k<size; k++ {
            if i > len(left)-1 && j <= len(right)-1 {
                slice[k] = right[j]
                j++
            } else if j > len(right)-1 && i <= len(left)-1 {
                slice[k] = left[i]
                i++
            } else if left[i] < right[j] {
                slice[k] = left[i]
                i++
            } else {
                slice[k] = right[j]
                j++
            }
        }
        return slice
    } 

    if len(nums) < 2 {
        return nums
    }
    mid:= len(nums)/2
    return merge(sortArray(nums[:mid]), sortArray(nums[mid:]))
}