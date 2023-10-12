func wordBreak(s string, wordDict []string) bool {
    dp := make([]bool, len(s)+1)

    dp[0] = true

    for i:=0; i<len(s)+1; i++ {
        if dp[i] {
            for _, word := range wordDict {
                if i+len(word) <= len(s) && s[i:i+len(word)] == word {
                    dp[i+len(word)] = true
                }
            }
        }
    }
    return dp[len(s)]
}