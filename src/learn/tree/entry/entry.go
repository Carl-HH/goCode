package main

import (
	"fmt"
	"goCode/src/learn/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if (myNode == nil || myNode.node == nil) {
		return
	}
	//包装成myTreeNode 才能调用postOrder方法

	left :=  myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3, Left: nil, Right: nil}
	root2 := tree.Node{3, &(tree.Node{}), nil}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Right.Right = tree.CreateTreeNode(3)
	fmt.Println(root)
	fmt.Println(root2)
	root.Print()

	//nil指针也能调用方法
	var pRoot *(tree.Node)
	pRoot.Set(3)

	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
