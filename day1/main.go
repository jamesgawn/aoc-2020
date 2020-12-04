package day1

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strconv"
	"time"
)

func ExecuteSolution(input io.Reader) {
	scanner := bufio.NewScanner(input)
	arr := make([]int, 0)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			num, _ := strconv.Atoi(text)
			arr = append(arr, num)
		} else {
			break
		}
	}
	n1, n2, _ := findPair(arr, 2020)
	resultPair := n1 * n2
	log.Println("Pair - Matching pair numbers are " + strconv.Itoa(n1) + " & " + strconv.Itoa(n2))
	log.Println("Pair - Result: " + strconv.Itoa(resultPair))
	n1, n2, n3, _ := findTruple(arr, 2020)
	resultTruple := n1 * n2 * n3
	log.Println("Truple - Matching pair numbers are " + strconv.Itoa(n1) + ", " + strconv.Itoa(n2) + " & " + strconv.Itoa(n3))
	log.Println("Truple - Result: " + strconv.Itoa(resultTruple))
}

func findPair(array []int, sumMatch int) (int, int, error) {
	defer timeTrack(time.Now(), "findPair")
	for _, n1 := range array {
		for _, n2 := range array {
			var sum = n1 + n2
			if sum == sumMatch {
				return n1, n2, nil
			}
		}
	}
	return -1, -1, errors.New("unable to find matching values")
}

func findTruple(array []int, sumMatch int) (int, int, int, error) {
	defer timeTrack(time.Now(), "findTruple")
	for _, n1 := range array {
		for _, n2 := range array {
			for _, n3 := range array {
				var sum = n1 + n2 + n3
				if sum == sumMatch {
					return n1, n2, n3, nil
				}
			}
		}
	}
	return -1, -1, -1, errors.New("unable to find matching values")
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
