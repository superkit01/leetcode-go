package main

import("sort")

func getKthMagicNumber(k int) int {

    cache:=make([]int,0)
	cache=append(cache,1)
	unique:=make(map[int]bool,0)
	unique[1]=true


	for i:=1;i<k;i++{

		temp:=cache[0]
		cache=cache[1:]
		if _,ok:=unique[temp*3];!ok{
            unique[temp*3]=true
			cache=append(cache,temp*3)
		}
		if _,ok:=unique[temp*5];!ok{
              unique[temp*5]=true
			cache=append(cache,temp*5)
		}
		if _,ok:=unique[temp*7];!ok{
              unique[temp*7]=true
			cache=append(cache,temp*7)
		}
		sort.Ints(cache)
	}

	return cache[0]
	
}