/*
Use: find the peak of the mountain (max value) but it located like that [0, 1, 2, 3, 2, 0]
Input: array that is a mountain 
Output: the peak index
*/
func peakIndexInMountainArray(arr []int) int {

    max := arr[0]
    maxIndex := 0

    for i := 0; i < len(arr); i++ {

        if max < arr[i] {
            max = arr[i]
            maxIndex = i
        }

    }
    return maxIndex
}
// add the follow up of O(log n)
/*
Use: find the peak of the mountain (max value) but it located like that [0, 1, 2, 3, 2, 0]
Input: array that is a mountain 
Output: the peak index
*/
func peakIndexInMountainArray(arr []int) int {

    right := len(arr) - 1
    left := 0
    var middle int

    // run till left == right
    for left < right {

        middle = (left + right) / 2

        // going down
        if arr[middle] > arr[middle + 1] {
            right = middle

        // going up
        } else {
            left = middle + 1
        }
    }

    return left
}
