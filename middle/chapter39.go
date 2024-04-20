package middle


func combinationSum(candidates []int, target int) [][]int {
    
    ans:=make([][]int,0)
    recursive(candidates,0,target,[]int{},&ans)
    return ans

}

//dfs  i位置 使用或不使用
func recursive(candidates []int,index int,remaining int,current []int,ans *[][]int){
    if remaining <0 {
        return 
    }
    
    if remaining == 0{ 
        temp3:=make([]int,0)
        temp3=append(temp3,current...)
        
        *ans=append(*ans,temp3)
        return
    }
    if index >= len(candidates){
        return
    }
    
    temp:=make([]int,0)
    temp=append(temp,current...)
    temp=append(temp,candidates[index])
    
    recursive(candidates, index,remaining-candidates[index],temp,ans)
    
    temp2:=make([]int,0)
    temp2=append(temp2,current...)
    recursive(candidates, index+1,remaining,temp2,ans)
   
    
    
}