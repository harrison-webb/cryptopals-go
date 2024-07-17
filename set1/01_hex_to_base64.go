package set1

func hexToBase64(hex []byte) []byte {
	var binary []int
	for _, symbol := range hex {
		// convert hex symbol to binary and add it to the 'binary' array
		binary = append(binary, hexCharToBinary(symbol)...)
	}

	var base64 []byte
	// now the entire input hex string is represented as binary and can be converted to base 64
	for i := 0; i < len(binary); i += 6 {
		// loop over 'binary' array 6 bits at a time (base64 symbols are represented with 6 bits)
		sixBits := binary[i : i+6]
		base64 = append(base64, binaryToBase64(sixBits)...)
	}

	return base64
}

func hexCharToBinary(char byte) []int {
	var decimal int
	if isDigit(char) {
		decimal = int(char) - 48
	} else if isLetter(char) {
		decimal = int(char) - 87 // decimal value of 'a' in hex is 10
	} else {
		panic("Invalid byte recieved in charToBinary: not a letter or digit")
	}

	binary := decimalToBinaryNibble(decimal)
	return binary
}

// Return true if the provided char is a digit (0-9)
func isDigit(char byte) bool {
	return char >= 48 && char <= 57
}

// Return true if the provided char is a letter (a-z or A-Z)
func isLetter(char byte) bool {
	// ascii codes for lowercase: 97-122  ascii codes for uppercase: 65-90
	return (char >= 97 && char <= 122) || (char >= 65 && char <= 90)
}

// Convert an int to binary (represented as []int)
// Only intended for values < 16
func decimalToBinaryNibble(decimal int) []int {
	if decimal > 15 {
		panic("decimalToBinary: input > 15")
	}

	binaryDigitArray := []int{0, 0, 0, 0} // big endian
	i := 3                                // start inserting at right side of array and move left each iteration of loop

	for decimal > 0 {
		if i < 0 {
			panic("decimalToBinary: i became negative")
		}
		remainder := decimal % 2
		binaryDigitArray[i] = remainder
		decimal = decimal / 2
		i -= 1
	}

	return binaryDigitArray
}

func binaryToBase64(binary []int) []byte {
	decimal := 0
	exponent := len(binary) - 1 // iterate over 'binary' from L to R, so first val will represent 2^(len - 1). E.g. 1101 = 2^3 + 2^2 + 0 + 2^0

	for _, val := range binary {
		decimal += val * pow(2, exponent)
		exponent--
	}

	// take decimal value of base64 and convert to ascii value for that base64 digit ( 0 = A, 26 = a, 52 = 0)
	var result int
	if decimal >= 52 { // >=52 -> digit
		result = decimal - 4 // ascii code for 0 is 48, so to get the ascii digit for the specified base64 value we need to subtract 4
	} else if decimal >= 26 { // 26-51 -> lowercase letter
		result = decimal + 71 // base64 code for 'a' is 26, ascii code for 'a' is 97
	} else if decimal >= 0 {
		result = decimal + 65 // base64 code for 'A' is 0, ascii code for 'A' is 65
	} else {
		panic("binaryToBase64: invalid value received")
	}
	return []byte{byte(result)}
}

// x^y
func pow(x, y int) int {
	if y < 0 {
		panic("pow does not handle negative exponnents")
	}
	if y == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}

	base := x
	for y > 1 {
		x = base * x
		y--
	}

	return x
}
