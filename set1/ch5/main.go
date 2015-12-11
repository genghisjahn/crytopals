package main

//http://cryptopals.com/sets/1/challenges/5/

import (
	"encoding/hex"
	"fmt"
)

func buildkey(rkey string, target string) string {
	result := ""
	rlen := len(rkey)
	tlen := len(target)
	num := tlen / rlen
	for i := 0; i < num; i++ {
		result += rkey

	}
	m := tlen % rlen
	if m > 0 {
		result += rkey[0:m]
	}
	return result
}

func main() {

	line1 := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

	basekey := "ICE"
	key := buildkey(basekey, line1)
	result1 := xorstrings(line1, key)
	fmt.Println(hex.EncodeToString(result1))
}

func xorstrings(s1, s2 string) []byte {
	h1 := hex.EncodeToString([]byte(s1))
	h2 := hex.EncodeToString([]byte(s2))
	hexData1, err1 := hex.DecodeString(h1)
	if err1 != nil {
		fmt.Println("1", err1)
		panic(err1)
	}

	hexData2, err2 := hex.DecodeString(h2)
	if err2 != nil {
		fmt.Println("2", err2)
		panic(err2)
	}
	result := []byte{}
	if len(hexData1) == len(hexData2) {
		for k := range hexData1 {
			r := hexData1[k] ^ hexData2[k]
			result = append(result, r)
		}
	}
	return result
}

func ScorePhrase(words string) int {
	result := 0
	freq := make(map[string]int)
	check1 := []string{"E", "T", "A"}
	check2 := []string{"O", "I", "N"}
	check3 := []string{"S", "H", "R"}
	check4 := []string{"D", "L", "U", " "}
	_, _, _ = check2, check3, check4
	length := float64(len(words))
	for _, v := range words {
		k := string(v)
		if c, ok := freq[k]; ok {
			freq[k] = c + 1
		} else {
			freq[k] = 1
		}
	}

	for _, c := range check1 {
		if float64(freq[c])/length > .017 {
			result += 4
		}

	}
	for _, c := range check2 {
		if float64(freq[c])/length > .017 {
			result += 3
		}
	}
	for _, c := range check3 {
		if float64(freq[c])/length > .017 {
			result += 2
		}
	}
	for _, c := range check4 {
		if float64(freq[c])/length > .017 {
			result += 1
		}
	}

	return result
}
