package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法1 迭代
func reverseList1(head *ListNode) *ListNode {

	pre := (*ListNode)(nil)
	ListNodeCurrent := head

	// 判断ListNodeCurrent 是否为空
	for ListNodeCurrent != (*ListNode)(nil) {
		// 如果不为空
		// 将ListNodeCurrent.Next 保存到tmp
		NodeTmp := ListNodeCurrent.Next

		// 将将ListNodeCurrent 赋值为pre 的值
		ListNodeCurrent.Next = pre

		// pre 与 ListNodeCurrent 向后移一格
		pre = ListNodeCurrent
		ListNodeCurrent = NodeTmp
	}
	// ListNodeCurrent 为空则到达列表结尾 返回最后的node
	return pre
}

// 方法2 递归
func reverseList2(head *ListNode) *ListNode {
	if head == (*ListNode)(nil) || head.Next == (*ListNode)(nil) {
		return head
	}

	p := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = (*ListNode)(nil)
	return p

}
