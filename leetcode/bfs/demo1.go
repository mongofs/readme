package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
	1.层序遍历有限思考广度有限遍历可借助queue
*/
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	temq := [] *TreeNode{root}
	for i:=0;len(temq)> 0;i++ {
		nextQ := []*TreeNode{}
		res = append(res, []int{})
		for j:= 0;j< len(temq);j++{
			node := temq[j]
			res[i] = append(res[i], node.Val)
			if node.Left !=nil {
				nextQ = append(nextQ,node.Left )
			}
			if node.Right !=nil {
				nextQ = append(nextQ,node.Right )
			}
		}
		temq =nextQ
	}
	return res
}

/*
	使用两个queue 进行返回
*/