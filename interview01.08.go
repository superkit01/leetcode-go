package main


func setZeroes2(matrix [][]int)  {
	row:=make(map[int]bool)
	col:=make(map[int]bool)

	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[i]);j++{
			if matrix[i][j]==0{
				row[i]=true
				col[j]=true
			}
		}
	}

	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[i]);j++{
			if v,ok:=row[i];ok{
				if v{
					matrix[i][j]=0
				}
			}
			if v,ok:=col[j];ok{
				if v{
					matrix[i][j]=0
				}
			}
		}
	}

}