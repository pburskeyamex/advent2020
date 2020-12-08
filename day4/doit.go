package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	passports := Parse()
	validCount := 0
	for _, passport := range passports {
		if passport.validate() {
			validCount++
		}
	}
	log.Printf("Valid passport count: %d", validCount)
}

type passport struct {
	//byr (Birth Year)
	//iyr (Issue Year)
	//eyr (Expiration Year)
	//hgt (Height)
	//hcl (Hair Color)
	//ecl (Eye Color)
	//pid (Passport ID)
	//cid (Country ID)

	birthYear      string
	issueYear      string
	expirationYear string
	height         string
	hairColor      string
	eyeColor       string
	passportID     string
	countryID      string
}

func (p *passport) validateRange(aString string, low int, high int) bool {
	valid := false
	var aNumber int
	aNumber, err := strconv.Atoi(aString)
	if err == nil {
		valid = true
		valid = valid && aNumber >= low
		valid = valid && aNumber <= high
	}
	return valid
}

func (p *passport) validateHeight(aString string) bool {
	valid := false
	//60in
	//190cm
	if strings.Contains(aString, "in") {
		aStringPortion := strings.ReplaceAll(aString, "in", "")
		aNumber, err := strconv.Atoi(aStringPortion)
		if err == nil {

			if aNumber >= 59 && aNumber <= 76 {
				valid = true
			}
		}
	} else if strings.Contains(aString, "cm") {
		aStringPortion := strings.ReplaceAll(aString, "cm", "")
		aNumber, err := strconv.Atoi(aStringPortion)
		if err == nil {

			if aNumber >= 150 && aNumber <= 193 {
				valid = true
			}
		}
	}

	return valid
}

func (p *passport) validate() bool {

	//v := func(aString string) bool {
	//	return len(aString) > 0
	//}
	var valid bool
	valid = true
	valid = valid && p.validateRange(p.birthYear, 1920, 2002)
	valid = valid && p.validateRange(p.issueYear, 2010, 2020)
	valid = valid && p.validateRange(p.expirationYear, 2020, 2030)
	valid = valid && p.validateHeight(p.height)
	valid = valid && p.validateHairColor(p.hairColor)
	valid = valid && p.validateEyeColor(p.eyeColor)
	valid = valid && p.validatePassportID(p.passportID)
	//valid = v(p.countryID)
	return valid
}

func (p *passport) validateEyeColor(aString string) bool {
	valid := false

	pattern := "^(amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth)$"

	valid, err := regexp.MatchString(pattern, aString)
	if err != nil {
		log.Fatal(err)
	}
	return valid
}

func (p *passport) validateHairColor(aString string) bool {
	valid := false
	pattern := "^#[a-fA-F0-9]{6}$"
	valid, err := regexp.MatchString(pattern, aString)
	if err != nil {
		log.Fatal(err)
	}
	return valid
}

func (p *passport) validatePassportID(aString string) bool {
	valid := false
	pattern := "^[0-9]{9}$"
	valid, err := regexp.MatchString(pattern, aString)
	if err != nil {
		log.Fatal(err)
	}
	return valid
}

func (p *passport) parse(aString string) *passport {

	if len(aString) > 0 {
		segments := strings.Split(aString, " ")

		var aSegment string
		for _, aSegment = range segments {
			splits := strings.Split(aSegment, ":")
			key := splits[0]
			value := splits[1]

			if key == "byr" {
				p.birthYear = value
			} else if key == "iyr" {
				p.issueYear = value
			} else if key == "eyr" {
				p.expirationYear = value
			} else if key == "hgt" {
				p.height = value
			} else if key == "hcl" {
				p.hairColor = value
			} else if key == "ecl" {
				p.eyeColor = value
			} else if key == "pid" {
				p.passportID = value
			} else if key == "cid" {
				p.countryID = value
			} else {
				log.Fatal("Key: " + key + " Not understood")
			}

		}
	}

	return p
}

func Parse() []*passport {
	file, err := os.Open("data/day_4_data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	passports := make([]*passport, 0)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var aPassport *passport
	for i := 0; fileScanner.Scan(); i++ {

		aString := fileScanner.Text()

		length := len(passports)
		if length == 0 || len(aString) == 0 {
			aPassport = &passport{
				birthYear:      "",
				issueYear:      "",
				expirationYear: "",
				height:         "",
				hairColor:      "",
				eyeColor:       "",
				passportID:     "",
				countryID:      "",
			}
			passports = append(passports, aPassport)
		} else {
			aPassport = passports[length-1]
		}

		aPassport.parse(aString)

	}

	file.Close()

	return passports
}
