package day5

import (
	"aoc2020/pkg/test"
	"testing"
)

func TestValidateFindSeatWithValidInput0(t *testing.T) {
	result, _ := findSeat("FBFBBFFRLR")
	test.EqualsInt(t, 44, result.row)
	test.EqualsInt(t, 5, result.column)
	test.EqualsInt(t, 357, result.id)
}

func TestValidateFindSeatWithValidInput1(t *testing.T) {
	result, _ := findSeat("BFFFBBFRRR")
	test.EqualsInt(t, 70, result.row)
	test.EqualsInt(t, 7, result.column)
	test.EqualsInt(t, 567, result.id)
}

func TestValidateFindSeatWithValidInput2(t *testing.T) {
	result, _ := findSeat("FFFBBBFRRR")
	test.EqualsInt(t, 14, result.row)
	test.EqualsInt(t, 7, result.column)
	test.EqualsInt(t, 119, result.id)
}

func TestValidateFindSeatWithValidInput3(t *testing.T) {
	result, _ := findSeat("BBFFBBFRLL")
	test.EqualsInt(t, 102, result.row)
	test.EqualsInt(t, 4, result.column)
	test.EqualsInt(t, 820, result.id)
}
