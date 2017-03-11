package ceasarcipher

//CeasarCipher stores the key and plain text intended
type CeasarCipher struct {
	Key int32
}

func (c *CeasarCipher) encryptChar(char int32) int32 {
	key := c.Key
	newChar := char
	if c.Key > 94 {
		key = c.Key % 94
	}

	if char < 32 || char > 126 {
		newChar = char
	} else {
		if key+char > 126 {
			key = (key + char) - 126
			newChar = 31 + key
		} else {
			newChar = char + key
		}
	}

	return newChar
}

func (c *CeasarCipher) decryptChar(char int32) int32 {
	key := c.Key
	newChar := char
	if c.Key > 94 {
		key = c.Key % 94
	}

	if char-key < 32 {
		key = 32 - (char - key)
		newChar = 127 - key
	} else {
		newChar = char - key
	}

	return newChar
}

//Encrypt ...
func (c *CeasarCipher) Encrypt(plainText string) string {
	ascii := []int32(plainText)
	for i := 0; i < len(ascii); i++ {
		ascii[i] = c.encryptChar(ascii[i])
	}
	return string(ascii)
}

//Decrypt ...
func (c *CeasarCipher) Decrypt(plainText string) string {
	ascii := []int32(plainText)
	for i := 0; i < len(ascii); i++ {
		ascii[i] = c.decryptChar(ascii[i])
	}
	return string(ascii)
}
