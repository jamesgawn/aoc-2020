package day4

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr     string
	iyr     string
	eyr     string
	hgt     string
	hcl     string
	ecl     string
	pid     string
	cid     string
	validV1 bool
	validV2 bool
}

func ExecuteSolution(input io.Reader) {
	passports := parsePassport(input)
	validatePassportV1(passports)
	validatePassportV2(passports)
	validPassportsV1 := 0
	validPassportsV2 := 0
	for _, validatedPassport := range passports {
		if validatedPassport.validV1 {
			validPassportsV1++
		}
		if validatedPassport.validV2 {
			validPassportsV2++
		}
		passportPrinter(validatedPassport)
	}
	log.Println("Total Valid Passports (V1): " + strconv.Itoa(validPassportsV1))
	log.Println("Total Valid Passports (V2): " + strconv.Itoa(validPassportsV2))
}

func parsePassport(input io.Reader) []passport {
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
				switch key {
				case "byr":
					currentPassport.byr = value
				case "iyr":
					currentPassport.iyr = value
				case "eyr":
					currentPassport.eyr = value
				case "hgt":
					currentPassport.hgt = value
				case "hcl":
					currentPassport.hcl = value
				case "ecl":
					currentPassport.ecl = value
				case "pid":
					currentPassport.pid = value
				case "cid":
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

func validatePassportV1(passports []passport) {
	for index, passport := range passports {
		passports[index].validV1 = passport.byr != "" &&
			passport.iyr != "" &&
			passport.eyr != "" &&
			passport.hgt != "" &&
			passport.hcl != "" &&
			passport.ecl != "" &&
			passport.pid != ""
	}
}

func validatePassportV2(passports []passport) {
	for index, passport := range passports {
		passports[index].validV2 = ValidatePassportBirthYear(passport) &&
			ValidatePassportIssueYear(passport) &&
			ValidatePassportExpiryYear(passport) &&
			ValidatePassportHeight(passport) &&
			ValidatePassportHairColour(passport) &&
			ValidatePassportEyeColour(passport) &&
			ValidatePassportNumber(passport)
	}
}

//byr (Birth Year) - four digits; at least 1920 and at most 2002.
func ValidatePassportBirthYear(passport passport) bool {
	return validateDate(passport.byr, 1920, 2002)
}

//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func ValidatePassportIssueYear(passport passport) bool {
	return validateDate(passport.iyr, 2010, 2020)
}

//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func ValidatePassportExpiryYear(passport passport) bool {
	return validateDate(passport.eyr, 2020, 2030)
}

func validateDate(date string, earliest int, latest int) bool {
	year, err := strconv.Atoi(date)
	if err != nil {
		return false
	} else {
		return year >= earliest && year <= latest
	}
}

/* hgt (Height) - a number followed by either cm or in:
 * If cm, the number must be at least 150 and at most 193.
 * If in, the number must be at least 59 and at most 76.
 */
func ValidatePassportHeight(passport passport) bool {
	rawHeight := passport.hgt
	if len(rawHeight) < 2 {
		return false
	}
	rawScale := rawHeight[len(rawHeight)-2:]
	rawValue, err := strconv.Atoi(rawHeight[0 : len(rawHeight)-2])
	if err != nil {
		return false
	}
	if rawScale == "cm" {
		return rawValue >= 150 && rawValue <= 193
	}
	if rawScale == "in" {
		return rawValue >= 59 && rawValue <= 76
	}
	return false
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func ValidatePassportHairColour(passport passport) bool {
	matched, _ := regexp.Match("^#[a-f0-9]{6}$", []byte(passport.hcl))
	return matched
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func ValidatePassportEyeColour(passport passport) bool {
	matched, _ := regexp.Match("^amb|blu|brn|gry|grn|hzl|oth$", []byte(passport.ecl))
	return matched
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func ValidatePassportNumber(passport passport) bool {
	matched, _ := regexp.Match("^[0-9]{9}$", []byte(passport.pid))
	return matched
}

func passportPrinter(passport passport) {
	log.Println("byr: " + passport.byr + " iyr: " + passport.iyr + " eyr: " + passport.eyr + " hgt: " + passport.hcl + " hcl: " + passport.hgt + " ecl: " + passport.ecl + " pid: " + passport.pid + " cid: " + passport.cid + " validV1: " + strconv.FormatBool(passport.validV1) + " validV2: " + strconv.FormatBool(passport.validV2))
}
