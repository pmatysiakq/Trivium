package TriviumGo

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Trivium struct {
	// https://www.ecrypt.eu.org/stream/p3ciphers/trivium/trivium_p3.pdf
	// https://github.com/uisyudha/Trivium
	State [288]uint8
	Key string
	Iv string
}

//func LogState

func NewTrivium(key, iv string) *Trivium {
	return &Trivium{
		Iv: iv,
		Key: key,
	}
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
	var words []string
	for i := 0; i < len(b); i += 8 {
		var word string
		for j :=0; j < 8; j++ {
			word += strconv.Itoa(int(b[i+j]))
		}
		words = append(words, word)
	}
	for i := 0; i < len(words); i++ {
		plaintext += string(bitString(words[i]).AsByteSlice())
	}

	return
}

func (t *Trivium) Encrypt(msg string) (ciphertext []uint8) {

	messageBinArray := StringToBin(msg)

	keyStream := t.GenerateKeyStream(len(messageBinArray))

	if len(messageBinArray) != len(keyStream) {
		fmt.Println("LOG::Message and KeyStream are of different size!")
		os.Exit(0)
	}

	for i := 0; i < len(messageBinArray); i++ {
		ciphertext = append(ciphertext, messageBinArray[i] ^ keyStream[i])
	}
	return
}

func (t *Trivium) Decrypt(ciphertext []uint8) (message string) {
	var msg []uint8
	keyStream := t.GenerateKeyStream(len(ciphertext))

	for i := 0; i < len(ciphertext); i++ {
		msg = append(msg, ciphertext[i] ^ keyStream[i])
	}

	message = BinToString(msg)

	return
}

func (t *Trivium) UpdateState(t1, t2, t3 uint8){
	var newState [288]uint8
	for i := 0; i < len(t.State) - 1; i++ {
		newState[i+1] = t.State[i]
	}

	t.State = newState

	t.State[0] = t3
	t.State[93] = t1
	t.State[177] = t2
}

func (t *Trivium) GenerateKeyStream(msgLen int) (keyStream []uint8) {
	t.Initialize()

	counter := 0

	for counter < msgLen {
		keyStream = append(keyStream, t.KeyStreamGenerator())
		counter++
	}
	return
}

func (t *Trivium) KeyStreamGenerator() (keyStream uint8) {
	t1 := t.State[65] ^ t.State[92]
	t2 := t.State[161] ^ t.State[176]
	t3 := t.State[242] ^ t.State[267]

	keyStream = t1 ^ t2 ^ t3

	t1 = t1 ^ t.State[90] & t.State[91] ^ t.State[170]
	t2 = t1 ^ t.State[174] & t.State[175] ^ t.State[263]
	t3 = t1 ^ t.State[285] & t.State[286] ^ t.State[68]

	t.UpdateState(t1, t2, t3)

	return
}

func (t *Trivium) Initialize() {
	var initState []uint8
	for i := 0; i < len(t.Key); i++ {
		temp, err := strconv.Atoi(string(t.Key[i]))
		if err != nil {
			fmt.Printf("error::%s\n", err)
		}
		initState = append(initState, uint8(temp))
	}

	for i := 0; i < 13; i++ {
		initState = append(initState, uint8(0))
	}

	for i := 0; i < len(t.Iv); i++ {
		temp, err := strconv.Atoi(string(t.Iv[i]))
		if err != nil {
			fmt.Printf("error::%s\n", err)
		}
		initState = append(initState, uint8(temp))
	}

	for i := 0; i < 112; i++ {
		initState = append(initState, uint8(0))
	}
	initState = append(initState, []uint8{1, 1, 1}...)

	if len(initState) == 288 {
		fmt.Printf("LOG::State assembly sucess!\n")
	} else {
		fmt.Printf("LOG::State assembly failed!\n")
	}

	for index, value := range initState {
		t.State[index] = value
	}

	for i := 0; i < 4*288; i++ {
		t.KeyStreamGenerator()
	}
}

type bitString string

func (b bitString) AsByteSlice() []byte {
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
