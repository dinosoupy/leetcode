func climbStairs(n int) int {
    zero, one := 1, 1
    var next int

    for i := 1; i < n; i++ {
        next = zero + one
        zero = one
        one = next
    }

    return one
}