package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	arr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			arr = append(arr, text)
		} else {
			break
		}
	}

	fmt.Println(arr)
}
