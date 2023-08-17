func maxArea(height []int) int {
    maxArea:=0
    left:=0
    right:=len(height)-1

    for left<right{
        a:=area(left, height[left], right, height[right])
        if a>maxArea {
            maxArea=a
        }
        if height[left]<height[right]{
            left++
        } else {
            right--
        }
    }
    return maxArea
}

func area(i1, h1, i2, h2 int) int {
    if h1<h2{
        return h1*(i2-i1)
    } else {
        return h2*(i2-i1)
    }
}