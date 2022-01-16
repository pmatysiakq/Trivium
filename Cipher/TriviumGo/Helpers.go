package TriviumGo

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
)

// HexToBin converts  hexadecimal representation - string to binary slice - []byte
func HexToBin(hexTxt string) (binTxt []uint8) {
	for i := 0; i < len(hexTxt); i += 2 {

		value, _ := strconv.ParseInt(string(hexTxt[i+1]), 16, 8)
		binTxt = append(binTxt, (uint8(value)&0b0001))
		value, _ = strconv.ParseInt(string(hexTxt[i+1]), 16, 8)
		binTxt = append(binTxt, (uint8(value)&0b0010)>>1)
		value, _ = strconv.ParseInt(string(hexTxt[i+1]), 16, 8)
		binTxt = append(binTxt, (uint8(value)&0b0100)>>2)
		value, _ = strconv.ParseInt(string(hexTxt[i+1]), 16, 8)
		binTxt = append(binTxt, (uint8(value)&0b1000)>>3)

		value, _ = strconv.ParseInt(string(hexTxt[i]), 16, 8)
		binTxt = append(binTxt, (uint8(value)&0b0001))
		value, _ = strconv.ParseInt(string(hexTxt[i]), 16, 8)
		binTxt = append(binTxt, (uint8(value)&0b0010)>>1)
		value, _ = strconv.ParseInt(string(hexTxt[i]), 16, 8)
		binTxt = append(binTxt, (uint8(value)&0b0100)>>2)
		value, _ = strconv.ParseInt(string(hexTxt[i]), 16, 8)
		binTxt = append(binTxt, (uint8(value)&0b1000)>>3)
	}
	return
}

// BinToHex converts binary slice - []byte to hexadecimal representation - string
func BinToHex(val []uint8) (hexOutput string) {
	if len(val)%4 != 0 {
		panic(val)
	}

	for i := 0; i < len(val); i += 8 {
		decimal := val[i+4]*1 + val[i+5]*2 + val[i+6]*4 + val[i+7]*8
		hexOutput += strings.ToUpper(strconv.FormatUint(uint64(decimal), 16))

		decimal = val[i]*1 + val[i+1]*2 + val[i+2]*4 + val[i+3]*8
		hexOutput += strings.ToUpper(strconv.FormatUint(uint64(decimal), 16))
	}

	return hexOutput
}

// GenerateXZeroMsg is implemented for testing purposes only
func GenerateXZeroMsg(X int) (msg string){
	for i :=0; i < X; i++ {
		msg += "0"
	}
	return
}

// RandStringRunes generates hexadecimal string containing n characters
// Avaliable characters: [0-9] and [A-F]
func RandStringRunes(n int) string {
	var letterRunes = []rune("0123456789ABCDEF")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// PrintMemUsage outputs the total memory being used.
// Returns this total memory parameter
func PrintMemUsage(msg string) uint64{
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("\t%s = %v MiB\n", msg, bToMb(m.TotalAlloc))

	return m.TotalAlloc

}

// bToMb converts bytes to MegaBytes
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}