package main

import (
	"encoding/hex"
	"fmt"
	"encoding/base64"
)

func main() {
	hexValue := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Println(hexToBase64(hexValue))
}

func hexToBase64(hexValue string) string {
	decodedBytes, err := hex.DecodeString(hexValue)
	if err != nil {
		fmt.Println("Error", err)
	}

	base64Value := base64.StdEncoding.EncodeToString(decodedBytes)
	return base64Value
}
