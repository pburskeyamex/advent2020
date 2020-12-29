package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ticket struct {
	data []int
}

type ticketRules struct {
	rules []*rule
}

func (me *ticketRules) build(rules []*rule, sampleTickets []*ticket) {

	columns := len(sampleTickets[0].data)
	me.rules = make([]*rule, columns)

	var mungedRules []map[*rule]int
	mungedRules = make([]map[*rule]int, 0)

	for x := 0; x < columns; x++ {

		sampleData := make([]int, len(sampleTickets))
		for i := 0; i < len(sampleData); i++ {
			aSampleTicket := sampleTickets[i]
			sampleData[i] = aSampleTicket.data[x]
		}

		columnMap := make(map[*rule]int, 0)
		mungedRules = append(mungedRules, columnMap)

		/*
			given some sample data, which rule fits....
		*/
		for i := 0; i < len(sampleData); i++ {
			aSample := sampleData[i]

			for _, aRule := range rules {
				if aRule.valid(aSample) {

					value, _ := columnMap[aRule]
					value++
					columnMap[aRule] = value

				}

			}

		}

	}

	sampleSize := len(sampleTickets)
	reduceFurther := true

	for i := 0; reduceFurther; i++ {
		me.reduce(sampleSize, mungedRules)
		reduceFurther = false
		for _, aRuleMap := range mungedRules {

			if len(aRuleMap) > 1 {
				reduceFurther = true
				break
			}

		}
	}
	me.reduce(sampleSize, mungedRules)

}

func (me *ticketRules) reduce(sampleSize int, mungedRules []map[*rule]int) {

	/*
		first pass.... easy ones
	*/
	for _, aRuleMap := range mungedRules {

		for key, value := range aRuleMap {
			if value == sampleSize {
				/*
					keep this one....
				*/
			} else {
				delete(aRuleMap, key)
			}

		}

	}

	//
	//
	///*
	//second pass, get the the ones that are not already sorted out in phase 1
	// */
	for key, aRuleMap := range mungedRules {
		if len(aRuleMap) == 1 {
			for aRule, _ := range aRuleMap {
				me.rules[key] = aRule
			}
			foundRule := me.rules[key]

			for _, aRuleMapToSearch := range mungedRules {
				for aPotentialRuleToDeleteKey, _ := range aRuleMapToSearch {
					if aPotentialRuleToDeleteKey == foundRule {
						delete(aRuleMapToSearch, aPotentialRuleToDeleteKey)
					}
				}
			}
		}
	}

}

func (me *ticket) valid(rules []*rule) bool {

	for x := 0; x < len(me.data); x++ {
		isValid := false
		aNumber := me.data[x]
		for i := 0; !isValid && i < len(rules); i++ {
			isValid = rules[i].valid(aNumber)
		}

		if !isValid {
			return false
		}

	}
	return true
}

type rule struct {
	name        string
	constraints []*rangeConstraint
}

func (me *rule) valid(aNumber int) bool {
	isValid := false
	for i := 0; !isValid && i < len(me.constraints); i++ {
		isValid = me.constraints[i].valid(aNumber)
	}
	return isValid
}

type rangeConstraint struct {
	low  int
	high int
}

func (me *rangeConstraint) valid(aNumber int) bool {
	isValid := (aNumber >= me.low)
	isValid = (isValid && (aNumber <= me.high))
	return isValid
}

func main() {

	expectation := 71
	if errorRate := part1("day_16_sample_data.txt"); errorRate != expectation {
		panic(fmt.Sprintf("Expected: %d", expectation))
	}
	expectation = 29851
	if errorRate := part1("day_16_data.txt"); errorRate != expectation {
		panic(fmt.Sprintf("Expected: %d", expectation))
	}

	expectation = 3029180675981
	if errorRate := part2("day_16_data.txt"); errorRate != expectation {
		panic(fmt.Sprintf("Expected: %d", expectation))
	}

}

func part1(fileName string) int {
	var rules []*rule
	//var myTickets []int
	var nearByTickets []*ticket
	rules, _, nearByTickets = Parse(fileName)

	errors := make([]int, 0)
	for _, aTicket := range nearByTickets {

		for x := 0; x < len(aTicket.data); x++ {
			isValid := false
			aNumber := aTicket.data[x]
			for i := 0; !isValid && i < len(rules); i++ {
				isValid = rules[i].valid(aNumber)
			}

			if !isValid {
				errors = append(errors, aTicket.data[x])
				break
			}

		}
	}

	errorRate := 0
	for _, aTicket := range errors {
		errorRate += aTicket
	}
	return errorRate
}

func part2(fileName string) int {
	var rules []*rule
	var myTicket *ticket
	var nearByTickets []*ticket
	rules, myTicket, nearByTickets = Parse(fileName)

	var ticketsToConsider []*ticket

	for _, aTicket := range nearByTickets {
		if aTicket.valid(rules) {
			ticketsToConsider = append(ticketsToConsider, aTicket)
		}
	}

	ticketRules := &ticketRules{}
	ticketRules.build(rules, ticketsToConsider)

	/*
		gather rules
	*/
	errorRate := 1
	for index, rule := range ticketRules.rules {
		if strings.Contains(rule.name, "departure") {
			myValue := myTicket.data[index]
			errorRate = errorRate * myValue
		}
	}

	return errorRate
}

func Parse(aFilePart string) (rules []*rule, myTicket *ticket, nearByTickets []*ticket) {
	filename := fmt.Sprintf("data/%s", aFilePart)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rules = make([]*rule, 0)
	myTicket = &ticket{}
	nearByTickets = make([]*ticket, 0)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	parsePosition := 1
	for i := 0; fileScanner.Scan(); i++ {
		aString := fileScanner.Text()
		/*
			class: 1-3 or 5-7
		*/

		spaceCount := strings.Count(aString, " ")
		if spaceCount >= 3 {
			runes := []rune(aString)
			index := strings.Index(aString, ":")
			name := string(runes[:index])

			rule := &rule{
				name:        name,
				constraints: make([]*rangeConstraint, 0),
			}
			rules = append(rules, rule)

			remainingString := string(runes[index+1:])
			splits := strings.Split(remainingString, " or ")

			for _, aString := range splits {
				aString = strings.Trim(aString, " ")
				splits = strings.Split(aString, "-")
				low, _ := strconv.Atoi(splits[0])
				high, _ := strconv.Atoi(splits[1])
				aConstraint := &rangeConstraint{
					low:  low,
					high: high,
				}
				rule.constraints = append(rule.constraints, aConstraint)
			}

		} else if aString == "" {

		} else if aString == "your ticket:" {
			parsePosition++
		} else if aString == "nearby tickets:" {
			parsePosition++
		} else {
			if parsePosition == 2 {

				splits := strings.Split(aString, ",")
				for _, aString := range splits {
					aNumber, _ := strconv.Atoi(aString)
					myTicket.data = append(myTicket.data, aNumber)

				}

			} else if parsePosition == 3 {
				aTicket := &ticket{data: make([]int, 0)}
				nearByTickets = append(nearByTickets, aTicket)
				splits := strings.Split(aString, ",")
				for _, aString := range splits {
					aNumber, _ := strconv.Atoi(aString)
					aTicket.data = append(aTicket.data, aNumber)

				}
			}
		}

	}

	file.Close()

	return rules, myTicket, nearByTickets
}
