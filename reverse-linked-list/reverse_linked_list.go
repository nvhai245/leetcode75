package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// recursive
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
