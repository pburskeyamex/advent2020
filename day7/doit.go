package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bag struct {
	color          string
	associatedBags []*bagAssociation
}

func createBag(color string) *bag {
	b := bag{
		color:          color,
		associatedBags: make([]*bagAssociation, 0),
	}
	return &b
}

func (b *bag) containsColor(color string, dictionary map[string]*bag) bool {
	itDoes := b.color == color
	if !itDoes {
		for _, association := range b.associatedBags {
			itDoes = association.containsColor(color, dictionary)
			if itDoes {
				break
			} else {
				/*
					check to see if this color as a parent in the broader dictionary contains our subject color
				*/
				potentialCandidateDeepBag := dictionary[association.bag.color]
				if potentialCandidateDeepBag.containsColor(color, dictionary) {
					itDoes = true
					break
				}
			}
		}
	}

	return itDoes
}

type bagAssociation struct {
	quantity int
	bag      *bag
}

func createBagAssociation(quantity int, anAssociatedBag *bag) *bagAssociation {
	b := bagAssociation{
		quantity: quantity,
		bag:      anAssociatedBag,
	}
	return &b
}

func resolveBag(color string, dictionary map[string]*bag) *bag {

	_, ok := dictionary[color]
	if !ok {
		dictionary[color] = createBag(color)
	}
	aBag := dictionary[color]
	return aBag
}

func (b *bagAssociation) containsColor(color string, dictionary map[string]*bag) bool {
	itDoes := b.bag.containsColor(color, dictionary)

	return itDoes
}

func main() {

	data := Parse()
	dictionary := parseDictionary(data)

	bags := find(dictionary, "shiny gold")

	for _, aBag := range bags {
		fmt.Println(fmt.Sprintf("Color: %s ...", aBag.color))
	}
	fmt.Println(fmt.Sprintf("Found: %d bags...", len(bags)-1))

	shinyGold := dictionary["shiny gold"]
	count := howManyBagsInside(shinyGold)
	fmt.Println(fmt.Sprintf("%d bags are inside of Color: %s ...", count, shinyGold.color))

}

func howManyBagsInside(aBag *bag) int {
	count := 0
	if aBag != nil {
		for _, associatedBag := range aBag.associatedBags {
			count += associatedBag.quantity
			count += howManyBagsInside(associatedBag.bag)
		}
	}
	return count
}

//
//func rootColor(color string) string {
//	rootColor := color
//	index := strings.Index(color, " ")
//	if index > -1 {
//		rootColor = color[index+1:]
//	}
//	return rootColor
//}

//func findRootColorMapping(dictionary map[string]bag) (map[string][]string) {
//
//	mapping := make(map[string][]string)
//
//	for _, aContainer := range dictionary {
//		aColor :=rootColor(aContainer.color)
//		_, ok := mapping[aColor]
//		if !ok{
//			//stringArray := []string{}
//			mapping[aColor] = make([]string, 0)
//		}
//
//	}
//	//for key, value := range mapping{
//	//	fmt.Println(key)
//	//	if len(value) > 0{
//	//
//	//	}
//	//}
//
//	//mapping["yellow"] = append(mapping["yellow"], "snot")
//	//
//
//
//	av := func(aMap map[string][]string, rootColor string, synonymColors ... string) {
//
//		for _, aSynonymColor := range synonymColors{
//			found := false
//
//			for _, aColor := range aMap[rootColor]{
//				if aColor == aSynonymColor{
//					found = true
//					break
//				}
//			}
//
//			if !found{
//				aMap[rootColor] = append(aMap[rootColor], aSynonymColor)
//			}
//		}
//
//	}
//
//
//	av(mapping, "yellow", "gold")
//	//av(mapping, "green", "olive", "lime")
//	//av(mapping, "brown", "tan", "bronze", "beige")
//	//av(mapping, "red", "tomato", "crimson", "fuchsia", "maroon")
//	//av(mapping, "blue", "teal", "turquoise", "aqua", "cyan", "indigo")
//	//av(mapping, "purple", "plum", "lavender", "violet")
//	//av(mapping, "white", "")
//	//av(mapping, "black", "")
//	//av(mapping, "orange", "coral", "salmon")
//	//av(mapping, "magenta", "")
//	//av(mapping, "gray", "silver")
//	//av(mapping, "chartreuse", "")
//
//
//
//
//
//
//
//
//
//	return mapping
//}

func find(dictionary map[string]*bag, color string) []*bag {
	answer := make([]*bag, 0)

	for _, aContainer := range dictionary {
		if aContainer.containsColor(color, dictionary) {
			answer = append(answer, aContainer)
		}
	}
	return answer
}

//
//func isColorRelatedTo(color string, rootColor string, mapping map[string][]string) bool {
//	itIs := false
//
//	_, ok := mapping[rootColor]
//	if ok{
//		candidateSynonyms := mapping[rootColor]
//		for _, aCandidateSynonymColor := range candidateSynonyms{
//			if aCandidateSynonymColor == color{
//				itIs = true
//				break
//			}
//		}
//	}
//
//	return itIs
//}

func resolveOrCreateBag(color string, dictionary map[string]*bag) *bag {

	aBag := resolveBag(color, dictionary)
	return aBag

}

func parseDictionary(data []string) map[string]*bag {
	dictionary := make(map[string]*bag, 0)

	for _, aString := range data {

		index := strings.Index(aString, "bags")
		color := aString[0:index]
		color = strings.Trim(color, " ")
		index = strings.Index(aString, "contain")
		contains := aString[index+len("contain"):]

		container := resolveOrCreateBag(color, dictionary)
		//dictionary[container.color] = container
		contains = strings.Trim(contains, " ")
		if contains != "no other bags." {

			parseNumberColorBag := func(aMungeString string) *bagAssociation {
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

				dictionaryAssociation := resolveOrCreateBag(color, dictionary)
				anAssociation := &bagAssociation{
					quantity: quantity,
					bag:      dictionaryAssociation,
				}

				return anAssociation
			}

			if strings.Contains(contains, ",") {
				containsColors := strings.Split(contains, ",")
				for _, aContainsColor := range containsColors {
					aContainerBag := parseNumberColorBag(aContainsColor)
					container.associatedBags = append(container.associatedBags, aContainerBag)
				}
			} else {
				aContainerBag := parseNumberColorBag(contains)
				container.associatedBags = append(container.associatedBags, aContainerBag)
			}
		}
		//log.Println(contains)

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
