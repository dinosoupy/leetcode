func longestConsecutive(nums []int) int {
    // make a map:
    // key: num from nums
    // value: bool indicating if key is in nums
    m:= make(map[int]bool, len(nums))
    for _, n:= range nums {
        m[n] = true
    }

    longest:= 0
    for n:= range m {
        fmt.Print(n)
        if !m[n-1] {
            length:= 1
            for {
                if !m[n+length] {
                    break
                } 
                length++
            } 
            if length>longest {
                longest = length
            }
        } 
    }
    return longest
}

