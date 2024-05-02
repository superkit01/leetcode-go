package middle


func canFinish(numCourses int, prerequisites [][]int) bool {
    //入度
    indegree:=make([]int,numCourses)
    
    //建图
    graph:=make(map[int][]int,0)
    for _,v:=range prerequisites {
        indegree[v[0]]++
        if _,ok:=graph[v[1]];!ok{
            graph[v[1]]=make([]int,0)
        }
        graph[v[1]]=append(graph[v[1]],v[0])   
    }

    //bfs 遍历 入度0
    queue:=make([]int,0)
    for i,v:=range indegree{
        if v==0{
            queue=append(queue,i)
        }
    }
    
    
    for len(queue)>0 {
        temp:=make([]int,0)
        for _,v:=range queue{
            for _,e:=range  graph[v]{
                indegree[e]--
                if indegree[e]==0 {
                    temp=append(temp,e)
                }
            }
        }
        queue=temp
    }
    
    for _,v:=range indegree {
        if v>0 {
            return false
        }
    }
    return true
    
}