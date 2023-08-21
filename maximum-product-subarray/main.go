func maxProduct(nums []int) int {
   maxP, minP, result:= nums[0], nums[0], nums[0]
   for _, v:= range nums[1:] {
       tempMax:=max(v, max(maxP*v, minP*v))
       minP=min(v, min(maxP*v, minP*v))

       maxP=tempMax
       result=max(result, maxP)
   }
   return result
}

func max(x, y int) int {
    if x>y {
        return x
    } 
    return y
}

func min(x, y int) int {
    if x<y{
        return x
    } 
    return y
}