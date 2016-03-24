package set1

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
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

func ScoreString(test string) (score int) {
	against := "ETAOIN SHRDLU etaoin shrdlu"
	scoreRaw := 0
	for i := 0; i < len(test); i++ {
		letter := string(test[i])
		if strings.Contains(against, letter) {
			scoreRaw++
		}
	}

	return scoreRaw
}

func filler(arr []byte, with byte) []byte {
	for i := range arr {
		arr[i] = with
	}
	return arr
}

func doDecrypt3(hexString string) []Decrypt {
	hexBytes, _ := hex.DecodeString(hexString)
	// put all the btes for a-z, ' ', and A-Z in a slice
	letters := make([]byte, 53)
	for i := 65; i < 91; i++ {
		letters = append(letters, byte(i))
	}
	for i := 97; i < 123; i++ {
		letters = append(letters, byte(i))
	}
	letters = append(letters, byte(32))

	scores := make([]Decrypt, 0)
	for _, v := range letters {
		tester := make([]byte, len(hexBytes))
		tester = filler(tester, v)
		result, _ := XOR(hexBytes, tester)
		score := ScoreString(string(result))
		scores = append(scores, Decrypt{v, result, score})
	}
	return scores
}

type Decrypt struct {
	Against byte
	Result  []byte
	Score   int
}

func (d Decrypt) String() string {
	return fmt.Sprintf("%q got score: %d against %q", d.Result, d.Score, d.Against)
}

type ByScore []Decrypt

func (s ByScore) Len() int           { return len(s) }
func (s ByScore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByScore) Less(i, j int) bool { return s[i].Score < s[j].Score }

func Chall3(hexString string) Decrypt {
	results := doDecrypt3(hexString)
	for _, result := range results {
		fmt.Println(result)
	}
	sort.Sort(ByScore(results))

	return results[len(results)-1]
}
