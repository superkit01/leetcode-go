package middle

func MaxOperations(nums []int) int {
	if len(nums)==2 {
        return 1
    }

    cache:=make(map[[3]int]int,0)

    var dfs  func(i,j,score int)int
    dfs=func(i,j,score int)int{
        if v,ok:=cache[[3]int{i,j,score}];ok{
            return v
        }
        if i>=j{
            cache[[3]int{i,j,score}]=0
            return 0
        }
        ans:=0
        if nums[i]+nums[j]==score {
            ans=max(dfs(i+1,j-1,score)+1,ans)
        }
        if nums[i]+nums[i+1]==score {
            ans=max(dfs(i+2,j,score)+1,ans)
        }
        if nums[j]+nums[j-1]==score {
            ans=max(dfs(i,j-2,score)+1,ans)
        }

        cache[[3]int{i,j,score}]=ans
        return ans
    }

   return max(max(dfs(0,len(nums)-1,nums[0]+nums[1]), dfs(0,len(nums)-1,nums[0]+nums[len(nums)-1])),dfs(0,len(nums)-1,nums[len(nums)-2]+nums[len(nums)-1]) )
        

}
