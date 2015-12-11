package main

//http://cryptopals.com/sets/1/challenges/2/

import (
	"encoding/hex"
	"fmt"
)

func main() {
	msg1 := "1c0111001f010100061a024b53535009181c"
	msg2 := "686974207468652062756c6c277320657965"
	hexData1, err1 := hex.DecodeString(msg1)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	hexData2, err2 := hex.DecodeString(msg2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	result := []byte{}
	if len(hexData1) == len(hexData2) {
		for k := range hexData1 {
			r := hexData1[k] ^ hexData2[k]
			result = append(result, r)
		}
	}
	fmt.Println(hex.EncodeToString(result))

}
