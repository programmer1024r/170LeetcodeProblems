/*
Use: find the first (smallest) letter that bigger then the target
Input: sorted array of letters and the target
Output: first bigger then target letter
*/
func nextGreatestLetter(letters []byte, target byte) byte {

    solution := letters[0]

    for i := 0; i < len(letters); i++ {

        if letters[i] > target {
            solution = letters[i]
            break
        }
    }

    return solution
}
