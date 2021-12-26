package linkedlist

//反转链表

func ReverseList( pHead *ListNode ) *ListNode {
	// write code here

	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	mid := pHead.Next
	high := mid.Next

	pHead.Next = nil

	for high != nil {
		mid.Next = pHead
		pHead = mid
		mid = high
		high = high.Next
	}
	mid.Next = pHead

	return mid
}