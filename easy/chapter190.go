package main

func reverseBits(num uint32) uint32 {
	var result uint32 = 0
	for i := 0; i < 32; i++ {
		bit := num % 2
		num = num >> 1

		result = result<<1 | bit
	}
	return result
}
