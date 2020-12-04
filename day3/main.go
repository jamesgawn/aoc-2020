package day3

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func ExecuteSolution(input io.Reader) {
	arr := make([][]rune, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) != 0 {
			arr = append(arr, []rune(text))
		} else {
			break
		}
	}

	part1 := obtainTrees("part 1", arr, 1, 3)
	fmt.Println("Part 1: " + strconv.Itoa(part1))

	part21 := obtainTrees("part2.1", arr, 1, 1)
	part22 := obtainTrees("part2.2", arr, 1, 3)
	part23 := obtainTrees("part2.3", arr, 1, 5)
	part24 := obtainTrees("part2.4", arr, 1, 7)
	part25 := obtainTrees("part2.5", arr, 2, 1)
	part2 := part21 * part22 * part23 * part24 * part25
	fmt.Println("Part 2: " + strconv.Itoa(part2))
}

func obtainTrees(name string, course [][]rune, downSpeed int, rightSpeed int) int {
	height := len(course)
	width := len(course[0])
	numberOfTrees := 0
	i := 0
	for posY := 0; posY < height; posY = posY + downSpeed {
		rowArray := course[posY]
		posX := obtainXPosition(i, width, rightSpeed)
		value := string(rowArray[posX : posX+1])
		treePresent := value == "#"
		if treePresent {
			numberOfTrees++
		}
		fmt.Println("Course: " + name + " Y:" + strconv.Itoa(posY) + " X:" + strconv.Itoa(posX) + " V: " + value + " TreePresent: " + strconv.FormatBool(treePresent) + " TreeCount:" + strconv.Itoa(numberOfTrees))
		i++
	}
	fmt.Println("Course " + name + " tree count:" + strconv.Itoa(numberOfTrees))
	return numberOfTrees
}

func obtainXPosition(i int, width int, rightwardSpeed int) int {
	a := rightwardSpeed * i
	b := a % (width)
	if a < width {
		return a
	} else {
		return b
	}
}
