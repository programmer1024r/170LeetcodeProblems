/*
Use: return the index of the given target in the array if not exsits return -1
Input: sorted nums and target
Output: the indexTarget
*/
func search(nums []int, target int) int {

    indexTarget := -1

    if nums != nil {
        // initialize pointers
        low := 0
        high := len(nums) - 1
        middle := (low + high) / 2
        // go throw till the low pointer is bigger then the high
        for low <= high {

            if nums[middle] < target {

                low = middle + 1

            } else if nums[middle] > target {

                high = middle - 1

            } else {
                // found the target
                indexTarget = middle
                break
            }
            middle = (low + high) / 2

        }

    }


    return indexTarget
}
