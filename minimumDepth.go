/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {

    // Base case
    if root == nil {return 0}

    queue := make([]*TreeNode, 0)
    queue = append(queue, root)

    level := 0
    found := false

    for len(queue) != 0 && found == false {

        level_size := len(queue)
        for i := 0; i < level_size; i++ {
            // dequeue
            curr := queue[0]
            queue = queue[1:]

            // first leaf
            if curr.Left == nil && curr.Right == nil {
                found = true
            }

            if curr.Left != nil {
                queue = append(queue, curr.Left)

            }

            if curr.Right != nil {
                queue = append(queue, curr.Right)

            }
        }

        level++

    }

 return level
}
