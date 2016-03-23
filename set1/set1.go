package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Chall1(hexString string) (string, error) {
	hexBytes, err := hex.DecodeString(hexString)
	output := base64.StdEncoding.EncodeToString(hexBytes)
	return output, err
}

func XOR(input1, input2 []byte) ([]byte, error) {
	ret := make([]byte, len(input2))
	for i := 0; i < len(input1); i++ {
		ret[i] = input1[i] ^ input2[i]
	}
	return ret, nil
}

func main() {
	hexString :=
		"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	out, _ := Chall1(hexString)
	fmt.Printf("Value is %v", out)

}
