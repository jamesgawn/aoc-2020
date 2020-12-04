package day4

import (
	"aoc2020/pkg/test"
	"testing"
)

func TestValidatePassportBirthYearWithValidInput(t *testing.T) {
	passport := passport{byr: "1950"}
	result := ValidatePassportBirthYear(passport)
	test.EqualsBool(t, true, result)
}

func TestValidatePassportBirthYearWithNaN(t *testing.T) {
	passport := passport{byr: "aasf50"}
	result := ValidatePassportBirthYear(passport)
	test.EqualsBool(t, false, result)
}

func TestValidatePassportBirthYearWithTooOld(t *testing.T) {
	passport := passport{byr: "1910"}
	result := ValidatePassportBirthYear(passport)
	test.EqualsBool(t, false, result)
}

func TestValidatePassportBirthYearWithTooYoung(t *testing.T) {
	passport := passport{byr: "2021"}
	result := ValidatePassportBirthYear(passport)
	test.EqualsBool(t, false, result)
}

func TestValidatePassportExpiryYearWithValidInput(t *testing.T) {
	passport := passport{eyr: "2030"}
	result := ValidatePassportExpiryYear(passport)
	test.EqualsBool(t, true, result)
}

func TestValidatePassportHeightWithValidInput(t *testing.T) {
	passport := passport{hgt: "70in"}
	result := ValidatePassportHeight(passport)
	test.EqualsBool(t, true, result)
}

func TestValidatePassportHairColourWithValidInput(t *testing.T) {
	passport := passport{hcl: "#f4ab21"}
	result := ValidatePassportHairColour(passport)
	test.EqualsBool(t, true, result)
}

func TestValidatePassportEyeColourWithValidInput(t *testing.T) {
	passport := passport{ecl: "amb"}
	result := ValidatePassportEyeColour(passport)
	test.EqualsBool(t, true, result)
}

func TestValidatePassportNumberWithValidInput(t *testing.T) {
	passport := passport{pid: "000000000"}
	result := ValidatePassportNumber(passport)
	test.EqualsBool(t, true, result)
}
