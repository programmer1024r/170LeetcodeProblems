package main

import "fmt"

func main() {
    fmt.Println(countBits2(2))
}
/*
Use: count the number of ones in the binary representation of 
    the numbers in range of 0 <= number <= n
Input: n the range
Output: an array of the ones sum
*/
func countBits(n int) []int {

    var solution []int
    solution = append(solution, 0)

    for len(solution) <= n {

        length := len(solution)

        for i := 0; i < length; i++ {
            solution = append(solution, solution[i] + 1)
        }
    }

    return solution[:n + 1]
}

// Another implementation

/*
Use: count the number of ones in the binary representation of 
    the numbers in range of 0 <= number <= n
Input: n the range
Output: an array of the ones sum
*/
func countBits2(n int) []int {

    var dp []int
    dp = append(dp, 0)
    offset := 1

    for i := 1; i <= n; i++ {
        // Update the offset
        if offset * 2 == i {
            offset = i
        }
        // we use the last "set" of numbers Plus one because we are in the next "set"
        dp = append(dp, dp[i - offset] + 1)


    }

    return dp
}
