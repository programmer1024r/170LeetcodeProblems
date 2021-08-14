/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {

    fastP := head
    slowP := head

    for fastP != nil && fastP.Next != nil {

        // Update the pointers 
        fastP = fastP.Next.Next
        slowP = slowP.Next
    }

    return slowP
}
