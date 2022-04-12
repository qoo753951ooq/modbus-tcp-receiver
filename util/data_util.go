package util

import (
	"bytes"
	"encoding/binary"

	"github.com/shopspring/decimal"
)

func CombineString(values ...string) string {
	var buffer bytes.Buffer

	for _, v := range values {
		buffer.WriteString(v)
	}

	return buffer.String()
}

func SubString(str string, begin, lenght int) (substr string) {

	//fmt.Println("Substring =", str)
	rs := []rune(str)
	lth := len(rs)
	//fmt.Printf("begin=%d, end=%d, lth=%d\n", begin, lenght, lth)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + lenght

	if end > lth {
		end = lth
	}
	//fmt.Printf("begin=%d, end=%d, lth=%d\n", begin, lenght, lth)
	return string(rs[begin:end])
}

func GetDataTimeSecondToZero(datatime string) string {

	var buffer bytes.Buffer

	nowRemoveSec := SubString(datatime, 0, len(datatime)-2)
	buffer.WriteString(nowRemoveSec)
	buffer.WriteString("00")

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

func BitIndex2ArrayIndex(bitIndex int) int {

	var arrayIndex int

	arrayIndex = 15 - bitIndex

	if arrayIndex < 0 {
		return -1
	}

	return arrayIndex
}

func MultiplyDP(value uint16, decimal_places string) float64 {

	var result float64

	mulValue := decimal.NewFromInt32(int32(value))

	switch decimal_places {
	case "0":
		result, _ = mulValue.Mul(decimal.NewFromInt(1)).Float64()
	case "1":
		result, _ = mulValue.Mul(decimal.NewFromFloat(0.1)).Float64()
	case "2":
		result, _ = mulValue.Mul(decimal.NewFromFloat(0.01)).Float64()
	case "3":
		result, _ = mulValue.Mul(decimal.NewFromFloat(0.001)).Float64()
	case "4":
		result, _ = mulValue.Mul(decimal.NewFromFloat(0.0001)).Float64()
	}

	return result
}
