package main

import (
	"bufio"
	"os"
	"strconv"
)

type passwordForVerification struct {
	minCharacterRequired int
	maxCharacterRequired int
	requiredCharacter    string
	password             string
}

func main() {
	arr := make([]passwordForVerification, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			minChar, _ := strconv.Atoi(text[0:1])
			maxChar, _ := strconv.Atoi(text[2:3])
			requiredCharacter := text[4:5]
			password := text[7:]
			arr = append(arr, passwordForVerification{minChar, maxChar, requiredCharacter, password})
		} else {
			break
		}
	}
}
