/*
Use: find the maximum sum of a contiguous subarray
Input: array 
Output: the maximum sum
*/
func maxSubArray(nums []int) int {
    // minimum value of integer
    maxSum := -int(^uint(0) >> 1) - 1

    if nums != nil {

        totalSum := 0
        // need to go throw the array
        for i := 0; i < len(nums); i++ {

            // if the sum that we have is negative then update it to the current number
            if totalSum < 0 {
                totalSum = nums[i]

            } else {
                // add the current number to the total sum
                totalSum += nums[i]
            }
            // Update max sum
            if totalSum > maxSum {
                maxSum = totalSum
            }

        }
    }

    return maxSum

}
