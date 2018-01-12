package main

import (
	"encoding/hex"
	"fmt"
	"encoding/base64"
)

func main() {
	encodedString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	decodedString := XORagainstLetter(encodedString, "X")
	sDec, _ := base64.StdEncoding.DecodeString(hexToBase64(decodedString))
	fmt.Println(string(sDec))

}

func hexToBase64(hexValue string) string {
	decodedBytes, err := hex.DecodeString(hexValue)
	if err != nil {
		fmt.Println("Error", err)
	}

	base64Value := base64.StdEncoding.EncodeToString(decodedBytes)
	return base64Value
}

func XORagainstLetter(base, cipher string) string {
	var res []byte
	baseBytes, _ := hex.DecodeString(base)
	cipherByte := fmt.Sprintf("%s", cipher)

	for i := 0; i < len(baseBytes); i++ {
		res = append(res, baseBytes[i]^cipherByte[0])
	}

	return hex.EncodeToString(res)
}
