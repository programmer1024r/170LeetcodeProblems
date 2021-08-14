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

// Second time solution

/*
Use: find the maximum sum of a contiguous subarray
Input: array 
Output: the maximum sum
*/
func maxSubArray(nums []int) int {
    currentSum := 0
    maxSum := nums[0] // max number

    for i := 0; i < len(nums); i++ {
        // update current sum
        if currentSum < 0 {
            currentSum = nums[i]

        } else { // else continue adding the sum
            currentSum += nums[i]
        }

        // update max sum
        if currentSum > maxSum {
            maxSum = currentSum
        }
    }

    return maxSum
}
