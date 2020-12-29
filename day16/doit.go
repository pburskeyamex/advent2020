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
	//if part1([]int64{0,3,6}, 2020) != 436{
	//	panic("Test Failed")
	//}
	//
	//if part1([]int64{1,3,2}, 2020) != 1{
	//	panic("Test Failed")
	//}
	//if part1([]int64{2,1,3}, 2020) != 10{
	//	panic("Test Failed")
	//}
	//
	//if part1([]int64{1,2,3}, 2020) != 27{
	//	panic("Test Failed")
	//}
	//if part1([]int64{2,3,1}, 2020) != 78{
	//	panic("Test Failed")
	//}
	//if part1([]int64{3,2,1}, 2020) != 438{
	//	panic("Test Failed")
	//}
	//if part1([]int64{3,1,2}, 2020) != 1836{
	//	panic("Test Failed")
	//}
	//
	///*
	//	puzzle input
	//*/
	//lastSpoken := part1([]int64{5, 1, 9, 18, 13, 8, 0}, 2020)
	//if lastSpoken != 376{
	//	panic("Part 1 failed")
	//}
	//log.Println(fmt.Sprintf("Last Spoken Part 1: %d", lastSpoken))

	if part2([]int{0, 3, 6}, 30000000) != 175594 {
		panic("Test Failed")
	}

	if part2([]int{1, 3, 2}, 30000000) != 2578 {
		panic("Test Failed")
	}
	if part2([]int{2, 1, 3}, 30000000) != 3544142 {
		panic("Test Failed")
	}

	if part2([]int{1, 2, 3}, 30000000) != 261214 {
		panic("Test Failed")
	}
	if part2([]int{2, 3, 1}, 30000000) != 6895259 {
		panic("Test Failed")
	}
	if part2([]int{3, 2, 1}, 30000000) != 18 {
		panic("Test Failed")
	}
	if part2([]int{3, 1, 2}, 30000000) != 362 {
		panic("Test Failed")
	}

	lastSpoken := part2([]int{5, 1, 9, 18, 13, 8, 0}, 30000000)
	//if lastSpoken != 376{
	//	panic("Part 1 failed")
	//}
	log.Println(fmt.Sprintf("Last Spoken Part 2: %d", lastSpoken))

}

func part1(data []int64, target int64) int64 {

	lastSpoken := make([]int64, target)

	/*
		push our starting words
	*/
	for i := 0; i < len(data); i++ {
		word := data[i]
		//lastSpoken = append(lastSpoken, word)
		lastSpoken[i] = word
	}

	for i := int64(len(data)); i < target; i++ {
		word := lastSpoken[len(lastSpoken)-1]
		//log.Println(fmt.Sprintf("Last Spoken: %d",  word))

		wordsToSearch := lastSpoken
		var first, second int64
		first, second = searchForLastTime(wordsToSearch, word)
		lastTimeSpoken := int64(0)
		if second == 0 {
			lastTimeSpoken = 0
		} else {
			lastTimeSpoken = (second) - (first)
		}

		lastSpoken[i] = lastTimeSpoken
		//lastSpoken = append(lastSpoken, lastTimeSpoken)
	}
	lastWords := lastSpoken[len(lastSpoken)-10:]
	if len(lastWords) > 0 {

	}
	lastWord := lastSpoken[len(lastSpoken)-1]
	return lastWord
}

func part2(data []int, target int) int {

	spoken := make([]int, target)
	turn := 1

	for _, input := range data[:len(data)-1] {
		spoken[input] = turn
		turn++
	}

	var prev int
	speak := data[len(data)-1]

	for ; turn <= target; turn++ {
		prev = speak

		if t := spoken[speak]; t != 0 {
			speak = turn - t
		} else {
			speak = 0
		}

		spoken[prev] = turn
	}

	return prev
}

func searchForLastTime(data []int64, target int64) (first int64, second int64) {

	numbers := make([]int64, 0)

	//for i := int64(0); i < int64(len(data)); i++ {
	//	word := data[i]
	//	if word == target {
	//		numbers = append(numbers, i)
	//		if len(numbers) ==2{
	//			break
	//		}
	//	}
	//}

	for i := len(data) - 1; i >= 0; i-- {
		word := data[i]
		if word == target {
			numbers = append(numbers, int64(i))
			if len(numbers) == 2 {
				break
			}
		}
	}

	if len(numbers) > 1 {
		end := len(numbers) - 1
		second = numbers[end-1] + 1
		first = numbers[end] + 1
	}

	//if abs(first - second) > 100000{
	//	log.Println(fmt.Sprintf("First: %d Second: %d", first, second))
	//}

	return first, second
}
func abs(x int64) int64 {
	if x < 0 {
		return (x * -1)
	}
	return x
}
