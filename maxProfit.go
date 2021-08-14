/*
Use: find the max profit to buy and sell in a given array of prices 
Input: array of prices 
Output: max profit
*/
func maxProfit(prices []int) int {

    maxProfit := 0

    if prices == nil {
        return maxProfit
    } else {

        valleyP := 0

        for i := 1; i < len(prices); i++ {
            peakP := i
            // Update the valley if there is a lower valley
            if prices[valleyP] > prices[peakP] {
                valleyP = peakP
            }
            // Update the max profit
            if maxProfit < prices[peakP] - prices[valleyP] {
                maxProfit = prices[peakP] - prices[valleyP]
            }
        }
    }
    return maxProfit
}

// Second time solving

/*
Use: find the max profit to buy and sell in a given array of prices 
Input: array of prices 
Output: max profit
*/
func maxProfit(prices []int) int {
    left := 0
    right := 1
    maxProfit := 0

    for i := 1; i < len(prices); i++ {
        // jump left to right if right is smallest then him
        if prices[right] < prices[left] {
            left = right
        }
        // update maxProfit if current profit is bigger
        if prices[right] - prices[left] > maxProfit {
            maxProfit = prices[right] - prices[left]
        }
        // update right
        right = i + 1

    }

    return maxProfit
}
