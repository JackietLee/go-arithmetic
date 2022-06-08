package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reConstructBinaryTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) != len(inorder) || len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	rootIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			rootIndex = i
		}
	}
	inL, inR := inorder[:rootIndex], inorder[rootIndex+1:]
	preL, preR := preorder[1:rootIndex+1], preorder[rootIndex+1:]
	left := reConstructBinaryTree(preL, inL)
	right := reConstructBinaryTree(preR, inR)
	t := new(TreeNode)
	t.Val = rootVal
	t.Left = left
	t.Right = right
	return t
}

func printPre(root *TreeNode) {
	if root != nil {
		fmt.Printf("%d ", root.Val)
		printPre(root.Left)
		printPre(root.Right)
	}
}

func printIn(root *TreeNode) {
	if root != nil {
		printIn(root.Left)
		fmt.Printf("%d ", root.Val)
		printIn(root.Right)
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	var stack []*TreeNode
	stack = append(stack, root)
	var inorderIndex int
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}
