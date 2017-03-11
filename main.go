package main

import (
	"fmt"

	bf "github.com/mohamedkhattab/ceasar-cipher/bruteforce"
	cc "github.com/mohamedkhattab/ceasar-cipher/ceasarcipher"
)

func main() {
	cipher := cc.MultiCeasar{[][2]int32{
		[2]int32{0, 4},
	}}
	enc := cipher.Encrypt("school sucks")
	fmt.Println(enc)

	cracker := bf.BruteForceMulti{}
	key, text := cracker.Crack(enc)
	fmt.Printf("Key: %d, message: %s", key, text)
}
