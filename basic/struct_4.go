package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func main7() {
	root := &TreeNode{Value: 10}
	root.Left = &TreeNode{Value: 20}
	root.Right = &TreeNode{Value: 30}
  
	fmt.Println(root.Value)
	fmt.Println(root.Left.Value)
	fmt.Println(root.Right.Value)

}