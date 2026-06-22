package easy


//balloon
func maxNumberOfBalloons(text string) int {
    cache := make(map[rune]int,0)
    for _,v:=range text {
        cache[v]++
    }

    return min(cache['b'], cache['a'], cache['n'], cache['l']/2, cache['o']/2)
}

