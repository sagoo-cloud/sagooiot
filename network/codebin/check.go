package codebin

// Sum 和
func Sum(buf []byte) byte {
	var sum byte = 0
	l := len(buf)
	for i := 0; i < l; i++ {
		sum += buf[i]
	}
	return sum
}

// Xor 异或
func Xor(buf []byte) byte {
	var xor = buf[0]
	l := len(buf)
	for i := 1; i < l; i++ {
		xor ^= buf[i]
	}
	return xor
}
