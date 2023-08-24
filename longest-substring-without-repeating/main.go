func lengthOfLongestSubstring(s string) int {
    start:=0
    end:=0
    max:=0
    m:=make(map[byte]int)

    for end<len(s) {
        if _, ok:=m[s[end]]; !ok {
            m[s[end]]=end
            end++
            if len(m)>max {
                max=len(m)
            }
        } else {
            delete(m, s[start])
            start++
        }
    }
    return max
}