package set1

func hexToBase64(hex []byte) []byte {
	var binary []int
	for _, symbol := range hex {
		// convert hex symbol to binary and add it to the 'binary' array
		binary = append(binary, charToBinary(symbol)...)
	}

	// now the entire input hex string is represented as binary and can be converted to base 64
	for i := 0; i < len(binary); i += 6 {
		// loop over 'binary' array 6 bits at a time (base64 symbols are represented with 6 bits)
	}
}

func charToBinary(char byte) []int {
	var decimal int
	if isDigit(char) {
		decimal = int(char) - 48
	} else if isLetter(char) {
		decimal = int(char) - 87 // decimal value of 'a' in hex is 10
	} else {
		panic("Invalid byte recieved in charToBinary: not a letter or digit")
	}

	return decimalToBinary(decimal)
}

func isDigit(char byte) bool {
	return char >= 48 && char <= 57
}

func isLetter(char byte) bool {
	// ascii codes for lowercase: 97-122 // ascii codes for uppercase: 65-90
	return (char >= 97 && char <= 122) || (char >= 65 && char <= 90)
}

func decimalToBinary(decimal int) []int {
	binaryDigitArray := []int{0, 0, 0, 0} // big endian
	i := 3                                // start inserting at right side of array and move left each iteration of loop

	for decimal > 0 {
		remainder := decimal % 2
		binaryDigitArray[i] = remainder
		decimal = decimal / 2
		i -= 1
		if i < 0 {
			panic("decimalToBinary: i became negative")
		}
	}

	return binaryDigitArray
}
