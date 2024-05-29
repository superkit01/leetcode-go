package lcr 

func trainingPlan(actions []int) int {
    bit:=[32]int{}
    for _,v:=range actions{
        i:=0
        for v>0 {
            bit[i] += v&0b1
            v=v>>1
            i++
        }
    }

    ans:=0
    for i,v:=range bit {
        if v%3 !=0 {
            ans|= 1<<i
        }
    }
    return ans
}