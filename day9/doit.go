package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strconv"
)

func pointerName(aFunction interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(aFunction).Pointer()).Name()
}

func humanOperationID(anOperationID int) int {
	return anOperationID + 1
}

func main() {

	data := Parse("day_9_data.txt")

	numbers := make([]int, len(data))
	for index := 0; index < len(data); index++ {
		numbers[index], _ = strconv.Atoi(data[index])
	}

	preamble := 25
	var found bool
	var offset int
	found = true
	var a, b, target int
	for offset = 0; found && offset < len(numbers); offset++ {
		target = numbers[preamble+offset]
		found, a, b = findCombination(numbers, preamble, offset, target)
	}
	if found {
		log.Println(fmt.Sprintf("X: %d Y:%d.... Sum: %d Target: %d", a, b, a+b, target))
	} else {
		var magicNumber int
		log.Println(fmt.Sprintf("Unable to match for Target: %d", target))
		found, numbers, magicNumber = findCombinationAlternative(numbers, preamble, 0, target)
		log.Println(fmt.Sprintf("Encryption Weakness: %d", magicNumber))
	}

}

func findCombination(numbers []int, preable int, offset int, target int) (bool, int, int) {
	var found bool
	var a, b int

	aSliceOfNumbers := numbers[offset : preable+offset]

	for x := 0; !found && x < len(aSliceOfNumbers); x++ {
		for y := len(aSliceOfNumbers) - 1; !found && y >= 0; y-- {
			a = aSliceOfNumbers[x]
			b = aSliceOfNumbers[y]
			total := a + b
			log.Println(fmt.Sprintf("X: %d Y:%d.... Sum: %d Target: %d", a, b, total, target))

			if total == target {
				found = true
			}
		}
	}
	return found, a, b
}

func findCombinationAlternative(numbers []int, preable int, offset int, target int) (bool, []int, int) {
	var found bool
	var magicNumber int
	var foundNumbers []int
	aSliceOfNumbers := numbers[offset:]

	bucket := 0
	for x := 0; !found && x < len(aSliceOfNumbers); x++ {

		a := aSliceOfNumbers[x]
		bucket += a

		if bucket == target {
			found = true
			foundNumbers = aSliceOfNumbers[:x+1]
			sort.Slice(foundNumbers, func(i int, j int) bool {
				return foundNumbers[i] < foundNumbers[j]
			})
			first := foundNumbers[0]
			last := foundNumbers[len(foundNumbers)-1]
			magicNumber = first + last
			break

		} else if bucket > target {
			return findCombinationAlternative(numbers, preable, offset+1, target)
		}
	}

	return found, numbers, magicNumber
}

func Parse(aFilePart string) []string {
	filename := fmt.Sprintf("data/%s", aFilePart)
	file, err := os.Open(filename)
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
