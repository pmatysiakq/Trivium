package TriviumGo

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func HexToBin(hexTxt string) (binTxt []uint8) {
	binTxt = StringToBin(HexToStr(hexTxt))

	return
}

func HexToStr(hexTxt string) (strTxt string) {
	byteArrTxt, _ := hex.DecodeString(hexTxt)
	strTxt = string(byteArrTxt)

	return
}

func StringToBin(s string) (binArray []uint8) {
	var binString string
	for _, c := range s {
		byteValue := fmt.Sprintf("%b", c)
		if len(byteValue) < 8 {
			count := 8 - len(byteValue)
			byteValue = strings.Repeat("0", count) + byteValue
		}
		binString = fmt.Sprintf("%s%s", binString, byteValue)
	}
	for i := 0; i < len(binString); i++ {
		temp, err := strconv.Atoi(string(binString[i]))
		if err != nil {
			fmt.Printf("error::%s\n", err)
		}
		binArray = append(binArray, uint8(temp))
	}
	return
}

func BinToString(b []uint8) (plaintext string) {
	var text string

	for i := 0; i < len(b); i++ {
		text += strconv.Itoa(int(b[i]))
	}

	plaintext = string(BitStrToByteSlice(text))

	return
}

func BitStrToByteSlice(b string) []byte {
	var out []byte
	var str string

	for i := len(b); i > 0; i -= 8 {
		if i-8 < 0 {
			str = string(b[0:i])
		} else {
			str = string(b[i-8 : i])
		}
		v, err := strconv.ParseUint(str, 2, 8)
		if err != nil {
			panic(err)
		}
		out = append([]byte{byte(v)}, out...)
	}
	return out
}