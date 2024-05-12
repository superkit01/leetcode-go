package week397

import "math"

func findPermutationDifference(s string, t string) int {
    posMap:=make(map[rune]int,0)
    for i,v:=range s{
        posMap[v]=i
    }
    ans:=0
    for i,v:=range t{
        ans+=int(math.Abs(float64(i-posMap[v] )))    
    }
    return ans
    
}