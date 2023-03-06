package middle

import "math"

func minimumDeletions(s string) int {
	countB:=0
	dp:=make([]int,len(s))
	dp[0]=0

	if s[0]=='b'{
		countB++
	}

	for i:=1;i<len(s);i++{
		if s[i]=='b'{
			countB++
			dp[i]=dp[i-1]
			continue
		}

		if s[i]=='a'{
			dp[i]=int(math.Min(float64(dp[i-1]+1),float64(countB) ))
		}
	}

	return dp[len(dp)-1]
}
