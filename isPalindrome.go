/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {

    var halfList []int
    isPalindrome := true
    fastP := head
    slowP := head

    if head != nil {

        // Go through half of the linked-list
        for fastP != nil && fastP.Next != nil {
            // saving the value
            halfList = append(halfList, slowP.Val)
            // Update the pointers 
            fastP = fastP.Next.Next
            slowP = slowP.Next
        }
        // in case the linked list with odd length
        if fastP != nil && fastP.Next == nil {
            slowP = slowP.Next
        }
        // Go through the second half of the linked list
        for i := len(halfList) - 1; i >= 0 && slowP != nil; i -- {

            /* if the first half is not identical to 
            the second then this is not palindrome */
            if halfList[i] != slowP.Val {
                isPalindrome = false
            }

            slowP = slowP.Next
        }

    }
    return isPalindrome
}

// recursion solution
var tmp *ListNode

func isPalindrome(head *ListNode) bool {
	tmp = head

	return check(head)
}

func check(p *ListNode) bool {
	if p == nil {
		return true
	}

	isPal := false
	if check(p.Next) && tmp.Val == p.Val {
		isPal = true
	}
	tmp = tmp.Next

	return isPal
}
