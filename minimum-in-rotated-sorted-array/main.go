func findMin(nums []int) int {
   low, high := 0, len(nums)-1
   for low<high {
       mid:=(low+high)/2
       // inflection point is to the right of mid
       if nums[mid]>nums[high] {
           low=mid+1
       } else { // inflection point is to the left of mid
           high=mid
       }
   }
   return nums[low]
}

