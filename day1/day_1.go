package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data/day_1_data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []uint

	for fileScanner.Scan() {
		aNumber, err := strconv.ParseUint(fileScanner.Text(), 10, 0)
		if err != nil {
			panic(err)
		}
		fileTextLines = append(fileTextLines, uint(aNumber))
	}

	file.Close()

	//for _, eachline := range fileTextLines {
	//	fmt.Println(eachline)
	//}

	var target uint
	target = 2020
	var number1, number2 uint
	number1, number2, err = addTwoNumbersToGet(target, fileTextLines, fileTextLines)
	log.Println(fmt.Printf("Number1: %d Number2: %d.... %d..... Answer: %d", number1, number2, (number1 + number2), (number1 * number2)))

	var number3 uint
	number1, number2, number3, err = addThreeNumbersToGet(target, fileTextLines, fileTextLines)
	log.Println(fmt.Printf("Number1: %d Number2: %d....Number3: %d.... %d..... Answer: %d", number1, number2, number3, (number1 + number2 + number3), (number1 * number2 * number3)))

}

func addTwoNumbersToGet(target uint, a []uint, b []uint) (uint, uint, error) {

	for _, number1 := range a {

		for _, number2 := range b {
			candidate := number1 + number2
			if candidate == target {
				return number1, number2, nil
			}
		}
	}
	return 0, 0, nil
}

func addThreeNumbersToGet(target uint, a []uint, b []uint) (uint, uint, uint, error) {

	for _, number1 := range a {

		for _, number2 := range b {
			for _, number3 := range b {
				candidate := number1 + number2 + number3
				if candidate == target {
					return number1, number2, number3, nil
				}
			}
		}
	}
	return 0, 0, 0, nil
}
