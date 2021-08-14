/*
Use: remove all elements with the given value from the linkedlist
Input: head, value 
Output: the linked list without the nodes asked
*/
func removeElements(head *ListNode, val int) *ListNode {

    if head != nil {

        head.Next = removeElements(head.Next, val)

        if head.Val == val {
            head = head.Next
        }
    }

    return head
}


