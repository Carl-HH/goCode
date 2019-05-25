package tree

import (
	"fmt"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

//为结构定义方法 ， 类似func  print(node treeNode)
func (node TreeNode) Print() {
	fmt.Print(node.Value, " ")
}

//使用指针作为方法接受者
func (node *TreeNode) Set(value int) {
	if node == nil {
		fmt.Println("node is nil")
	} else {
		node.Value = value
	}
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()

}

func CreateTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}
