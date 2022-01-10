// https://www.ecrypt.eu.org/stream/p3ciphers/trivium/trivium_p3.pdf
package TriviumGo

import (
	"fmt"
	"os"
)

// Trivium represents Trivium Cipher
type Trivium struct {
	State [288]uint8
	Key string
	Iv string
}

// NewTrivium creates new cipher with `key` and `iv` arguments
// Returns new Trivium object
func NewTrivium(key, iv string) *Trivium {
	return &Trivium{
		Iv: iv,
		Key: key,
	}
}

// Encrypt is used to encrypt given message
// Returns encrypted message in []int8 format compatible with decrypt
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

// Decrypt is used to retrieve encrypted message
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

	var KEY []uint8 = HexToBin(t.Key)
	for i := 0; i < len(KEY); i++ {
		initState = append(initState, KEY[i])
	}

	for i := 0; i < 13; i++ {
		initState = append(initState, uint8(0))
	}

	var IV []uint8 = HexToBin(t.Iv)
	for i := 0; i < len(IV); i++ {
		initState = append(initState, IV[i])
	}

	for i := 0; i < 112; i++ {
		initState = append(initState, uint8(0))
	}
	initState = append(initState, []uint8{1, 1, 1}...)

	if len(initState) == 288 {
		//fmt.Printf("LOG::State assembly sucess! | KeyStream length: %v\n", len(initState))
	} else {
		fmt.Printf("LOG::State assembly failed! | KeyStream length: %v\n", len(initState))
	}

	for index, value := range initState {
		t.State[index] = value
	}

	for i := 0; i < 4*288; i++ {
		t.KeyStreamGenerator()
	}
}
