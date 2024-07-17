package set1

import (
	"bytes"
	"fmt"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	output := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	want := []byte(output)

	got := hexToBase64([]byte(input))

	if !bytes.Equal(want, got) {
		t.Fatalf("hex to base64 failed!\nGot: %v\nWanted: %v", got, want)
	}

	fmt.Printf("Success! Output: %v\n", string(got))

}
