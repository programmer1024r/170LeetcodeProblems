/*
Use: Find the number of distinct ways to get to the target number by 1 or 2
Input: target number (int)
Output: number of combinations 
*/
func climbStairs(n int) int {

    // on the n stair you have one option and it's not to move, and on the n -1 is always to jump 1 means one option 
    // if n = 5, the nextStair is on the 4th stair and the afterNextStair is the 5th stair
    nextStair := 1
    afterNextStair := 1

    // going backwards through the  combinations
    // every time you go to the next stair or to the next next stair. means all the combinations from your stair +1 and all the combinations from your stair + 2
    for i :=  2; i < n; i++ {
        currentStair := nextStair + afterNextStair
        afterNextStair = nextStair
        nextStair = currentStair

    }

    return currentStair
}
