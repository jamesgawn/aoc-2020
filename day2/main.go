package day2

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type passwordForVerification struct {
	minCharacterRequired int
	maxCharacterRequired int
	requiredCharacter    string
	password             string
}

type passwordResult struct {
	passwordForVerification
	valid bool
}

func ExecuteSolution(input io.Reader) {
	scanner := bufio.NewScanner(input)
	arr := make([]passwordForVerification, 0)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			endOfMin := strings.Index(text, "-")
			startOfMax := endOfMin + 1
			endOfMax := strings.Index(text, " ")
			startOfChar := endOfMax + 1
			endOfChar := endOfMax + 2
			startOfPassword := endOfChar + 2
			minChar, _ := strconv.Atoi(text[0:endOfMin])
			maxChar, _ := strconv.Atoi(text[startOfMax:endOfMax])
			requiredCharacter := text[startOfChar:endOfChar]
			password := text[startOfPassword:]
			arr = append(arr, passwordForVerification{minChar, maxChar, requiredCharacter, password})
		} else {
			break
		}
	}
	part1ValidCount := 0
	part2ValidCount := 0
	for _, toVerify := range arr {
		resultPart1 := verifyPasswordValidityPart1(toVerify)
		if resultPart1.valid {
			part1ValidCount++
		}
		resultPart2 := verifyPasswordValidityPart2(toVerify)
		if resultPart2.valid {
			part2ValidCount++
		}
	}
	fmt.Println("Total Valid (Part 1): " + strconv.Itoa(part1ValidCount))
	fmt.Println("Total Valid (Part 2): " + strconv.Itoa(part2ValidCount))
}

func verifyPasswordValidityPart1(verifiable passwordForVerification) passwordResult {
	characterCount := 0
	requiredRune := verifiable.requiredCharacter
	for _, character := range verifiable.password {
		if string(character) == requiredRune {
			characterCount++
		}
	}
	return passwordResult{verifiable, characterCount >= verifiable.minCharacterRequired && characterCount <= verifiable.maxCharacterRequired}
}

func verifyPasswordValidityPart2(verifiable passwordForVerification) passwordResult {
	requiredRune := verifiable.requiredCharacter
	firstCharMatch := verifiable.password[verifiable.minCharacterRequired-1:verifiable.minCharacterRequired] == requiredRune
	secondCharMatch := verifiable.password[verifiable.maxCharacterRequired-1:verifiable.maxCharacterRequired] == requiredRune
	return passwordResult{verifiable, firstCharMatch != secondCharMatch}
}
