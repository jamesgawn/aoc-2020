package day6

import (
	"aoc2020/pkg/test"
	"log"
	"os"
	"testing"
)

func TestParseGroupCustomsDeclarationValid1(t *testing.T) {
	input := []string{
		"abc",
	}
	customsDeclaration := parseGroupCustomsDeclaration(input)
	test.EqualsInt(t, 3, len(customsDeclaration.anyYes))
}

func TestParseGroupCustomsDeclarationValid2(t *testing.T) {
	input := []string{
		"a",
		"b",
		"c",
	}
	customsDeclaration := parseGroupCustomsDeclaration(input)
	test.EqualsInt(t, 3, len(customsDeclaration.anyYes))
}

func TestParseGroupCustomsDeclarationValid3(t *testing.T) {
	input := []string{
		"ab",
		"ac",
	}
	customsDeclaration := parseGroupCustomsDeclaration(input)
	test.EqualsInt(t, 3, len(customsDeclaration.anyYes))
}

func TestParseGroupCustomsDeclarationValid4(t *testing.T) {
	input := []string{
		"a",
		"a",
		"a",
		"a",
	}
	customsDeclaration := parseGroupCustomsDeclaration(input)
	test.EqualsInt(t, 1, len(customsDeclaration.anyYes))
}

func TestParseGroupCustomsDeclarationValid5(t *testing.T) {
	input := []string{
		"b",
	}
	customsDeclaration := parseGroupCustomsDeclaration(input)
	test.EqualsInt(t, 1, len(customsDeclaration.anyYes))
}

func TestExecuteSolution(t *testing.T) {
	file, err := os.Open("./input_test.txt")
	if err != nil {
		log.Println("I can't solve every problem... not yet anyway")
		log.Fatal(err)
	}

	result := ExecuteSolution(file)
	test.EqualsInt(t, 11, result)

	cErr := file.Close()
	if cErr != nil {
		log.Fatal(cErr)
	}
}
