package middle

func minimumArea(grid [][]int) int {
    x:=[]int{len(grid[0]),-1}
    y:=[]int{len(grid),-1}
    for i,v:=range grid {
        for j,m:=range v {
            if m==1 {
                    if x[0]>j{
                        x[0]=j
                    }
                    if x[1]<j+1 {
                        x[1]=j+1
                    }
                    if y[0]>i{
                        y[0]=i
                    }
                    if y[1]<i+1 {
                        y[1]=i+1
                    }
                    
            }
        }
    }

    return (x[1]-x[0]) * (y[1]-y[0])
    
}
