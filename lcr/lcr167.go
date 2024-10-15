package lcr

//某套连招动作记作序列 arr，其中 arr[i] 为第 i 个招式的名字。请返回 arr 中最多可以出连续不重复的多少个招式。
//abba
func dismantlingAction(arr string) int {
	cache := make(map[rune]int, 0) 
	max := 0
	curr := 0
	for i, v := range arr {
		if index, ok := cache[v]; ok{
			if index < i-curr{
				curr++
			}else{
				curr = i - index
			}
		} else {
			curr++
		}
		cache[v] = i
		if max < curr {
			max = curr
		}
	}
	return max

}
