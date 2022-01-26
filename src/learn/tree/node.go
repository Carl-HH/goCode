package tree

import (
	"fmt"
)

type Node struct {
	/**
	  首字母大写 则访问权限为public
	*/
	Value       int
	Left, Right *Node
}

//为结构定义方法 ， 类似func  print(node treeNode)
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

//使用指针作为方法接受者,只有使用指针才可以改变结构内容
//nil指针也可以调用方法
func (node *Node) Set(value int) {
	if node == nil {
		fmt.Println("node is nil")
	} else {
		node.Value = value
	}
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()

}

func (node *Node) postOrder() {
	if node == nil {
		return
	}
	node.postOrder()
}

func CreateTreeNode(value int) *Node {
	return &Node{Value: value}
}
