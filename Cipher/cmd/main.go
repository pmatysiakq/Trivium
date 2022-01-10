package main

import (
	"fmt"
	"github.com/pmatysiakq/Trivium/Cipher/TriviumGo"
)

func main() {
	var key, iv string

	key = "576557416e744d417850"
	iv = "486954726976756d2121"

	triviumEncrypt := TriviumGo.NewTrivium(key, iv)
	ciphertext := triviumEncrypt.Encrypt("Ala ma kota!")
	fmt.Println("Ciphertext:", ciphertext)

	triviumDecrypt := TriviumGo.NewTrivium(key, iv)
	plaintext := triviumDecrypt.Decrypt(ciphertext)
	fmt.Println("Plaintext:", plaintext)
}

//func main() {
//	var key, iv, message string
//	var encrypt, decrypt bool
//
//	flag.StringVar(&key, "k", "", "A key to encrypt/decrypt message (80 bit - binary)")
//	flag.StringVar(&iv, "i", "", "An iv to encrypt/decrypt message (80 bit - binary")
//	flag.BoolVar(&encrypt, "e", true, "Encrypt message")
//	flag.BoolVar(&decrypt, "d", true, "Decrypt message")
//	flag.StringVar(&message, "m", "", "Message to be encrypted/decrypted")
//
//	// fdeFsSfgeDeSD (== 80 bitów) TODO
//	//key = "10011000110001000100100001101001011000010000101110110100111111110110110101001011"
//	//iv = "11111100100111011010110010010011110010110110001111001100010110110110100101110100"
//
//	if encrypt {
//		triviumEncrypt := TriviumGo.NewTrivium(key, iv)
//		//ciphertext := triviumEncrypt.Encrypt("Ale ma kota, a kot ma ale XD !!!@@#")
//		ciphertext := triviumEncrypt.Encrypt(message)
//		fmt.Println("Ciphertext:", ciphertext)
//	} else if decrypt {
//		triviumDecrypt := TriviumGo.NewTrivium(key, iv)
//		plaintext := triviumDecrypt.Decrypt([]uint8(message))
//		fmt.Println("Plaintext:", plaintext)
//	}