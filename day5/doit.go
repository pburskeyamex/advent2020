package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	sample := "BBFFBBFRLL"
	rowNumber, seatNumber, seatID := resolveBoardingPass(sample)
	log.Printf("Boarding Pass: %s .... Row Number: %d Seat Number: %d Seat ID: %d\n", sample, rowNumber, seatNumber, seatID)

	highestSeatId := -1
	data := Parse()
	seatIDs := make([]int, 0)
	for _, aString := range data {
		rowNumber, seatNumber, seatID := resolveBoardingPass(aString)
		log.Printf("Boarding Pass: %s .... Row Number: %d Seat Number: %d Seat ID: %d\n", sample, rowNumber, seatNumber, seatID)
		if seatID > highestSeatId {
			highestSeatId = seatID
		}
		seatIDs = append(seatIDs, seatID)
	}

	log.Printf("Highest Seat ID: %d", highestSeatId)

	sort.Slice(seatIDs, func(i int, j int) bool {
		return seatIDs[i] < seatIDs[j]
	})

	for i := 0; i < len(seatIDs); i++ {
		value := seatIDs[i]

		message := fmt.Sprintf("Considering seat: %d", value)
		if (i + 1) < len(seatIDs) {
			nextValue := seatIDs[i+1]

			if nextValue-value != 1 {
				message = message + fmt.Sprintf("**********")
				log.Println(fmt.Sprintf("Missing seat: %d", value+1))
				break
			}
		}
		//log.Println(message)
	}

}

func Parse() []string {
	file, err := os.Open("data/day_5_data.txt")
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

func resolveBoardingPass(aString string) (rowNumber int, seatNumber int, seatID int) {
	rowNumber = reduceToRow(aString)
	seatNumber = reduceToSeat(aString)
	seatID = (rowNumber * 8) + seatNumber

	return rowNumber, seatNumber, seatID
}

func reduceToSeat(aString string) int {

	var seatNumber int
	seatNumber = -1
	seatSlice := strings.Split(aString, "")[7:]
	numberSlice := buildInts(8)
	for ok := true; ok && (seatNumber < 0); {

		seatSlice, numberSlice = popSeat(seatSlice, numberSlice)
		if len(seatSlice) == 0 {
			seatNumber = numberSlice[0]
		}
	}
	return seatNumber
}

func reduceToRow(aString string) int {

	var rowNumber int
	rowNumber = -1
	rowSlice := strings.Split(aString, "")[:7]
	numberSlice := buildInts(128)
	for ok := true; ok && (rowNumber < 0); {

		rowSlice, numberSlice = popRow(rowSlice, numberSlice)
		if len(rowSlice) == 0 {
			rowNumber = numberSlice[0]
		}
	}
	return rowNumber
}

func popRow(stringSlice []string, numberSlice []int) ([]string, []int) {

	bottom, top := splitSliceIntoTwo(numberSlice)
	var remainingSlice []int
	aDirection := stringSlice[0]
	if aDirection == "F" {
		remainingSlice = bottom
	} else if aDirection == "B" {
		remainingSlice = top
	}

	return stringSlice[1:], remainingSlice

}

func popSeat(stringSlice []string, numberSlice []int) ([]string, []int) {

	bottom, top := splitSliceIntoTwo(numberSlice)
	var remainingSlice []int
	aDirection := stringSlice[0]
	if aDirection == "L" {
		remainingSlice = bottom
	} else if aDirection == "R" {
		remainingSlice = top
	}

	return stringSlice[1:], remainingSlice

}

func splitSliceIntoTwo(numberSlice []int) ([]int, []int) {
	length := len(numberSlice)
	a := (length / 2)

	bottom := numberSlice[:a]
	top := numberSlice[a:]
	return bottom, top
}

func buildInts(size int) []int {
	anArray := make([]int, size)
	for i := 0; i < size; i++ {
		anArray[i] = i
	}
	return anArray
}
