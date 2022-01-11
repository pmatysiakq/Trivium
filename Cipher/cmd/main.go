package main

import (
	"flag"
	"fmt"
	"github.com/pmatysiakq/Trivium/Cipher/TriviumGo"
)

func main() {
	var key, iv, message, cipher string
	var encrypt, decrypt, printPlain bool

	flag.StringVar(&key, "k", "", "A key to encrypt/decrypt message (80 bit, hex)")
	flag.StringVar(&iv, "i", "", "An iv to encrypt/decrypt message (80 bit, hex)")
	flag.BoolVar(&encrypt, "e", false, "Encrypt message")
	flag.BoolVar(&decrypt, "d", false, "Decrypt encrypted message. Use flag '-c' to provide ciphertext (hex)")
	flag.BoolVar(&printPlain, "p", false, "If -d is provided, then decrypted msg will be printed in plaintext")
	flag.StringVar(&message, "m", "", "Message to be encrypted/decrypted (Plaintext not encoded)")
	flag.StringVar(&cipher,"c", "", "Ciphertext to decrypt in hex format")

	flag.Parse()

	if encrypt {
		fmt.Println("---------- ENCRYPTION ----------")
		triviumEncrypt := TriviumGo.NewTrivium(key, iv)
		cipherHex := triviumEncrypt.Encrypt(message)
		fmt.Println("Ciphertext:", cipherHex)
		fmt.Println("--------------------------------")
	}

	if decrypt {
		fmt.Println("---------- DECRYPTION ----------")
		triviumDecrypt := TriviumGo.NewTrivium(key, iv)
		plaintext, plainHex := triviumDecrypt.Decrypt(cipher)
		fmt.Println("Plaintext (hex):", plainHex)

		if printPlain {
			fmt.Println("Plaintext:", plaintext)
		}
		fmt.Println("--------------------------------")
	}

}
