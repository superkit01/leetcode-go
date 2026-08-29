package lcr

func encryptionCalculate(dataA int, dataB int) int {

	for dataB != 0 {
		c := dataA & dataB    //进位
		dataA = dataA ^ dataB //本位
		dataB = (c << 1)
	}
	return dataA

}
