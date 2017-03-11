package bruteforce

import (
	"fmt"
	"strings"

	cc "github.com/mohamedkhattab/ceasar-cipher/ceasarcipher"
)

type BruteForceMulti struct{}

func (bfs *BruteForceMulti) Crack(cipherText string) (map[int]int32, []string) {
	keys := make(map[int]int32)
	text := make([]string, 10)

	cipher := cc.MultiCeasar{[][2]int32{
		[2]int32{0, 0},
	}}

	words := strings.Split(cipherText, " ")
	var word string

	for i := 0; i < len(words); i++ {
	Loop:
		for j := 0; j < 54; j++ {

			cipher.Key = [][2]int32{
				[2]int32{0, int32(j)},
			}

			word = cipher.Decrypt(words[i])

			for k := 0; k < len(store); k++ {
				if word == store[k] {
					fmt.Println("word: ", word, "key: ", j)
					text = append(text, word)
					keys[i] = int32(j)
					break Loop
				}
			}
		}
	}

	return keys, text
}
