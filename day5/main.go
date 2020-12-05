package day5

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strconv"
)

type seat struct {
	raw    string
	row    int
	column int
	id     int
}

func ExecuteSolution(input io.Reader) {
	seats := make([]seat, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()
		seat, _ := findSeat(text)
		seats = append(seats, seat)
	}
	largestSeatId := 0
	for _, seat := range seats {
		if seat.id > largestSeatId {
			largestSeatId = seat.id
		}
	}
	log.Println("Part 1 - Largest Seat ID:" + strconv.Itoa(largestSeatId))
}

func findSeat(text string) (seat, error) {
	seat, err := findRow(text, text, 0, 127)
	return calculateSeatId(seat), err
}

func calculateSeatId(seat seat) seat {
	seat.id = (seat.row * 8) + seat.column
	return seat
}

func findRow(fulltext string, text string, minRow int, maxRow int) (seat, error) {
	nextDivisor := text[0]
	numberOfRowsToRemove := ((maxRow - minRow) / 2) + 1
	if nextDivisor == 'F' {
		if len(text) == 4 {
			return findColumn(fulltext, text[1:], minRow, 0, 7)
		} else {
			return findRow(fulltext, text[1:], minRow, maxRow-numberOfRowsToRemove)
		}
	} else if nextDivisor == 'B' {
		if len(text) == 4 {
			return findColumn(fulltext, text[1:], maxRow, 0, 7)
		} else {
			return findRow(fulltext, text[1:], minRow+numberOfRowsToRemove, maxRow)
		}
	}
	return seat{}, errors.New("Unexpected input while finding row: " + string(nextDivisor))
}

func findColumn(fulltext string, text string, row int, minColumn int, maxColumn int) (seat, error) {
	nextDivisor := string(text[0])
	numberOfColumnsToRemove := ((maxColumn - minColumn) / 2) + 1
	if nextDivisor == "L" {
		if len(text) == 1 {
			return seat{
				raw:    fulltext,
				row:    row,
				column: minColumn,
			}, nil
		} else {
			return findColumn(fulltext, text[1:], row, minColumn, maxColumn-numberOfColumnsToRemove)
		}
	} else if nextDivisor == "R" {
		if len(text) == 1 {
			return seat{
				raw:    fulltext,
				row:    row,
				column: maxColumn,
			}, nil
		} else {
			return findColumn(fulltext, text[1:], row, minColumn+numberOfColumnsToRemove, maxColumn)
		}
	} else {
		return seat{}, errors.New("Unexpected input while finding column: " + string(nextDivisor))
	}
}
