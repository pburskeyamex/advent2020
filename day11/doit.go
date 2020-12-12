package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func deepCopy(data [][]string) [][]string {
	copyData := make([][]string, 0)
	for i, _ := range data {
		copyRow := make([]string, len(data[i]))
		copyData = append(copyData, copyRow)
		for j, aString := range data[i] {
			copyData[i][j] = aString
		}
	}
	return copyData
}

func main() {

	data := Parse("day_11_data.txt")

	var realData [][]string
	realData = make([][]string, 0)

	for _, aString := range data {
		realData = append(realData, strings.Split(aString, ""))
	}

	//var picture [][]string
	//var results []string
	//var coordinates [][]int
	//x := 0
	//y := 0
	sight := 1
	//coordinates = diagonalToDimensions(x, y, sight)
	//picture, results = populateAdjacentSeatsInGraph(sight, coordinates, realData, x, y)
	//prettyPrint(picture)
	//log.Print(results)

	changing := true
	available, occupied := count(realData)
	dataToConsider := realData
	for i := 0; changing; i++ {
		_, _, emptySeatPhaseData := adjustSeating(sight, dataToConsider)
		//prettyPrint(emptySeatPhaseData)
		thisAvailable, thisOccupied := count(emptySeatPhaseData)
		if thisAvailable == available && thisOccupied == occupied {
			changing = false
		} else {
			available = thisAvailable
			occupied = thisOccupied
		}

		dataToConsider = emptySeatPhaseData
		log.Println(fmt.Sprintf("Iteration: %d Available: %d Occupied: %d", i, available, occupied))
	}

}

func adjustSeating(sight int, originalData [][]string) ([][]string, [][]string, [][]string) {
	occupyPhaseData := deepCopy(originalData)
	for i := 0; i < len(originalData); i++ {
		for j := 0; j < len(originalData[i]); j++ {
			if willASeatBecomeFilled(sight, i, j, originalData) {
				occupySeat(i, j, occupyPhaseData)
			}
		}
	}

	//prettyPrint(occupyPhaseData)
	emptySeatPhaseData := deepCopy(occupyPhaseData)

	//log.Println("Adjusting seats")

	for i := 0; i < len(occupyPhaseData); i++ {
		for j := 0; j < len(occupyPhaseData[i]); j++ {
			if willASeatBecomeVacant(sight, i, j, occupyPhaseData) {
				emptySeat(i, j, emptySeatPhaseData)
			}
		}

	}

	return originalData, occupyPhaseData, emptySeatPhaseData
}

func count(data [][]string) (available int, occupied int) {

	for x := 0; x < len(data); x++ {
		for y := 0; y < len(data[x]); y++ {
			seat := data[x][y]
			if isAvailable(seat) {
				available++
			} else if isOccupied(seat) {
				occupied++
			}
		}
	}

	return available, occupied
}

func populateAdjacentSeatsInGraph(sight int, coordinates [][]int, data [][]string, interestingX int, interestingY int) ([][]string, []string) {

	var strings []string
	var picture [][]string
	strings = make([]string, 0)

	widthOfSightFloor := 0
	widthOfSightCeiling := 0
	for i := 0; i < len(coordinates); i++ {
		for j := 0; j < len(coordinates[i]); j++ {
			value := coordinates[i][j]
			if value < widthOfSightFloor {
				widthOfSightFloor = value
			}
			if value > widthOfSightCeiling {
				widthOfSightCeiling = value
			}
		}
	}

	actualSight := (widthOfSightCeiling - widthOfSightFloor) + 1

	picture = make([][]string, actualSight)
	for i := 0; i < len(picture); i++ {
		picture[i] = make([]string, actualSight)
	}
	pictureX := 0
	pictureY := 0
	for i := 0; i < len(coordinates); i++ {
		aPositionArray := coordinates[i]
		x := aPositionArray[0]
		y := aPositionArray[1]

		aString := " "
		//log.Println(fmt.Sprintf("X %d  Y %d" ,x, y))
		if x == interestingX && y == interestingY {
			//log.Println("We should skip this next one....")
		} else {
			if x >= 0 && y >= 0 && x < len(data) && y < len(data[x]) {
				aString = data[x][y]
				strings = append(strings, aString)
			}
		}
		//fmt.Println(fmt.Sprintf("pictureX: %d pictureY:%d", pictureX, pictureY))
		picture[pictureX][pictureY] = aString
		pictureY++
		if pictureY == actualSight {
			pictureY = 0
			pictureX++
		}

	}

	return picture, strings
}

