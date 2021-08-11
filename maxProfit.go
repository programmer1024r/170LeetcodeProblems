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
