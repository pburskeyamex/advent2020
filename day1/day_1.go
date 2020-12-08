package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data/day_1_data.txt")
	if err != nil{
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []uint

	for fileScanner.Scan() {
		aNumber , err := strconv.ParseUint(fileScanner.Text(), 10, 0)
		if err != nil{
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
	fmt.Printf("Number1: %d Number2: %d.... %d..... Answer: %d", number1, number2, (number1 + number2), (number1*number2))
}

func addTwoNumbersToGet(target uint, a []uint, b []uint) (uint, uint, error) {

	for _, number1 :=  range a{

		for _, number2 :=  range b{
			candidate := number1 + number2
			if candidate == target{
				return number1, number2, nil
			}
		}
	}
	return 0, 0, nil
}
