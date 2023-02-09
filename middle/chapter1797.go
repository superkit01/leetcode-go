package middle

type AuthenticationManager struct {
	timeToLive int
	cache      map[string]int
}

func Constructor2(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		timeToLive: timeToLive,
		cache:      make(map[string]int, 0),
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.cache[tokenId] = currentTime

}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if v, ok := this.cache[tokenId]; ok {
		if v+this.timeToLive > currentTime {
			this.cache[tokenId] = currentTime
		}
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	max := currentTime - this.timeToLive
	result := 0
	for _, v := range this.cache {
		if v > max {
			result++
		}
	}
	return result

}
