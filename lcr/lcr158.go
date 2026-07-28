package lcr

func inventoryManagementII(stock []int) int {
	vote, cnt := 0, 0
	for _, v := range stock {
		if cnt == 0 {
			vote = v
		}
		if v == vote {
			cnt++
		} else {
			cnt--
		}
	}
	return vote

}
