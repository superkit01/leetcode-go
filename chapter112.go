package main


func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    
   return recusive(root, targetSum)
    

}

func recusive(node *TreeNode,remaining int) bool{
    if  node.Left == nil && node.Right == nil {
        if node.Val == remaining{
            return true
        }else{
            return false
        }
    }
    
    if node.Left != nil  && node.Right == nil{
        return recusive(node.Left,remaining - node.Val)
    }
    if node.Right != nil  && node.Left == nil{
        return recusive(node.Right,remaining - node.Val)
    } else {
       return  recusive(node.Left,remaining - node.Val)  || recusive(node.Right,remaining - node.Val)
    }

    
}

