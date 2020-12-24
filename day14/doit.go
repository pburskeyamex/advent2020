package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	data := Parse("day_14_data.txt")
	//part1(data)
	part2(data)

}

func part1(data []string) {

	memory := make(map[uint64]uint64, 0)

	var maskAsString []string

	for i := 0; i < len(data); i++ {
		anInstruction := data[i]
		////log.Println(anInstruction)
		//var mask uint64
		//mask |= 11
		//fmt.Printf("Binary: %b", mask)

		if strings.Index(anInstruction, "mask") == 0 {
			maskAsString = parseMaskInstruction(anInstruction)
		} else if strings.Index(anInstruction, "mem") == 0 {
			var address uint64
			var value uint64
			var currentMapValue uint64
			address, value = parseMemInstruction(anInstruction)

			/*
				does address exist
			*/
			currentMapValue, ok := memory[address]
			if !ok {
				memory[address] = currentMapValue
			}
			/*
				we are always overrighting with what is being set....
			*/
			currentMapValue = value
			currentMapValue = applyPart1MaskTo(maskAsString, currentMapValue)

			memory[address] = currentMapValue

		} else {
			panic(fmt.Sprintf("Instruction: %s not understood", anInstruction))
		}

	}

	sum := uint64(0)
	for key, element := range memory {
		fmt.Println(fmt.Sprintf("Address: %d Binary Value: %b Decimal Value %d", key, element, element))
		sum += element
	}
	fmt.Println(fmt.Sprintf("Final Memory Total: %d", sum))

}

func part2(data []string) {

	memory := make(map[uint64]uint64, 0)

	var maskAsString []string

	for i := 0; i < len(data); i++ {
		anInstruction := data[i]
		////log.Println(anInstruction)
		//var mask uint64
		//mask |= 11
		//fmt.Printf("Binary: %b", mask)

		if strings.Index(anInstruction, "mask") == 0 {
			maskAsString = parseMaskInstruction(anInstruction)
		} else if strings.Index(anInstruction, "mem") == 0 {
			var address uint64
			var value uint64
			var currentMapValue uint64
			address, value = parseMemInstruction(anInstruction)

			///*
			//	does address exist
			//*/
			//currentMapValue, ok := memory[address]
			//if !ok {
			//	currentMapValue = address
			//	memory[address] = currentMapValue
			//}
			/*
				we are always assigning the address as the value.....
			*/
			currentMapValue = address

			mungedMemory := applyPart2MaskTo(maskAsString, currentMapValue)
			for key, _ := range mungedMemory {
				memory[key] = value
			}

		} else {
			panic(fmt.Sprintf("Instruction: %s not understood", anInstruction))
		}

	}

	sum := uint64(0)
	for key, element := range memory {
		fmt.Println(fmt.Sprintf("Address: %d Binary Value: %b Decimal Value %d", key, element, element))
		sum += element
	}
	fmt.Println(fmt.Sprintf("Final Memory Total: %d", sum))

}
func applyPart2MaskTo(mask []string, currentMapValue uint64) map[uint64]uint64 {
	var mungedValue uint64
	mungedValue = currentMapValue
	lengthOfMask := len(mask)

	/*
		first phase, apply the mask to get our munged value
	*/

	for i := 0; i < lengthOfMask; i++ {

		position := (lengthOfMask - i) - 1
		/*
			we want to parse in reverse...
		*/
		aMaskValue := mask[position]
		if aMaskValue == "X" {

		} else {
			aBit, _ := strconv.ParseUint(aMaskValue, 10, 64)
			//fmt.Println(fmt.Sprintf("Binary: %b", mungedValue))
			//fmt.Println(fmt.Sprintf("Shifting %d >> %d", aBit, i))

			if aBit == 0 {
				// yummy
			} else if aBit == 1 {
				/*
					write a 1.....
				*/
				mungedValue = mungedValue | (aBit << i)
			}

			//isBitSet := mungedValue & (1 << i) != 0
			//if isBitSet && aBit == 0{
			//	mungedValue = mungedValue ^ (1<<i)
			//} else if (aBit == 1 && !isBitSet){
			//
			//}

			//fmt.Println(fmt.Sprintf("Binary: %b", mungedValue))
		}
	}

	fmt.Println("Applied mask to base value")
	fmt.Println(fmt.Sprintf("Binary: %b", mungedValue))
	fmt.Println(fmt.Sprintf("Mask: %s", mask))
	mungedMemoryAddresses := make(map[uint64]uint64, 0)

	positionsToVary := make([]int, 0)
	for i := 0; i < lengthOfMask; i++ {

		position := (lengthOfMask - i) - 1
		/*
			we want to parse in reverse...
		*/
		aMaskValue := mask[position]
		if aMaskValue == "X" {
			positionsToVary = append(positionsToVary, i)
		}
	}

	proof := math.Pow(2, float64(len(positionsToVary)))

	mungedMemoryAddresses = varyBitCombinations(positionsToVary, mungedValue, mungedMemoryAddresses, int(proof))

	return mungedMemoryAddresses
}

