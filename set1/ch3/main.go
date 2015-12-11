package main

//http://cryptopals.com/sets/1/challenges/3/

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	tkeys := "abcdefghjklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	msg := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	var key byte
	var highest = 0
	var highkey string
	var highmsg string
	for _, v := range tkeys {
		key = byte(v)

		hexData, err := hex.DecodeString(msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		result := []byte{}

		for k := range hexData {
			r := hexData[k] ^ key
			result = append(result, r)
		}
		score := ScorePhrase(strings.ToUpper(string(result)))
		if score > highest {
			highest = score
			highkey = string(key)
			highmsg = string(result)
		}

		if score > 20 {
			//fmt.Println("Msg:", string(result), string(key), score)
		}

	}
	fmt.Println("Best Candidate:", highkey, highest, highmsg)
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
