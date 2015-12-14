package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	a := "this is a test"
	b := "wokka wokka!!!"

	if strHamDist(a, b) != 37 {
		panic("strHamDist func is broken!")
	}

	encryptedText, errDec := base64.StdEncoding.DecodeString(ciphertext)
	if errDec != nil {
		panic(errDec)
	}
	lenscores := getKeyLengthScores(1, 29, 30, string(encryptedText))
	for _, s := range lenscores {
		blocks := blocksplit([]byte(encryptedText), s.Length)
		tblocks := transpose(blocks, s.Length)
		bscores := scoreblocks(tblocks)
		key := ""
		for _, bs := range bscores {
			key += bs.Key
		}
		//fmt.Println("Length:", s.Length, "Key:", key)
		key = ""
	}

}

func getPossibleKey(sb skscores) string {
	result := ""
	for _, c := range sb {
		result += c.Key
	}
	return result
}

func scoreblocks(tblocks [][]byte) skscores {
	result := skscores{}
	for _, b := range tblocks {
		sks := processPhrase(string(b))
		result = append(result, sks)
	}
	sort.Sort(result)
	return result
}

func transpose(blocks [][]byte, size int) [][]byte {
	tblocks := [][]byte{}
	for i := 0; i < size; i++ {
		t := []byte{}
		tblocks = append(tblocks, t)
	}
	for _, v := range blocks {
		for a, b := range v {
			m := a % size
			tblocks[m] = append(tblocks[m], b)
		}

	}
	return tblocks
}

func blocksplit(source []byte, size int) [][]byte {
	r := [][]byte{}
	slen := len(source)
	for i := 0; i < slen; i += size {
		block := source[i : i+size]
		r = append(r, block)
	}

	return r
}

func getKeyLengthScores(topcount, minlen, maxlen int, data string) []keylengthscore {
	scores := klscores{}

	for i := minlen; i < maxlen; i++ {
		btext := []byte(data)
		first := btext[0:i]
		second := btext[i : 2*i]
		s1 := string(first)
		s2 := string(second)
		if len(first) != len(s1) {
			fmt.Println(first, s1, "!!")
		}
		if len(second) != len(s2) {
			fmt.Println(first, s2, "!!")
		}
		hd1 := strHamDist(s1, s2)
		hd2 := byteHamDist(first, second)
		normHD1 := float64(hd1) / float64(i)
		normHD2 := float64(hd2) / float64(i)
		kls := keylengthscore{i, normHD1}
		scores = append(scores, kls)
		fmt.Println(i, normHD2)
	}
	sort.Sort(scores)
	return scores[0:topcount]
}

func strHamDist(s1, s2 string) int {
	t := 0
	fmt.Println(len(s1), len(s2))
	for k, v := range s1 {
		n1 := int64(v)
		n2 := int64(s2[k])
		b1 := strconv.FormatInt(n1, 2)
		b2 := strconv.FormatInt(n2, 2)
		b1 = fmt.Sprintf("%07d", b1)
		b2 = fmt.Sprintf("%07d", b2)
		for k := range b1 {
			if b1[k] != b2[k] {
				t++
			}
		}

	}
	return t
}

func byteHamDist(s1, s2 []byte) int {
	t := 0
	for k, v := range s1 {
		n1 := int64(v)
		n2 := int64(s2[k])
		b1 := strconv.FormatInt(n1, 2)
		b2 := strconv.FormatInt(n2, 2)
		b1 = fmt.Sprintf("%07d", b1)
		b2 = fmt.Sprintf("%07d", b2)
		for k := range b1 {
			if b1[k] != b2[k] {
				t++
			}
		}

	}
	fmt.Println(len(s1), t)
	return t
}

func processPhrase(phrase string) singlekeyscore {
	r := singlekeyscore{}
	msg := phrase
	var key byte
	var highest = 0
	for _, v := range tkeys {
		key = byte(v)
		result := []byte{}

		for k := range msg {
			r := msg[k] ^ key
			result = append(result, r)
		}
		score := scorePhrase(strings.ToUpper(string(result)))
		if score >= highest {
			highest = score
			r.Key = string(key)
			r.RawText = phrase
			r.Score = score
			r.Decrypted = string(result)
		}

	}

	return r
}

func scorePhrase(words string) int {
	result := 0
	freq := make(map[string]int)
	check1 := []string{"E", "T", "A"}
	check2 := []string{"O", "I", "N"}
	check3 := []string{"S", "H", "R", " "}
	check4 := []string{"D", "L", "U"}
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
			result++
		}
	}

	return result
}

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
