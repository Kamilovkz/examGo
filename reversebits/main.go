package main

func ReverseBits(oct byte) byte {
	var res byte
	for i := 0; i < 8; i++ {
		if oct&(1<<i) != 0 {
			res |= 1 << (7 - i)
		}
	}
	return res
}
