package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	data := Parse("day_10_sample_data.txt")

	var deviceJoltageRating int
	var deltas []int

	deltas = make([]int, 5)
	numbers := make([]int, len(data))
	for index := 0; index < len(data); index++ {
		numbers[index], _ = strconv.Atoi(data[index])
	}

	sort.Slice(numbers, func(i int, j int) bool {
		return numbers[i] < numbers[j]
	})

	deviceJoltageRating = numbers[len(numbers)-1]
	deviceJoltageRating = deviceJoltageRating + 3
	numbers = append(numbers, deviceJoltageRating)

	last := 0
	for x := 0; x < len(numbers); x++ {
		current := numbers[x]
		diff := current - last
		deltas[diff]++
		last = current
		//log.Println(fmt.Sprintf("Index: %d Last: %d Current: %d Difference: %d", x, last, current, diff))
	}
	log.Println(deltas)
	combinations := make([]int, 0)
	combinations = append(combinations, 0)
	combinations = append(combinations, numbers...)

	log.Println(combinations)

	paths := make(map[int]int)
	paths[0] = 1
	for i := 1; i < len(combinations); i++ {
		sum := 0
		for j := i - 1; j >= 0 && combinations[i]-combinations[j] <= 3; j-- {
			sum += paths[combinations[j]]
		}
		paths[combinations[i]] = sum
	}
	log.Println(paths)
}

func combine(combinations []int) {

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
