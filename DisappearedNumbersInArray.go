/*
Use: find the missing numbers in array with range [1, n], n = array length
Input: array pointer, and the array size
Output: the missing numbers
 */
func findDisappearedNumbers(nums []int) []int {
    n := len(nums)
    hashTable :=  make(map[int]int) // The key will be the value
    var solution []int
    // Put nums in hash table 
    for i := 0; i < n ; i++ {
        hashTable[nums[i]] = nums[i]
    }
    // go throw numbers from 1 - n (len(nums)) and check if they are in the hash table
    for i := 1; i <= n; i++ {
        // the number is not in the hash table
        if hashTable[i] == 0 {
            solution = append(solution, i)
        }
    }
    return solution
}

// Better solution
func findDisappearedNumbers(nums []int) []int {
    n := len(nums)

	for i: = 0; i < n; i++ {
		nums[(nums[i] - 1) % n ] += n
	}

    out := []int{}

	for i : = 0; i < n; i++ {
        if nums[i] <= n {
			out = append(out,i+1)
		}
	}
	return out
}
