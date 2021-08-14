
/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

    hasCycle := false
    slowP :=  head
    fastP := head

    for fastP != nil && fastP.Next != nil {

        // update pointers
        slowP = slowP.Next
        fastP = fastP.Next.Next

        // check if the two pointers overlaped 
        if slowP == fastP {
            hasCycle = true
            break
        }

    }

    return hasCycle

}
