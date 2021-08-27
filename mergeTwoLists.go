/*
Use: merge two linked lists
Input: two sorted lists
Output: the final sorted list
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

    l1Ptr := l1
    l2Ptr := l2

    // edge cases
    if l1 == nil && l2 == nil{
        return nil

    } else if l1 == nil {
        return l2

    } else if l2 == nil {
        return l1
    }


    var head *ListNode

    // choose the head
    if l1Ptr.Val < l2Ptr.Val {
        head = l1Ptr
        l1Ptr = l1Ptr.Next

    } else {
        head = l2Ptr
        l2Ptr = l2Ptr.Next
    }

    curr := head

    // go throw the lists till one of the pointers are nil
    for l1Ptr != nil && l2Ptr != nil {

        if l1Ptr.Val < l2Ptr.Val {
            curr.Next = l1Ptr
            l1Ptr = l1Ptr.Next

        } else {
            curr.Next = l2Ptr
            l2Ptr = l2Ptr.Next
        }

        curr = curr.Next
    }

    // add the not nil pointer to be the contine of the linked list
    if l1Ptr != nil {
        curr.Next = l1Ptr

    } else {
        curr.Next = l2Ptr
    }

    return head
}
