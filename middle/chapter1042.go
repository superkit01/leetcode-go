package middle

func gardenNoAdj(n int, paths [][]int) []int{ 
    nearTable:=make(map[int][]int,0)
    

    for _,v:= range paths{
        if _,ok:=nearTable[v[0]];!ok{
            nearTable[v[0]]=make([]int,0)
        }
        nearTable[v[0]]=append(nearTable[v[0]],v[1])
        
        if _,ok:=nearTable[v[1]];!ok{
            nearTable[v[1]]=make([]int,0)
        }
        nearTable[v[1]]=append(nearTable[v[1]],v[0])
    }

    ans:=make([]int,n)
    
    outer:
    for i:=0;i<n;i++{
        colors:=[]int{1,2,3,4}
        colorsBool:=[]bool{false,false,false,false}
        for _,v:= range nearTable[i+1] {
            if ans[v-1] !=0 {
              colorsBool[ans[v-1]-1]=true
            }
        }
        for k:=range colors{
            if !colorsBool[k] {
               ans[i]=colors[k]
                continue outer
            }
        }
        
    }
    
    return ans

}