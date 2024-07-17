package set1

func xorHex(a, b []byte) []byte {
	var aBinary []int
	var bBinary []int

	for i, _ := range a {
		aBinary = append(aBinary, hexCharToBinary(a[i])...)
		bBinary = append(bBinary, hexCharToBinary(b[i])...)
	}

	xorBinary := xorBinary(aBinary, bBinary)

	xorHex := binaryToHex(xorBinary)

}

func xorBinary(a, b []int) []int {
	var result []int

	for i, _ := range a {
		result[i] = a[i] ^ b[i] // xor operator
	}

	return result
}
