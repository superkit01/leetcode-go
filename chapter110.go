package main

import "math"

func isBalanced(root *TreeNode) bool {
    if root==nil {
        return true
    }
    
    return math.Abs(float64(calDeep(root.Left)-calDeep(root.Right)))<=float64(1) && isBalanced(root.Left) && isBalanced(root.Right)
    

}


func calDeep(node *TreeNode) int{
    if node == nil {
        return 0
    }
    
    leftDeep:= calDeep(node.Left)
    rightDeep :=calDeep(node.Right)
    if leftDeep > rightDeep {
        return leftDeep+1
    }else {
        return rightDeep+1
    } 
}