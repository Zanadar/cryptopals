package set1

import (
	"encoding/hex"
	"reflect"
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

func TestCountLetters(t *testing.T) {
	t.Skip()
	result, keys := CountLetters("./fixtures/frankenstein.txt")
	for _, k := range keys {
		t.Log(k, result[k])
	}
}

func TestLetterPercents(t *testing.T) {
	t.Skip()
	result := LetterPercents("./fixtures/frankenstein.txt")
	for k, v := range result {
		t.Log(k, v)
	}
}

func TestScoreString(t *testing.T) {
	t.Skip()
	scoreThis := "AEIOU"
	if result := ScoreString(scoreThis); result != 5 {
		t.Errorf("Expected %d and got %d", 5, result)
	}
}

func TestChall3(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	results := Chall3(input)
	t.Logf("Winner is: %+v", results)
}

func TestChall4(t *testing.T) {
	inputFile := "./fixtures/4.txt"
	results := Chall4(inputFile)
	t.Logf("Winner is: %+v", results)
}

func TestChall5(t *testing.T) {
	toEncrypt := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	shouldBe := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	fillString := make([]byte, len([]byte(toEncrypt)))
	fillString = filler(fillString, []byte(key))
	result := hex.EncodeToString(Chall5(toEncrypt, key))

	if string(result) != shouldBe {
		t.Errorf("\n1: %s \n2: %s \nusing %q", shouldBe, result, fillString)
	}
}

func TestHammingDist(t *testing.T) {
	result := HammingDist([]byte("this is a test"), []byte("wokka wokka!!!"))
	if result != 37 {
		t.Error("Hamming Distance should be 37, but its", result)
	}
}

func TestFindKey(t *testing.T) {
	decode := DecodeBase64File("./fixtures/6.txt")
	results := FindKeySize(decode)
	for _, v := range results {
		t.Logf("Key is probably %v", v)
	}
}

func TestTransposeBlocks(t *testing.T) {
	result := TransposeBlocks([]byte{1, 2, 3, 4, 1, 2, 3, 4}, 4)
	shouldBe := [][]byte{{1, 1}, {2, 2}, {3, 3}, {4, 4}}
	if !reflect.DeepEqual(result, shouldBe) {
		t.Errorf("Should be %v but got %v", shouldBe, result)
	}
}

func TestChall6(t *testing.T) {
	test := DecodeBase64File("./fixtures/6.txt")
	// test, _ := hex.DecodeString("426c657463686579205061726b2072656a6f6963657320696e20746865206661637420746861742c20756e74696c20666169726c7920726563656e746c792c206974207761732070726f6261626c79204272697461696e27732062657374206b657074207365637265742e20546869732069732062656361757365206f6620746865207365637265637920737572726f756e64696e6720616c6c2074686520616374697669746965732063617272696564206f6e206865726520647572696e6720576f726c64205761722054776f20776173206f6620766974616c20696d706f7274616e636520746f206f7572206e6174696f6e616c20736563757269747920616e6420756c74696d61746520766963746f72792e")
	for i := 1; i < 41; i++ {
		result := CrackRepeatingXOR(test, i)
		t.Logf("%q", result)
	}

}
