package algo

import (
	"fmt"
	"test/entity"
	"testing"
)

func TestReverseNode(t *testing.T) {
	var prev *entity.Node
	for i := 0; i < 3; i++ {
		node := entity.Node{
			Val: i,
		}
		node.Next = prev
		prev = &node
	}
	prev.Print()
	res := ReverseNode(*prev, nil)
	fmt.Println("after reverse")
	res.Print()
}
