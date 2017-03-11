package ceasarcipher

//MultiCeasar ...
type MultiCeasar struct {
	Key [][2]int32
}

func (mc *MultiCeasar) encryptChar(char int32, level int) int32 {
	key := mc.Key[level][1]
	newChar := char
	if key > 52 {
		key = key % 52
	}

	if char < 65 || char > 122 {
		newChar = char
	} else if char > 90 && char < 97 {
		newChar = char
	} else {
		if c := char + key; c > 90 && c < 97 {
			offset := (char + key) - 91
			newChar = 97 + offset
		} else if key+char > 122 {
			key = (key + char) - 126
			newChar = 65 + key
			if c := newChar; c > 90 && c < 97 {
				offset := 91 - newChar
				newChar = 97 + offset
			}
		} else {
			newChar = char + key
		}
	}
	return newChar
}

func (mc *MultiCeasar) decryptChar(char int32, level int) int32 {
	key := mc.Key[level][1]
	newChar := char
	if key > 52 {
		key = key % 52
	}

	if char < 65 || char > 122 {
		newChar = char
	} else if char > 90 && char < 97 {
		newChar = char
	} else {
		// case 1: char - key < 65
		// -> 70 - 6 = 64 < 65 by 1
		// -> offset = 65 - (char - key)
		// -> 123 - offset
		// case 2: char - key between 90 && 97
		// -> 100 k=5 -> 95 <-> 90,97 offset by 2
		// -> offset = 97 - (char - key)
		// -> newChar = 90 - offset
		if char-key < 65 {
			key = 65 - (char - key)
			newChar = 123 - key
			if c := newChar; c > 90 && c < 97 {
				offset := 97 - newChar
				newChar = 91 - offset
			}
		} else if c := char - key; c > 90 && c < 97 {
			offset := 97 - (char - key)
			newChar = 91 - offset
		} else {
			newChar = char - key
		}
	}

	return newChar
}

//Encrypt ...
func (mc *MultiCeasar) Encrypt(plainText string) string {
	ascii := []int32(plainText)
	for i := 0; i < len(mc.Key); i++ {
		for j := mc.Key[i][0]; j < int32(len(ascii)); j++ {
			ascii[j] = mc.encryptChar(ascii[j], i)
		}
	}
	return string(ascii)
}

//Decrypt ...
func (mc *MultiCeasar) Decrypt(plainText string) string {
	// "enc text"
	// -> [# # # # #]
	// -> decryptChar(char)
	//   -> 67 k=5,0,k=3,5
	ascii := []int32(plainText)
	for i := 0; i < len(mc.Key); i++ {
		for j := mc.Key[i][0]; j < int32(len(ascii)); j++ {
			ascii[j] = mc.decryptChar(ascii[j], i)
		}
	}
	return string(ascii)
}
