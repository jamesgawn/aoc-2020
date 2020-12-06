package day6

import (
	"bufio"
	"io"
	"log"
	"strconv"
)

type customsDeclaration struct {
	answersByPerson []string
	anyYes          []string
	anyYesTotal     int
	allYes          []string
	allYesTotal     int
}

func ExecuteSolution(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	customsDeclarations := make([]customsDeclaration, 0)
	currentGroupAnswers := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			currentGroupAnswers = append(currentGroupAnswers, text)
		} else {
			customsDeclaration := parseGroupCustomsDeclaration(currentGroupAnswers)
			customsDeclarations = append(customsDeclarations, customsDeclaration)
			currentGroupAnswers = make([]string, 0)
		}
	}
	customsDeclaration := parseGroupCustomsDeclaration(currentGroupAnswers)
	customsDeclarations = append(customsDeclarations, customsDeclaration)

	totalYes := 0
	totalAllYes := 0
	for _, declaration := range customsDeclarations {
		totalYes = totalYes + declaration.anyYesTotal
		totalAllYes = totalAllYes + declaration.allYesTotal
	}
	log.Println("Part 1 - Any Yes Answers: " + strconv.Itoa(totalYes))
	log.Println("Part 2 - All Yes Answers: " + strconv.Itoa(totalAllYes))
	return totalYes
}

func parseGroupCustomsDeclaration(rawGroupData []string) customsDeclaration {
	uniqueGroupAnswers := parseYesAnswers(rawGroupData)
	allYesGroupAnswers := parseEveryoneAnsweredYes(rawGroupData)
	customsDeclaration := customsDeclaration{
		answersByPerson: rawGroupData,
		anyYes:          uniqueGroupAnswers,
		anyYesTotal:     len(uniqueGroupAnswers),
		allYes:          allYesGroupAnswers,
		allYesTotal:     len(allYesGroupAnswers),
	}
	return customsDeclaration
}

func parseEveryoneAnsweredYes(rawGroupData []string) []string {
	everyoneAnsweredYesLetters := make([]string, 0)
	for _, letter := range rawGroupData[0] {
		allAnsweredYes := true
		for _, currentPersonAnswers := range rawGroupData {
			hasLetter := false
			for _, currentAnswer := range currentPersonAnswers {
				if letter == currentAnswer {
					hasLetter = true
					break
				}
			}
			if hasLetter == false {
				allAnsweredYes = false
				break
			}
		}
		if allAnsweredYes {
			everyoneAnsweredYesLetters = append(everyoneAnsweredYesLetters, string(letter))
		}
	}
	return everyoneAnsweredYesLetters
}

func parseYesAnswers(rawGroupData []string) []string {
	uniqueGroupAnswers := make([]string, 0)
	for _, currentPersonAnswers := range rawGroupData {
		for _, currentAnswer := range currentPersonAnswers {
			uniqueGroupAnswers = appendIfMissing(uniqueGroupAnswers, string(currentAnswer))
		}
	}
	return uniqueGroupAnswers
}

func appendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
