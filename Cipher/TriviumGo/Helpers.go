package TriviumGo

import (
	"strconv"
)

func HexToBin(hexTxt string) (binTxt []uint8) {
	for i := 0; i < len(hexTxt); i++ {
		value, _ := strconv.ParseInt(string(hexTxt[i]), 16, 8)
		binTxt = append(binTxt, (uint8(value)&0b1000)>>3)
		value, _ = strconv.ParseInt(string(hexTxt[i]), 16, 8)
		binTxt = append(binTxt, (uint8(value)&0b0100)>>2)
		value, _ = strconv.ParseInt(string(hexTxt[i]), 16, 8)
		binTxt = append(binTxt, (uint8(value)&0b0010)>>1)
		value, _ = strconv.ParseInt(string(hexTxt[i]), 16, 8)
		binTxt = append(binTxt, uint8(value)&0b0001)
	}
	return
}

func BinToHex(val []uint8) (hexOutput string) {
	if len(val)%4 != 0 {
		panic(val)
	}

	for i := 0; i < len(val); i += 4 {
		decimal := val[i]*8 + val[i+1]*4 + val[i+2]*2 + val[i+3]*1
		hexOutput += strconv.FormatUint(uint64(decimal), 16)
	}

	return hexOutput
}

//func HexToStr(hexTxt string) (strTxt string) {
//	byteArrTxt, _ := hex.DecodeString(hexTxt)
//	strTxt = string(byteArrTxt)
//
//	return
//}

//func StringToBin(s string) (binArray []uint8) {
//	var binString string
//	for _, c := range s {
//		byteValue := fmt.Sprintf("%b", c)
//		if len(byteValue) < 8 {
//			count := 8 - len(byteValue)
//			byteValue = strings.Repeat("0", count) + byteValue
//		}
//		binString = fmt.Sprintf("%s%s", binString, byteValue)
//	}
//	for i := 0; i < len(binString); i++ {
//		temp, err := strconv.Atoi(string(binString[i]))
//		if err != nil {
//			fmt.Printf("error::%s\n", err)
//		}
//		binArray = append(binArray, uint8(temp))
//	}
//	return
//}

//func BinToString(b []uint8) (plaintext string) {
//	var text string
//
//	for i := 0; i < len(b); i++ {
//		text += strconv.Itoa(int(b[i]))
//	}
//
//	plaintext = string(BitStrToByteSlice(text))
//
//	return
//}

//func BitStrToByteSlice(b string) []byte {
//	var out []byte
//	var str string
//
//	for i := len(b); i > 0; i -= 8 {
//		if i-8 < 0 {
//			str = string(b[0:i])
//		} else {
//			str = string(b[i-8 : i])
//		}
//		v, err := strconv.ParseUint(str, 2, 8)
//		if err != nil {
//			panic(err)
//		}
//		out = append([]byte{byte(v)}, out...)
//	}
//	return out
//}
