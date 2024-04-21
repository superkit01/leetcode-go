package lcr

func combine(n int, k int) [][]int {
    ans:= make([][]int,0)
    
    dfs77(n,k,1,make([]int,0),&ans)
    
    return ans

}


func dfs77(n int, k int,index int, current []int, ans *[][]int){
    if len(current)>k {
        return
    }
    
    if len(current)==k {
        temp:=make([]int,0)
        temp=append(temp,current...)
        *ans=append(*ans,temp)
        return
    }
    
    if index>n {
        return
    }
    
    
    //选择当前位置
    temp2:=make([]int,0)
    temp2=append(temp2,current...)
    temp2=append(temp2,index)
    dfs77(n,k,index+1, temp2,ans)
    
    //不选择当前位置
    temp3:=make([]int,0)
    temp3=append(temp3,current...)
    dfs77(n,k,index+1, temp3,ans)
    
}