func occupySeat(row int, column int, data [][]string) {
	data[row][column] = "#"
}

func emptySeat(row int, column int, data [][]string) {
	data[row][column] = "L"
}

func willASeatBecomeFilled(sight int, row int, column int, data [][]string) bool {
	available := false

	seat := data[row][column]
	available = isAvailable(seat)
	if available {
		var picture [][]string
		var results []string
		var coordinates [][]int
		coordinates = diagonalToDimensions(row, column, sight)
		picture, results = populateAdjacentSeatsInGraph(sight, coordinates, data, row, column)
		found := false
		//prettyPrint(picture)

		for i := 0; !found && i < len(results); i++ {
			aPotentialSeat := results[i]
			found = !isAFloor(aPotentialSeat) && !isAvailable(aPotentialSeat)
		}
		available = !found

		pictureSays := true
		for x := 0; pictureSays && x < len(picture); x++ {
			for y := 0; pictureSays && y < len(picture[x]); y++ {
				aDot := picture[x][y]
				occupied := isOccupied(aDot)
				pictureSays = !occupied
			}
		}
		if pictureSays != available {
			log.Println("Start Data.......")
			prettyPrint(data)
			log.Println("End Data.......")

			log.Println("Start coordinates.......")
			prettyPrintInts(coordinates)
			log.Println("End coordinates.......")

			log.Println("Start results.......")
			prettyPrintSimple(results)
			log.Println("End results.......")

			log.Println("Start Picture.......")
			prettyPrint(picture)
			log.Println("End Picture.......")

			fmt.Println(fmt.Sprintf("Row: %d Column: %d", row, column))
			fmt.Println(fmt.Sprintf("Algorithm: %v Picture: %v", available, pictureSays))

			panic("Picture proof disagrees with other logic")
		}
	}

	return available
}

func prettyPrint(picture [][]string) {
	for _, row := range picture {
		fmt.Println(row)
	}
}

func prettyPrintSimple(picture []string) {
	for _, row := range picture {
		fmt.Println(row)
	}
}

func prettyPrintInts(picture [][]int) {
	for _, row := range picture {
		fmt.Println(row)
	}
}

func willASeatBecomeVacant(sight int, row int, column int, data [][]string) bool {
	willBecomeVacant := false
	seat := data[row][column]
	if !isAFloor(seat) {
		occupied := isOccupied(seat)

		if occupied {
			var picture [][]string
			var results []string
			var coordinates [][]int
			coordinates = diagonalToDimensions(row, column, sight)
			picture, results = populateAdjacentSeatsInGraph(sight, coordinates, data, row, column)

			//prettyPrint(picture)

			occupiedCount := 0
			for i := 0; i < len(results); i++ {
				aPotentialSeat := results[i]
				if isOccupied(aPotentialSeat) {
					occupiedCount++
				}
			}
			willBecomeVacant = (occupiedCount >= 4)

			pictureSays := false
			occupiedCount = 0
			for x := 0; x < len(picture); x++ {
				for y := 0; y < len(picture[x]); y++ {
					aDot := picture[x][y]
					if isOccupied(aDot) {
						occupiedCount++
					}
				}
			}
			pictureSays = (occupiedCount >= 4)
			if pictureSays != willBecomeVacant {
				panic("Picture proof disagrees with other logic")
			}
		}
	}

	return willBecomeVacant
}

func isAvailable(seat string) bool {
	return (seat == "L")
}

func isAFloor(seat string) bool {
	return (seat == ".")
}

func isOccupied(seat string) bool {
	return (seat == "#")
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

func diagonalToDimensions(i int, j int, sight int) [][]int {
	results := make([][]int, 0)

	for x := i - sight; x <= (i + sight); x++ {
		for y := j - sight; y <= (j + sight); y++ {
			coordinate := []int{x, y}
			results = append(results, coordinate)
		}
	}

	return results

}
