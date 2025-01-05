package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func (n *Node) Copy() *Node {
	return &Node{
		Val: n.Val,
	}
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	rh := head.Copy()
	ih := rh
	oh := head
	onMap := make(map[*Node]*Node, 1000)

	onMap[head] = ih
	for head.Next != nil {
		temp := head.Next.Copy()
		ih.Next, ih = temp, temp
		head = head.Next
		onMap[head] = ih
	}

	head, ih = oh, rh
	ih.Random = onMap[head.Random]
	for head.Next != nil {
		head = head.Next
		ih = ih.Next
		ih.Random = onMap[head.Random]
	}

	return rh
}
