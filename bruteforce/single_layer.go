package bruteforce

import (
	"strings"

	cc "github.com/mohamedkhattab/ceasar-cipher/ceasarcipher"
)

type BruteForceSingle struct{}

func (bfs *BruteForceSingle) Crack(cipherText string) (int, string) {
	// init cipher
	// loop for 94 iterations
	// decrypt each iteration wiht new key
	var key int
	var text string
	cipher := cc.CeasarCipher{0}

Loop:
	for i := 1; i <= 94; i++ {
		cipher.Key = int32(i)
		words := cipher.Decrypt(cipherText)
		word := strings.Split(words, " ")[0]
		for j := 0; j < len(store); j++ {
			if word != "" && word == store[j] {
				text = words
				key = i
				break Loop
			}
		}
	}

	return key, text
}
