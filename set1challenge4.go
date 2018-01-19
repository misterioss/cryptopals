package main

import (
	"io/ioutil"

	"fmt"
	"strings"
	"encoding/hex"
)

func main() {
	data, err := ioutil.ReadFile("./set1challenge4.txt")
	check(err)
	strs := strings.Split(string(data), "\n");
	for i := 0; i < len(strs); i++ {
		s := strs[i]

		makeMeHappy(s);
	}
}

func check(err error) {
	if err != nil {
		panic(err)
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

func makeMeHappy(s string) {
	checkStrings(s, 13, 88, 10)
}

func checkStrings(s string, from, to int, threshhold float64) {
	for i := from; i <= to; i++ {
		xored := XORagainstLetter(s, byte(i))
		decoded, _ := hex.DecodeString(xored)
		resultString := fmt.Sprintf("%s", decoded)
		score := scoreText(resultString)
		if score >= threshhold {
			fmt.Println(s, score, resultString, i)
			fmt.Println(score)
		}
		score = 0
	}
}

func scoreText(s string) float64 {
	stringToScore := strings.Split(strings.ToUpper(s), "")
	totalScore := 0.0
	letterFrequency := make(map[string]float64)
	letterFrequency["E"] = 12.02
	letterFrequency["T"] = 9.10
	letterFrequency["A"] = 8.12
	letterFrequency["O"] = 7.68
	letterFrequency["I"] = 7.31
	letterFrequency["N"] = 6.95
	letterFrequency["S"] = 6.28
	letterFrequency["R"] = 6.02
	letterFrequency["H"] = 5.92
	letterFrequency["D"] = 4.32
	letterFrequency["L"] = 3.98
	letterFrequency["U"] = 2.88
	letterFrequency["C"] = 2.71
	letterFrequency["M"] = 2.61
	letterFrequency["F"] = 2.30
	letterFrequency["Y"] = 2.11
	letterFrequency["W"] = 2.09
	letterFrequency["G"] = 2.03
	letterFrequency["P"] = 1.82
	letterFrequency["B"] = 1.49
	letterFrequency["V"] = 1.11
	letterFrequency["K"] = 0.69
	letterFrequency["X"] = 0.17
	letterFrequency["Q"] = 0.11
	letterFrequency["J"] = 0.10
	letterFrequency["Z"] = 0.07

	for i := 0; i < len(stringToScore); i++ {
		letter := stringToScore[i]

		val, prs := letterFrequency[letter];
		if (prs) {
			totalScore += val
		} else {
			totalScore -= 15
		}
	}

	return totalScore
}
