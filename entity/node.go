package entity

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func (n Node) Print() {
	fmt.Println(n.Val)
	if n.Next != nil {
		n.Next.Print()
	}
}
