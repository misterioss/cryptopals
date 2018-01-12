package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	shouldProduce := "746865206b696420646f6e277420706c6179"
	baseBuffer := "1c0111001f010100061a024b53535009181c"
	againstBuffer := "686974207468652062756c6c277320657965"
	fmt.Println(XOR(baseBuffer, againstBuffer) == shouldProduce)
}


func XOR(base, against string) string {
	var res []byte
	baseBytes, _ := hex.DecodeString(base)
	againstBytes, _ :=hex.DecodeString(against)

	for i:=0; i < len(baseBytes) ; i++  {
		res = append(res, baseBytes[i]^againstBytes[i])
	}

	return hex.EncodeToString(res)
}