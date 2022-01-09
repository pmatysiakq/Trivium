package main

import (
	"fmt"
	"github.com/pmatysiakq/Trivium/Cipher/TriviumGo"
)

func main() {
	var key, iv string

	// fdeFsSfgeDeSD (== 80 bitów) TODO
	key = "10011000110001000100100001101001011000010000101110110100111111110110110101001011"
	iv = "11111100100111011010110010010011110010110110001111001100010110110110100101110100"

	triviumEncrypt := TriviumGo.NewTrivium(key, iv)
	ciphertext := triviumEncrypt.Encrypt("Ale ma kota, a kot ma ale XD !!!@@#")
	fmt.Println("Ciphertext:", ciphertext)

	triviumDecrypt := TriviumGo.NewTrivium(key, iv)
	plaintext := triviumDecrypt.Decrypt(ciphertext)
	fmt.Println("Plaintext:", plaintext)
}
