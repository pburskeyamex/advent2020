package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type bag struct {
	color    string
	quantity int
	bags     []*bag
}

func main() {

	data := Parse()
	dictionary := parseDictionary(data)

	bags := find(dictionary, rootColor("shiny gold"))
	log.Println(bags)

}

func rootColor(color string) string {
	rootColor := color
	index := strings.Index(color, " ")
	if index > -1 {
		rootColor = color[index+1:]
	}
	return rootColor
}

func find(dictionary map[string]bag, color string) []bag {
	answer := make([]bag, 0)

	for _, aContainer := range dictionary {
		for _, aBag := range aContainer.bags {
			if rootColor(aBag.color) == color {
				answer = append(answer, aContainer)
			}
		}
	}
	return answer
}

func parseDictionary(data []string) map[string]bag {
	dictionary := make(map[string]bag, 0)

	for _, aString := range data {

		index := strings.Index(aString, "bags")
		color := aString[0:index]
		color = strings.Trim(color, " ")
		index = strings.Index(aString, "contain")
		contains := aString[index+len("contain"):]
		container := &bag{
			color:    color,
			quantity: 1,
			bags:     make([]*bag, 0),
		}
		contains = strings.Trim(contains, " ")
		if contains != "no other bags." {

			parseNumberColorBag := func(aMungeString string) *bag {
				aMungeString = strings.Trim(aMungeString, " ")
				if aMungeString == "no other bags." {
					return nil
				}
				aMungeString = strings.ReplaceAll(aMungeString, "bags.", "")
				aMungeString = strings.ReplaceAll(aMungeString, "bags", "")
				aMungeString = strings.ReplaceAll(aMungeString, "bag.", "")
				aMungeString = strings.ReplaceAll(aMungeString, "bag", "")
				aMungeString = strings.Trim(aMungeString, " ")
				index := strings.Index(aMungeString, " ")
				color := aMungeString[index:]
				aString := aMungeString[:index]
				quantity, _ := strconv.Atoi(aString)
				color = strings.Trim(color, " ")

				aBag := &bag{
					color:    color,
					quantity: quantity,
					bags:     nil,
				}

				return aBag
			}

			if strings.Contains(contains, ",") {
				containsColors := strings.Split(contains, ",")
				for _, aContainsColor := range containsColors {
					aContainerBag := parseNumberColorBag(aContainsColor)
					container.bags = append(container.bags, aContainerBag)
				}
			} else {
				aContainerBag := parseNumberColorBag(contains)
				container.bags = append(container.bags, aContainerBag)
			}
		}
		log.Println(contains)

		dictionary[container.color] = *container

	}
	return dictionary
}

func Parse() []string {
	file, err := os.Open("data/day_7_sample_data.txt")
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
