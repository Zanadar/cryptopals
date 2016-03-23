package set1

import (
	"encoding/hex"
	"testing"
)

func TestChall1(t *testing.T) {
	result, _ :=
		Chall1("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	shouldBe := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	if result != shouldBe {
		t.Errorf("Expected %s and got %s", shouldBe, result)
	}

}

func TestChall2(t *testing.T) {
	input := "1c0111001f010100061a024b53535009181c"
	inputBytes, _ := hex.DecodeString(input)
	XORagainst := "686974207468652062756c6c277320657965"
	otherBytes, _ := hex.DecodeString(XORagainst)
	shouldBe := "746865206b696420646f6e277420706c6179"

	result, _ := XOR(inputBytes, otherBytes)
	hexString := hex.EncodeToString(result)

	if hexString != shouldBe {
		t.Errorf("Expected %s and got %s", shouldBe, result)
	}
}
