package TriviumGo

import (
	"strconv"
	"strings"
)

// HexToBin converts  hexadecimal representation - string to binary slice - []byte
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

// BinToHex converts binary slice - []byte to hexadecimal representation - string
func BinToHex(val []uint8) (hexOutput string) {
	if len(val)%4 != 0 {
		panic(val)
	}

	for i := 0; i < len(val); i += 4 {
		decimal := val[i]*8 + val[i+1]*4 + val[i+2]*2 + val[i+3]*1
		hexOutput += strings.ToUpper(strconv.FormatUint(uint64(decimal), 16))
	}

	return hexOutput
}