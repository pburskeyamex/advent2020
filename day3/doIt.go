package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)
func main() {
	var treeCount, missCount uint
	var right, down int
	right, down = 1,2
	treeCount, missCount = Algorithm(right,down, 10)
	fmt.Printf("Tree count: %d Miss count: %d", treeCount, missCount)
}

type traversal struct{
	right int
	down int
}

func determineNSlopes(n int, right int, down int) []traversal{
	traversals := make([]traversal, 0)



	for i:= 0; i < n; i++{

		aTraversal := &traversal{
			right: right,
			down:  down,
		}

		if i > 0 {
			lastTraversal := traversals[i -1]
			adjustedRight := aTraversal.right + lastTraversal.right
			adjustedDown := aTraversal.down + lastTraversal.down
			aTraversal.right = adjustedRight
			aTraversal.down = adjustedDown
		}
		traversals = append(traversals, *aTraversal)
	}
	return traversals




}



func Algorithm( rightCount int, downCount int, printNLines int ) (treeCount uint, missCount uint){
	file, err := os.Open("data/day_3_data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	treeCount = 0
	missCount = 0

	treeSymbol := "X"
	missSymbol := "O"

	traversals := determineNSlopes(350, rightCount, downCount)
	if traversals != nil{

	}
	currentTraversal := 0

	for i := 0; fileScanner.Scan(); i++{

		aTraversal := traversals[currentTraversal]
		aString := fileScanner.Text()
		aString = replicateToLength(aString, aTraversal.right)
		mungedMap := aString

		var currentIndex int
		if i == aTraversal.down{
			currentTraversal++
			currentIndex = aTraversal.right

			stringArray := strings.Split(aString, "")
			if stringArray[currentIndex] == "#"{
				treeCount++
				stringArray[currentIndex] = treeSymbol
			}else{
				missCount++
				stringArray[currentIndex] = missSymbol
			}
			mungedMap = strings.Join(stringArray,"")
			if mungedMap == ""{
				log.Fatalln("arggggggg")
			}


		}



		//log.Println(fmt.Sprintf("Line: %d Right: %d Original: %s......... Map: %s", currentLine, currentIndex, aString, mungedMap))
		if printNLines > 0{
			printNLines--
			log.Println(fmt.Sprintf("Line: %d Right: %d Map: %s", i, currentIndex, mungedMap))
		}




	}




	file.Close()
	return treeCount, missCount
}


func replicateToLength(aString string, aLength int) string{
	if len(aString) > aLength{
		return aString
	}else{
		for len(aString) <= aLength{
			aString = aString + aString
		}
		return aString
	}
}
