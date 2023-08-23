func twoSum(numbers []int, target int) []int {
    low, high:=0, len(numbers)-1
    for low<high {
        if numbers[low]+numbers[high]>target {
            high--
        }
        if numbers[low]+numbers[high]<target {
            low++
        }
        if numbers[low]+numbers[high]==target {
            break
        }
    }
    return []int{low+1, high+1}
}