func varyBitCombinations(positionsToVary []int, mungedValue uint64, memory map[uint64]uint64, proof int) map[uint64]uint64 {
	for i := 0; i < len(positionsToVary); i++ {
		position := positionsToVary[i]
		value0, value1 := varyBitsAtPosition(position, mungedValue)
		memory[value0] = value0
		memory[value1] = value1
	}

	if len(memory) < proof {
		for _, value := range memory {
			memory = varyBitCombinations(positionsToVary, value, memory, proof)
		}
	}

	return memory
}

func varyBitsAtPosition(position int, mungedValue uint64) (value0 uint64, value1 uint64) {
	value0 = mungedValue
	value1 = mungedValue
	isBitSet := mungedValue&(1<<position) != 0
	if isBitSet {
		value0 = value0 ^ (1 << position)
	}
	value1 = mungedValue | (1 << position)

	//fmt.Println(fmt.Sprintf("Binary Value 0: %b Base 10: %d", value0, value0))
	//fmt.Println(fmt.Sprintf("Binary Value 1: %b Base 10: %d", value1, value1))

	return value0, value1
}

func applyPart1MaskTo(mask []string, currentMapValue uint64) uint64 {
	var mungedValue uint64
	mungedValue = currentMapValue
	lengthOfMask := len(mask)
	for i := 0; i < lengthOfMask; i++ {

		position := (lengthOfMask - i) - 1
		/*
			we want to parse in reverse...
		*/
		aMaskValue := mask[position]
		if aMaskValue != "X" {
			aBit, _ := strconv.ParseUint(aMaskValue, 10, 64)
			//fmt.Println(fmt.Sprintf("Binary: %b", mungedValue))
			//fmt.Println(fmt.Sprintf("Shifting %d >> %d", aBit, i))

			isBitSet := mungedValue&(1<<i) != 0
			if isBitSet && aBit == 0 {
				mungedValue = mungedValue ^ (1 << i)
			} else if aBit == 1 && !isBitSet {
				mungedValue = mungedValue | (aBit << i)
			}

			//fmt.Println(fmt.Sprintf("Binary: %b", mungedValue))
		}
	}

	return mungedValue
}

func parseMaskInstruction(instruction string) []string {

	runes := []rune(instruction)

	mask := string(runes[7:])
	mask = strings.Trim(mask, " ")

	return strings.Split(mask, "")
}

func parseMemInstruction(instruction string) (address uint64, value uint64) {
	start := strings.Index(instruction, "[")
	end := strings.Index(instruction, "]")

	runes := []rune(instruction)
	aString := string(runes[start+1 : end])
	address, _ = strconv.ParseUint(aString, 10, 64)

	start = strings.Index(instruction, "=")
	aString = string(runes[start+1:])
	aString = strings.Trim(aString, " ")
	value, _ = strconv.ParseUint(aString, 10, 64)

	return
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
