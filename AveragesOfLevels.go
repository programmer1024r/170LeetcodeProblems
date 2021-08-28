/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {

    var solution []float64
    var sum int

    // edge case
    if root == nil {
        return solution
    }

    queue := make([]*TreeNode, 0) // initailizing empty queue
    queue = append(queue, root)

    for len(queue) != 0 {

        level_size := len(queue)
        sum = 0

        for i := 0; i < level_size; i++ {
            // dequeue
            curr := queue[0]
            queue = queue[1:]
            sum += curr.Val

            // update level
            if curr.Left != nil {
                queue = append(queue, curr.Left)
            }

            if curr.Right != nil {
                queue = append(queue, curr.Right)
            }

        }
        solution = append(solution, float64(sum) / float64(level_size))
    }
    return solution
}
