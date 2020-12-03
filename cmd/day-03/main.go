package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := make([][]rune, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			arr = append(arr, []rune(text))
		} else {
			break
		}
	}

	width := len(arr[0])
	fmt.Println("Width: " + strconv.Itoa(width))
	numberOfTrees := 0
	for posY, rowArray := range arr {
		posX := obtainXPosition(posY, width)
		value := string(rowArray[posX : posX+1])
		treePresent := value == "#"
		if treePresent {
			numberOfTrees++
		}
		fmt.Println("Y:" + strconv.Itoa(posY) + " X:" + strconv.Itoa(posX) + " V: " + value + " TreePresent: " + strconv.FormatBool(treePresent) + " TreeCount:" + strconv.Itoa(numberOfTrees))
	}
}

func obtainXPosition(posY int, width int) int {
	a := 3 * posY
	b := a % (width)
	if a <= width {
		return a
	} else {
		return b
	}
}
