package main

//http://cryptopals.com/sets/1/challenges/1/

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	data := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	hexData, err := hex.DecodeString(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	sEnc := base64.StdEncoding.EncodeToString(hexData)
	fmt.Println(sEnc)
}
