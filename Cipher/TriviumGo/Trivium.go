// https://www.ecrypt.eu.org/stream/p3ciphers/trivium/trivium_p3.pdf
package TriviumGo

import (
	"encoding/hex"
	"fmt"
	"os"
)

// Trivium represents Trivium Cipher
type Trivium struct {
	State [288]uint8
	Key   string
	Iv    string
}

// NewTrivium creates new cipher with `key` and `iv` arguments
// Returns new Trivium object
func NewTrivium(key, iv string) *Trivium {
	return &Trivium{
		Iv:  iv,
		Key: key,
	}
}

// Encrypt is used to encode message, given as Hex string
// Returns encrypted message in hex format
func (t *Trivium) Encrypt(msg string) (cipherHex string) {

	if len(msg) % 2 != 0 {
		fmt.Printf("Please provide a valid HEX. Can't encode: %s", msg)
		os.Exit(2137)
	}
	messageBinArray := HexToBin(msg)

	keyStream := t.GenerateKeyStream(len(messageBinArray))

	if len(messageBinArray) != len(keyStream) {
		fmt.Println("LOG::Message and KeyStream are of different size!")
		os.Exit(2137)
	}
	var cipherBinArr []uint8
	for i := 0; i < len(messageBinArray); i++ {
		cipherBinArr = append(cipherBinArr, messageBinArray[i]^keyStream[i])
	}

	cipherHex = BinToHex(cipherBinArr)
	return
}

// Decrypt is used to decode encrypted message, given as Hex string
// Returns decoded plaintext
func (t *Trivium) Decrypt(cipherHex string) (plaintext, plainHex string) {

	if len(cipherHex) % 2 != 0 {
		fmt.Printf("Please provide a valid HEX. Can't decode: %s", cipherHex)
		os.Exit(2137)
	}
	messageBinArray := HexToBin(cipherHex)

	keyStream := t.GenerateKeyStream(len(messageBinArray))

	if len(messageBinArray) != len(keyStream) {
		fmt.Println("LOG::Message and KeyStream are of different size!")
		os.Exit(2137)
	}
	var cipherBinArr []uint8
	for i := 0; i < len(messageBinArray); i++ {
		cipherBinArr = append(cipherBinArr, messageBinArray[i]^keyStream[i])
	}

	plainHex = BinToHex(cipherBinArr)
	temp, err := hex.DecodeString(plainHex)
	if err != nil {
		os.Exit(2137)
	}
	plaintext = fmt.Sprintf("%s", temp)
	return
}

// UpdateState shifts values one step to the right and in the internal state
// and then replaces selected bits with values given as parameters
func (t *Trivium) UpdateState(t1, t2, t3 uint8) {
	var newState [288]uint8
	for i := 0; i < len(t.State)-1; i++ {
		newState[i+1] = t.State[i]
	}

	t.State = newState

	t.State[0] = t3
	t.State[93] = t1
	t.State[177] = t2
}

// GenerateKeyStream generates the KeyStream of the same length as length
// of the message which will be encrypted/decrypted
func (t *Trivium) GenerateKeyStream(msgLen int) (keyStream []uint8) {

	t.Initialize()

	counter := 0
	for counter < msgLen {
		keyStream = append(keyStream, t.KeyStreamGenerator())
		counter++
	}
	return
}

// KeyStreamGenerator is used to generate one bit of KeyStream which is returned
// Also internal State is updated
func (t *Trivium) KeyStreamGenerator() (keyStream uint8) {
	t1 := t.State[65] ^ t.State[92]
	t2 := t.State[161] ^ t.State[176]
	t3 := t.State[242] ^ t.State[287]

	keyStream = t1 ^ t2 ^ t3

	t1 = t1 ^ t.State[90] & t.State[91] ^ t.State[170]
	t2 = t2 ^ t.State[174] & t.State[175] ^ t.State[263]
	t3 = t3 ^ t.State[285] & t.State[286] ^ t.State[68]

	t.UpdateState(t1, t2, t3)

	return
}

// Initialize is used to generate initState based on Key nad Iv
// Then the internal state is Updated/Rotated over 4 full cycles (4*288)
func (t *Trivium) Initialize() {
	var initState []uint8

	var KEY, IV []uint8
	tmpKey := HexToBin(t.Key)
	tmpIv := HexToBin(t.Iv)

	for i:= 79; i >= 0; i-- {
		KEY = append(KEY, tmpKey[i])
		IV = append(IV, tmpIv[i])
	}

	if len(KEY) != 80 {
		fmt.Printf("Provide the Key of lenght 80 bits not %v bits\nQuitting!\n", len(KEY))
		os.Exit(2137)
	}

	for i := 0; i < len(KEY); i++ {
		initState = append(initState, KEY[i])
	}

	for i := 0; i < 13; i++ {
		initState = append(initState, uint8(0))
	}

	if len(IV) != 80 {
		fmt.Printf("Provide the IV of lenght 80 bits not %v bits\nQuitting!\n", len(IV))
		os.Exit(2137)
	}

	for i := 0; i < len(IV); i++ {
		initState = append(initState, IV[i])
	}

	for i := 0; i < 112; i++ {
		initState = append(initState, uint8(0))
	}
	initState = append(initState, []uint8{1, 1, 1}...)

	if len(initState) != 288 {
		fmt.Printf("LOG::State assembly failed! | KeyStream length: %v\n", len(initState))
		os.Exit(2137)
	}

	for index, value := range initState {
		t.State[index] = value
	}

	for i := 0; i < 4*288; i++ {
		t.KeyStreamGenerator()
	}
}
