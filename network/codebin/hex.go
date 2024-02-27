package codebin

import "encoding/hex"

var hexNumbers = []byte("0123456789ABCDEF")

// ByteToHex 编码
func ByteToHex(value byte) []byte {
	buf := make([]byte, 2)
	buf[0] = hexNumbers[value>>4]
	buf[1] = hexNumbers[value&0x0F]
	return buf
}

// WriteByteHex 编码
func WriteByteHex(buf []byte, value uint8) {
	buf[0] = hexNumbers[value>>4]
	buf[1] = hexNumbers[value&0x0F]
}

// WriteUint8Hex 编码
func WriteUint8Hex(buf []byte, value uint8) {
	buf[0] = hexNumbers[value>>4]
	buf[1] = hexNumbers[value&0x0F]
}

// WriteUint16Hex 编码
func WriteUint16Hex(buf []byte, value uint16) {
	h, l := value>>8, value&0xF
	buf[0] = hexNumbers[h>>4]
	buf[1] = hexNumbers[h&0x0F]
	buf[3] = hexNumbers[l>>4]
	buf[4] = hexNumbers[l&0x0F]

}

// ToHex 编码
func ToHex(values []byte) []byte {
	length := len(values)
	buf := make([]byte, length<<1) //length * 2
	for i := 0; i < length; i++ {
		value := values[i]
		j := i << 1 //i * 2
		buf[j] = hexNumbers[value>>4]
		buf[j+1] = hexNumbers[value&0x0F]
	}
	return buf
}

// FromHex 解码
func FromHex(values []byte) []byte {
	buf := make([]byte, len(values)>>1) //length / 2
	_, _ = hex.Decode(buf, values)
	return buf
}
