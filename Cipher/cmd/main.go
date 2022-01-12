package main

import (
	"flag"
	"fmt"
	"github.com/pmatysiakq/Trivium/Cipher/TriviumGo"
	"strings"
)

func main() {
	var key, iv, msg, cipher, mode string


	flag.StringVar(&key, "key", "", "A KEY to encrypt/decrypt message (80 bit, HEX)")
	flag.StringVar(&iv, "iv", "", "An IV to encrypt/decrypt message (80 bit, HEX)")
	flag.StringVar(&mode, "mode", "null", "e - encrypt, d - decrypt output HEX," +
		" dp - decrypt - output HEX and PLAINTEXT")
	flag.StringVar(&msg, "msg", "", "Message to be encrypted/decrypted (HEX))")
	flag.StringVar(&cipher,"cipher", "", "Ciphertext to decrypt (HEX)")

	flag.Parse()

	if strings.ToLower(mode) == "e" {
		fmt.Println("---------- ENCRYPTION ----------")
		triviumEncrypt := TriviumGo.NewTrivium(key, iv)
		cipherHex := triviumEncrypt.Encrypt(msg)
		fmt.Println("Ciphertext:", cipherHex)
		fmt.Println("--------------------------------")
	}

	if strings.ToLower(mode) == "d" ||  strings.ToLower(mode) == "dp" {
		fmt.Println("---------- DECRYPTION ----------")
		triviumDecrypt := TriviumGo.NewTrivium(key, iv)
		plaintext, plainHex := triviumDecrypt.Decrypt(cipher)
		fmt.Println("Plaintext (hex):", plainHex)

		if strings.ToLower(mode) == "dp" {
			fmt.Println("Plaintext:", plaintext)
		}
		fmt.Println("--------------------------------")
	}

}
