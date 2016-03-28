package set1

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

var charFreq = map[string]float64{
	"A": 8.167,
	"B": 1.492,
	"C": 2.782,
	"D": 4.253,
	"E": 12.702,
	"F": 2.228,
	"G": 2.015,
	"H": 6.094,
	"I": 6.966,
	"J": 0.153,
	"K": 0.772,
	"L": 4.025,
	"M": 2.406,
	"N": 6.749,
	"O": 7.507,
	"P": 1.929,
	"Q": 0.095,
	"R": 5.987,
	"S": 6.327,
	"T": 9.056,
	"U": 2.758,
	"V": 0.978,
	"W": 2.360,
	"X": 0.150,
	"Y": 1.974,
	"Z": 0.074,
	" ": 20.0,
}

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

func CountLetters(filename string) (map[string]int, []string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error with file:", err)
	}

	var counts = make(map[string]int)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		if char, _ := utf8.DecodeRuneInString(scanner.Text()); unicode.IsLetter(rune(char)) {
			counts[scanner.Text()] += 1
		}
	}
	keys := make([]string, 52)
	for k, _ := range counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return counts, keys
}

func LetterPercents(filename string) map[string]float64 {
	results, keys := CountLetters(filename)
	total := 0
	for _, v := range keys {
		total += results[v]
	}
	percents := make(map[string]float64)
	for k, v := range results {
		percent := float64(v) / float64(total)
		percents[k] = percent
	}
	return percents
}

func ScoreString(test string) float64 {
	scoreRaw := 0.0
	for i := 0; i < len(test); i++ {
		letter := string(test[i])
		letter = strings.ToUpper(letter)
		score, ok := charFreq[letter]
		if ok {
			scoreRaw += score
		}
	}
	return scoreRaw
}

func filler(arr []byte, with []byte) []byte {
	for i := range arr {
		arr[i] = with[i%len(with)] //repeat 'with' bytes
	}
	return arr
}

func doDecrypt3(cypher []byte) []Decrypt {
	scores := make([]Decrypt, 0)
	for i := 0; i < 256; i++ {
		tester := make([]byte, len(cypher))
		tester = filler(tester, []byte{byte(i)})
		result, _ := XOR(cypher, tester)
		score := ScoreString(string(result))
		scores = append(scores, Decrypt{byte(i), result, score})
	}
	return scores
}

type Decrypt struct {
	Against byte
	Result  []byte
	Score   float64
}

func (d Decrypt) String() string {
	return fmt.Sprintf("%q got score: %d against %q", d.Result, d.Score, d.Against)
}

type ByScore []Decrypt

func (s ByScore) Len() int           { return len(s) }
func (s ByScore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByScore) Less(i, j int) bool { return s[i].Score < s[j].Score }

func Chall3(hexString string) Decrypt {
	temp, _ := hex.DecodeString(hexString)
	stringBytes := []byte(temp)
	results := doDecrypt3(stringBytes)
	// for _, result := range results {
	// fmt.Println(result)
	// }
	sort.Sort(ByScore(results))

	return results[len(results)-1]
}

func Chall4(filename string) Decrypt {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error with file:", err)
	}

	var lineDecodes = make([]Decrypt, 0)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := Chall3(scanner.Text())
		lineDecodes = append(lineDecodes, line)
	}
	sort.Sort(ByScore(lineDecodes))

	return lineDecodes[len(lineDecodes)-1]
}

func Chall5(plaintext string, key string) []byte {
	bytesString := []byte(plaintext)
	keyXOR := make([]byte, len(bytesString))
	keyXOR = filler(keyXOR, []byte(key))
	result, _ := XOR(bytesString, keyXOR)
	return result
}

func HammingDist(first []byte, second []byte) int {
	// This is a pretty janky implementation but it works...
	diffs, _ := XOR(first, second)
	count := 0
	for _, b := range diffs {
		for b > 1 {
			b &= b - 1
			count++
		}
	}
	return count
}

type KeyResult struct {
	size          int
	normalHamming float64
}

type ByDist []KeyResult

func (d ByDist) Len() int           { return len(d) }
func (d ByDist) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d ByDist) Less(i, j int) bool { return d[i].normalHamming < d[j].normalHamming }

func DecodeBase64File(path string) []byte {
	file, _ := os.Open(path)
	fInfo, _ := file.Stat()
	fSize := fInfo.Size()
	decodeSize := base64.StdEncoding.DecodedLen(int(fSize))
	data := make([]byte, fSize)
	_, _ = file.Read(data)
	base64Bytes := make([]byte, decodeSize)
	base64.StdEncoding.Decode(base64Bytes, data)
	defer file.Close()
	return base64Bytes
}

func FindKeySize(data []byte) []KeyResult {
	results := make([]KeyResult, 0)
	for j := 1; j < 41; j++ {
		readBlock := bytes.NewBuffer(data)
		firstBlock := readBlock.Next(j)
		secondBlock := readBlock.Next(j)
		thirdBlock := readBlock.Next(j)
		fourhtBlock := readBlock.Next(j)
		fifthBlock := readBlock.Next(j)
		sixthBlock := readBlock.Next(j)
		seventhBlock := readBlock.Next(j)
		dists := make([]int, 6)
		dists[0] = HammingDist(firstBlock, secondBlock)
		dists[1] = HammingDist(secondBlock, thirdBlock)
		dists[2] = HammingDist(thirdBlock, fourhtBlock)
		dists[3] = HammingDist(fourhtBlock, fifthBlock)
		dists[4] = HammingDist(fifthBlock, sixthBlock)
		dists[5] = HammingDist(sixthBlock, seventhBlock)

		normalAvgDist := 0.0
		sum := 0.0

		for _, v := range dists {
			sum += float64(v)
		}
		sum /= 6
		normalAvgDist = sum / float64(j)

		results = append(results, KeyResult{j, normalAvgDist})
	}
	sort.Sort(ByDist(results))
	return results
}

func TransposeBlocks(cypher []byte, keyLength int) [][]byte {
	blocks := make([][]byte, keyLength)
	for i, _ := range blocks {
		blocks[i] = make([]byte, 0)
	}
	cypherBuf := bytes.NewBuffer(cypher)
	for cypherBuf.Len() > 0 {
		block := cypherBuf.Next(keyLength)
		for i, v := range block {
			blocks[i] = append(blocks[i], v)
		}
	}

	return blocks
}

func CrackRepeatingXOR(cypher []byte, keyLength int) []byte {
	blocks := TransposeBlocks(cypher, keyLength)
	// read cyper in KeyLength blocks. For each block, append block[0] to blocks[0]..block[n] to blocks[n]
	cracks := []Decrypt{}
	for _, v := range blocks {
		decrypt := doDecrypt3(v)
		sort.Sort(ByScore(decrypt))
		cracks = append(cracks, decrypt[len(decrypt)-1])
	}
	// for _, crack := range cracks {
	// fmt.Print("\nScore:", crack.Score, "\n\n")
	// fmt.Println(crack)
	// }
	key := []byte{}
	for _, crack := range cracks {
		key = append(key, crack.Against)
	}
	return key
}
