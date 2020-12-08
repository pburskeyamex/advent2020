package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {

	groups := make([]map[string]int, 0)

	data := Parse()
	for _, aString := range data {

		if len(aString) == 0 || len(groups) == 0 {
			aMap := make(map[string]int)
			groups = append(groups, aMap)
			aMap["PEOPLE"] = 0
		}

		if len(aString) > 0 {
			sample := strings.Split(aString, "")
			aMap := groups[len(groups)-1]
			parseIntoMap(sample, aMap)

			aMap["PEOPLE"] = aMap["PEOPLE"] + 1
		}
	}
	revisedTotal := 0
	originalTotal := 0
	for i := 0; i < len(groups); i++ {
		aMap := groups[i]
		count := len(aMap) - 1
		revisedCount := 0
		originalTotal += count

		peopleCount := aMap["PEOPLE"]
		if peopleCount == 1 {
			revisedCount += count
		} else {
			for key, value := range aMap {
				if key != "PEOPLE" && (value == peopleCount) {
					revisedCount++
				}
			}

		}

		revisedTotal += revisedCount
		log.Printf("Group: %d Count: %d Revised Count: %d Original Total: %d Revised Total: %d\n", i, count, revisedCount, originalTotal, revisedTotal)
	}

}

func parseIntoMap(strings []string, aMap map[string]int) {
	for _, aString := range strings {
		value := aMap[aString]
		aMap[aString] = value + 1
	}
}

func Parse() []string {
	file, err := os.Open("data/day_6_data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := make([]string, 0)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for i := 0; fileScanner.Scan(); i++ {

		aString := fileScanner.Text()
		data = append(data, aString)

	}

	file.Close()

	return data
}
