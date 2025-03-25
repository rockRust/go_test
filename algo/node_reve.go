package algo

import "test/entity"

// 反转链表
func ReverseNode(node entity.Node, prev *entity.Node) entity.Node {
	if node.Next == nil {
		node.Next = prev
		return node
	}
	next := *node.Next
	node.Next = prev
	prev = &node
	return ReverseNode(next, prev)
}
