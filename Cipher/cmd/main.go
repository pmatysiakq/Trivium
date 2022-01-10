package main

import (
	"flag"
	"fmt"
	"github.com/pmatysiakq/Trivium/Cipher/TriviumGo"
)

func main() {
	var key, iv, message string
	var encrypt, decrypt bool

	flag.StringVar(&key, "k", "", "A key to encrypt/decrypt message (80 bit, hex)")
	flag.StringVar(&iv, "i", "", "An iv to encrypt/decrypt message (80 bit, hex)")
	flag.BoolVar(&encrypt, "e", false, "Encrypt message")
	flag.BoolVar(&decrypt, "d", false, "Decrypt encrypted message in the fly. Disallowed without '-e' flag")
	flag.StringVar(&message, "m", "", "Message to be encrypted/decrypted (Plaintext not encoded)")
	flag.String("c", "", "Do not use. Not implemented!")

	flag.Parse()

	var ciphertext []uint8
	var cipherHex string

	if encrypt {
		fmt.Println("---------- ENCRYPTION ----------")
		triviumEncrypt := TriviumGo.NewTrivium(key, iv)
		ciphertext, cipherHex = triviumEncrypt.Encrypt(message)
		fmt.Println("Ciphertext:", cipherHex)
		fmt.Println("--------------------------------")
	} else if decrypt {
		fmt.Println("----------- ERROR -----------")
		fmt.Println("Decrypt not implemented!")
		fmt.Println("You can decrypt on fly currently encrypted message")
		fmt.Println("-----------------------------")
	}

	if encrypt && decrypt {
		fmt.Println("---------- DECRYPTION ----------")
		triviumDecrypt := TriviumGo.NewTrivium(key, iv)
		plaintext := triviumDecrypt.Decrypt(ciphertext)
		fmt.Println("Plaintext:", plaintext)
		fmt.Println("--------------------------------")
	}
}