package util

import (
	"bytes"
	"encoding/binary"
)

func CombineString(values ...string) string {
	var buffer bytes.Buffer

	for _, v := range values {
		buffer.WriteString(v)
	}

	return buffer.String()
}

func Bytes2Bits(data []byte) []int {
	dst := make([]int, 0)
	for _, v := range data {
		for i := 0; i < 8; i++ {
			move := uint(7 - i)
			dst = append(dst, int((v>>move)&1))
		}
	}
	return dst
}

func Bytes2Decimal(data []byte) []uint16 {
	result := make([]uint16, 0)

	for i := 1; i < len(data); i++ {

		if i%2 == 0 {
			continue
		}

		result = append(result, binary.BigEndian.Uint16(data[i-1:i+1]))
	}
	return result
}
