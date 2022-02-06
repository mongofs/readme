package main

//  dfs 最著名的就是树的遍历: 前中后 序遍历

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return append(append(preorderTraversal(root.Right), root.Val), preorderTraversal(root.Left)...)
}