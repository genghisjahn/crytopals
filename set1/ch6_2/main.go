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

	encryptedText, _ := base64.StdEncoding.DecodeString(ciphertext)
	lenscores := getKeyLengthScores(40, 2, 41, string(encryptedText))
	for _, s := range lenscores[0:3] {
		blocks := blocksplit([]byte(encryptedText), s.Length)
		tblocks := transpose(blocks, s.Length)
		_ = tblocks
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

func getKeyLengthScores(topcount, minlen, maxlen int, txt string) []keylengthscore {
	scores := klscores{}
	for i := minlen; i < maxlen; i++ {
		btext := []byte(txt)
		first := btext[0:i]
		second := btext[i : 2*i]
		third := btext[2*i : 3*i]
		fourth := btext[3*i : 4*i]

		s1 := string(first)
		s2 := string(second)
		s3 := string(third)
		s4 := string(fourth)
		hd1 := strHamDist(s1, s2)
		hd2 := strHamDist(s3, s4)
		normHD1 := float64(hd1) / float64(i)
		normHD2 := float64(hd2) / float64(i)
		avg := (normHD1 + normHD2) / 2.0
		_ = avg
		kls := keylengthscore{i, normHD1}
		scores = append(scores, kls)
	}
	sort.Sort(scores)
	return scores[0:topcount]
}

func strHamDist(s1, s2 string) int {
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
		// if v < 32 || v > 127 {
		// 	continue
		// }
		if c, ok := freq[k]; ok {
			freq[k] = c + 1
		} else {
			freq[k] = 1
		}
	}

	for _, c := range check1 {
		if float64(freq[c])/length > .017 {
			result += 1
		}

	}
	for _, c := range check2 {
		if float64(freq[c])/length > .017 {
			result += 1
		}
	}
	for _, c := range check3 {
		if float64(freq[c])/length > .017 {
			result += 1
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

var ciphertext = `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`
