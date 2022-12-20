package main


type TreeNode struct {
	Val  int
	Left *TreeNode
	Right *TreeNode

}


func maxDepth( root *TreeNode)int{
	if root==nil {
		return 0
	}
	result:=0
	queue:=make([]*TreeNode,0)
	queue=append(queue,root)

	for len(queue)>0 {
		result++

		tempQueue:=make([]*TreeNode,0)

		for _,v:=range queue{

			if v.Left!=nil{
				tempQueue=append(tempQueue,v.Left)
			}

			if v.Right!=nil{
				tempQueue=append(tempQueue,v.Right)
			}

		}
		queue=tempQueue
	}
	return result

}