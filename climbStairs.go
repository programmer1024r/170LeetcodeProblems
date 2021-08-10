/*
Use: Find the number of distinct ways to get to the target number by 1 or 2
Input: target number (int)
Output: number of combinations 
*/
func climbStairs(n int) int {
    var combinationsPerPlace []int
    for i := 0; i < n; i++ {
        combinationsPerPlace = append(combinationsPerPlace, 0)
    }
    // on the n stair you have one option and it's not to move, and on the n -1 is always to jump 1 means one option 
    nIndex := n - 1
    combinationsPerPlace[nIndex] = 1
    combinationsPerPlace[nIndex - 1] = 1

    // going backwards through the array of combinations
    // every time you go to the next stair or to the next next stair. means all the combinations from your stair +1 and all the combinations from your stair + 2
    for i := nIndex - 2; i >= 0; i-- {
        combinationsPerPlace[i] = combinationsPerPlace[i + 1] + combinationsPerPlace[i + 2]
    }

    return combinationsPerPlace[0]
}
