func longestCommonSubsequence(text1 string, text2 string) int {
    dp := make([][]int, len(text1)+1)
    for row := range dp {
        dp[row] = make([]int, len(text2)+1)
    }

    for row:= 1; row<len(text1)+1; row++ {
        for col:= 1; col<len(text2)+1; col++ {
            // if chars match, add 1 + northwest element
            if text1[row-1] == text2[col-1] {
                dp[row][col] = 1 + dp[row-1][col-1]
            // if chars dont match, take max of north and west
            } else {
                dp[row][col] = max(dp[row-1][col], dp[row][col-1])
            }
        }
    }
    return dp[len(text1)][len(text2)]
}

func max(x, y int) int {
    if x > y {return x} else {return y}
}