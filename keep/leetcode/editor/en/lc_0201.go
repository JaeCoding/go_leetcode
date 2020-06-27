package main

//编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
//
// 示例1:
//
//
// 输入：[1, 2, 3, 3, 2, 1]
// 输出：[1, 2, 3]
//
//
// 示例2:
//
//
// 输入：[1, 1, 1, 1, 2]
// 输出：[1, 2]
//
//
// 提示：
//
//
// 链表长度在[0, 20000]范围内。
// 链表元素在[0, 20000]范围内。
//
//
// 进阶：
//
// 如果不得使用临时缓冲区，该怎么解决？
// Related Topics 链表

//leetcode submit region begin(Prohibit modification and deletion)
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

func removeDuplicateNodes1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 声明并赋值一个map[int,bool] 初始化已有head
	occurred := map[int]bool{head.Val: true}
	// 声明 pos
	pos := head
	// 只要下一个不为nil
	for pos.Next != nil {
		cur := pos.Next
		if !occurred[cur.Val] {
			// 如果下一个值不在map，则将其加入map，并位移
			occurred[cur.Val] = true
			pos = pos.Next
		} else {
			// 否则当前直接指向下下
			pos.Next = pos.Next.Next
		}
	}
	pos.Next = nil
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
