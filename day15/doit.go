package main

import (
	"fmt"
	"log"
)

func main() {

	//
	//data := []int{0,3,6,0}
	//target := 0
	//result := 3
	//first, second := searchForLastTime(data, target)
	//if first != 1 && second != 4{
	//	panic("fail")
	//}
	//
	//
	//data = []int{0,3,6,0,3}
	//target = 3
	//result = 3
	//first, second = searchForLastTime(data, target)
	//if second - first != result{
	//	panic("fail")
	//}
	//
	//
	//data = []int{1,2,3}
	//target = 1
	//result = 3
	//first, second = searchForLastTime(data, target)
	//if first != result && second != 0{
	//	panic("fail")
	//}
	//
	//target = 2
	//result = 2
	//first, second = searchForLastTime(data, target)
	//if first != result && second != 0{
	//	panic("fail")
	//}
	//
	//
	//target = 3
	//result = 1
	//first, second = searchForLastTime(data, target)
	//if first != result && second != 0{
	//	panic("fail")
	//}
	//
	//target = 4
	//result = 0
	//first, second = searchForLastTime(data, target)
	//if first != result && second != 0{
	//	panic("fail")
	//}
	//
	//
	//if part1([]int{0,3,6}, 2020) != 436{
	//	panic("Test Failed")
	//}
	//
	//if part1([]int{1,3,2}, 2020) != 1{
	//	panic("Test Failed")
	//}
	//if part1([]int{2,1,3}, 2020) != 10{
	//	panic("Test Failed")
	//}
	//
	//if part1([]int{1,2,3}, 2020) != 27{
	//	panic("Test Failed")
	//}
	//if part1([]int{2,3,1}, 2020) != 78{
	//	panic("Test Failed")
	//}
	//if part1([]int{3,2,1}, 2020) != 438{
	//	panic("Test Failed")
	//}
	//if part1([]int{3,1,2}, 2020) != 1836{
	//	panic("Test Failed")
	//}

	/*
		puzzle input
	*/
	lastSpoken := part1([]int{5, 1, 9, 18, 13, 8, 0}, 2020)
	log.Println(fmt.Sprintf("Last Spoken Part 1: %d", lastSpoken))

	//if part2(data) != 123{
	//	panic("Test Failed")
	//}

}

func part1(data []int, target int) int {

	lastSpoken := make([]int, 0)

	/*
		push our starting words
	*/
	for i := 0; i < len(data); i++ {
		word := data[i]
		lastSpoken = append(lastSpoken, word)
	}

	for i := 4; i <= target; i++ {
		word := lastSpoken[len(lastSpoken)-1]
		//log.Println(fmt.Sprintf("Last Spoken: %d",  word))

		wordsToSearch := lastSpoken

		first, second := searchForLastTime(wordsToSearch, word)
		lastTimeSpoken := 0
		if second == 0 {
			lastTimeSpoken = 0
		} else {
			lastTimeSpoken = second - first
		}

		lastSpoken = append(lastSpoken, lastTimeSpoken)
	}
	lastWords := lastSpoken[len(lastSpoken)-10:]
	if len(lastWords) > 0 {

	}
	lastWord := lastSpoken[len(lastSpoken)-1]
	return lastWord
}

func searchForLastTime(data []int, target int) (first int, second int) {

	numbers := make([]int, 0)

	for i := 0; i < len(data); i++ {
		word := data[i]
		if word == target {
			numbers = append(numbers, i)
		}
	}

	if len(numbers) > 1 {
		end := len(numbers) - 1
		first = numbers[end-1] + 1
		second = numbers[end] + 1
	}

	return first, second
}
