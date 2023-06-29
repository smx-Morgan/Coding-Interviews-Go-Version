package src

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	res := reversePrint(head.Next)

	return append(res, head.Val)
}
