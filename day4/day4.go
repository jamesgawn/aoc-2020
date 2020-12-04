package day4

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

type passportResult struct {
	passport
	valid bool
}

func ExecuteSolution(input io.Reader) {
	passports := ParsePassport(input)
	validatedPassports := ValidatePassport(passports)
	validPassports := 0
	for _, validatedPassport := range validatedPassports {
		PassportPrinter(validatedPassport)
		if validatedPassport.valid {
			validPassports++
		}
	}
	fmt.Println("Total Valid Passports: " + strconv.Itoa(validPassports))
}

func ParsePassport(input io.Reader) []passport {
	passportArray := make([]passport, 0)
	currentPassport := passport{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			dataChunks := strings.Split(text, " ")
			for _, dataChunk := range dataChunks {
				data := strings.Split(dataChunk, ":")
				key := string(data[0])
				value := data[1]
				if key == "byr" {
					currentPassport.byr = value
				}
				if key == "iyr" {
					currentPassport.iyr = value
				}
				if key == "eyr" {
					currentPassport.eyr = value
				}
				if key == "hgt" {
					currentPassport.hgt = value
				}
				if key == "hcl" {
					currentPassport.hcl = value
				}
				if key == "ecl" {
					currentPassport.ecl = value
				}
				if key == "pid" {
					currentPassport.pid = value
				}
				if key == "cid" {
					currentPassport.cid = value
				}
			}
		} else {
			passportArray = append(passportArray, currentPassport)
			currentPassport = passport{}
		}
	}
	return passportArray
}

func ValidatePassport(passports []passport) []passportResult {
	passportResults := make([]passportResult, 0)
	for _, passport := range passports {
		valid := passport.byr != "" && passport.iyr != "" && passport.eyr != "" && passport.hgt != "" && passport.hcl != "" && passport.ecl != "" && passport.pid != ""
		passportResults = append(passportResults, passportResult{passport, valid})
	}
	return passportResults
}

func PassportPrinter(passport passportResult) {
	fmt.Println("byr: " + passport.byr + " iyr: " + passport.iyr + " eyr: " + passport.eyr + " hgt: " + passport.hcl + " hcl: " + passport.hgt + " ecl: " + passport.ecl + " pid: " + passport.pid + " cid: " + passport.cid + " valid: " + strconv.FormatBool(passport.valid))
}
