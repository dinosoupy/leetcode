func maxProfit(prices []int) int {
    profit := 0 
    min := prices[0]

    for _, v := range prices {
        if v < min {
            min = v
        } else if v - min > profit {
            profit = v - min
        }
    }
    return profit
}