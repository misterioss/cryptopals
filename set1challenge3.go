package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	encodedString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	codeA := 65
	codeX := 88
	codea := 97
	codex := 120

	for i := codeA; i <= codeX; i++ {
		xored := XORagainstLetter(encodedString, byte(i))
		decoded, _ := hex.DecodeString(xored)
		fmt.Printf("%s\n", decoded)
	}

	fmt.Println("\n\nSmall letters \n\n")
	for i := codea; i <= codex; i++ {
		xored := XORagainstLetter(encodedString, byte(i))
		decoded, _ := hex.DecodeString(xored)
		fmt.Printf("%s\n", decoded)
	}
}

func XORagainstLetter(base string, cipher byte) string {
	var res []byte
	baseBytes, _ := hex.DecodeString(base)

	for i := 0; i < len(baseBytes); i++ {
		res = append(res, baseBytes[i]^cipher)
	}

	return hex.EncodeToString(res)
}
