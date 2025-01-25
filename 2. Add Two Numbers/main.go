
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{}
	h := r
	rest := 0
	for {
		v := rest
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}

		h.Val = v % 10
		rest = v / 10

		if l1 == nil && l2 == nil && rest == 0 {
			break
		}
		h.Next = &ListNode{}
		h = h.Next
	}

	return r
}
