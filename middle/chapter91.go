package middle

func numDecodings(s string) int {

    if s[0]=='0' {
        return 0
    }
    
    
    // 前i个字母可解码数
    dp:=make([]int,len(s))
    dp[0]=1
    for i:=1; i<len(s); i++{
        if s[i] !='0' {
            dp[i]=dp[i-1]
        }
        if s[i-1]=='1' || (s[i-1]=='2' && s[i]<'7'){
            if i==1 {
                dp[i]+=1
            }else{
                dp[i]+=dp[i-2]
            }
        }
        
        
    }
    return dp[len(dp)-1]
}