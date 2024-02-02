package middle

import (
	"sort"
)

/**
* 贪心
* alice取i下标  bob取j下标
* alice[i]-bob[j]
*
* alice取j下标  bob取i下标
* alice[j]-bob[i]
*
* 哪一种大alice取哪一种
*
*
* 转化后alice[i]+bob[i]   alice[j]+bob[j] 取较大值
*
 */

type AbsIndex struct {
	Val   int
	Index int
}

func stoneGameVI(aliceValues []int, bobValues []int) int {

	absIndex := make([]*AbsIndex, len(aliceValues))
	for i := 0; i < len(aliceValues); i++ {
		absIndex[i] = &AbsIndex{
			Val:   aliceValues[i] + bobValues[i],
			Index: i,
		}
	}

	sort.Slice(absIndex, func(i, j int) bool {
		return absIndex[i].Val-absIndex[j].Val > 0
	})
	alice := 0
	bob := 0

	for i := 0; i < len(aliceValues); i++ {
		if i%2 == 0 {
			alice += aliceValues[absIndex[i].Index]
		} else {
			bob += bobValues[absIndex[i].Index]
		}
	}

	if alice == bob {
		return 0
	} else if alice > bob {
		return 1
	} else {
		return -1
	}

}
