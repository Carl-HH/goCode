package main

import (
	"fmt"
	"learn/tree"
)

func main() {
	var root tree.TreeNode

	root = tree.TreeNode{Value: 3, Left: nil, Right: nil}
	root2 := tree.TreeNode{3, &(tree.TreeNode{}), nil}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)
	root.Right.Right = tree.CreateTreeNode(3)
	fmt.Println(root)
	fmt.Println(root2)
	root.Print()

	//nil指针也能调用方法
	var pRoot *(tree.TreeNode)
	pRoot.Set(3)

	root.Traverse()
}
