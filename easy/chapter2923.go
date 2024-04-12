package easy

func findChampion(grid [][]int) int {

	for i := 0; i < len(grid); i++ {
		or:=0
		for j:=0;j<len(grid);j++ {
			or|=grid[j][i]
		}
		if or==0 {
			return i
		}


	}
	return -1

}
