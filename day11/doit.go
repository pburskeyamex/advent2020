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

type coordinate struct {
	y int
	x int
}

type picturePoint struct {
	realX    int
	realY    int
	pictureX int
	pictureY int
	value    string
	center   bool
}

func assertEquals(data [][]string, y int, x int, value string) {
	if data[y][x] != value {
		panic("Data mismatch")
	}
}

func main() {

	data := Parse("day_11_sample_data_2.txt")

	var realData [][]string
	realData = make([][]string, 0)

	for _, aString := range data {
		realData = append(realData, strings.Split(aString, ""))
	}

	prettyPrint(realData)

	//var picture [][]string
	//var results []string
	//var coordinates [][]int
	//x := 0
	//y := 0
	sight := 0
	tolerance := 5
	//coordinates = diagonalToDimensions(x, y, sight)
	//picture, results = populateAdjacentSeatsInGraph( coordinates, realData, x, y)
	//prettyPrint(picture)
	//log.Print(results)

	var picture [][]string
	var results []string
	var coordinates []*coordinate
	var directions []*direction
	var picturePoints []*picturePoint
	interestingCoordinate := &coordinate{
		y: 4,
		x: 3,
	}

	assertEquals(realData, interestingCoordinate.y, interestingCoordinate.x, "L")
	sight, coordinates, picture, results, directions, picturePoints = lineOfSight(sight, interestingCoordinate, realData)
	prettyPrint(picture)
	prettyPrintSimple(results)
	log.Print(coordinates)
	if len(directions) > 0 || len(picturePoints) > 0 {

	}

	changing := true
	available, occupied := count(realData)
	dataToConsider := realData
	for i := 0; changing; i++ {
		_, _, emptySeatPhaseData := adjustSeating(sight, dataToConsider, tolerance)
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

func adjustSeating(sight int, originalData [][]string, tolerance int) ([][]string, [][]string, [][]string) {
	occupyPhaseData := deepCopy(originalData)
	for y := 0; y < len(originalData); y++ {
		for x := 0; x < len(originalData[y]); x++ {
			interestingCoordinate := &coordinate{
				y: y,
				x: x,
			}
			if willASeatBecomeFilled(sight, interestingCoordinate, originalData) {
				occupySeat(interestingCoordinate, occupyPhaseData)
			}
		}
	}

	//prettyPrint(occupyPhaseData)
	emptySeatPhaseData := deepCopy(occupyPhaseData)

	//log.Println("Adjusting seats")

	for y := 0; y < len(occupyPhaseData); y++ {
		for x := 0; x < len(occupyPhaseData[y]); x++ {
			interestingCoordinate := &coordinate{
				y: y,
				x: x,
			}
			if willASeatBecomeVacant(sight, interestingCoordinate, occupyPhaseData, tolerance) {
				emptySeat(interestingCoordinate, emptySeatPhaseData)
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

func determineDistanceBetweenHighAndLow(coordinates []*coordinate, valueFunction func(aCoordinate *coordinate) int) int {

	aMap := make(map[int]int)
	//floor := -1
	//ceiling := -1
	for i := 0; i < len(coordinates); i++ {
		aCoordinate := coordinates[i]
		value := valueFunction(aCoordinate)
		aMap[value] = value
		//if floor < 0 || value < floor {
		//	floor = value
		//}
		//if ceiling < 0 || value > ceiling {
		//	ceiling = value
		//}
	}

	//return ceiling - floor
	return len(aMap)
}

func adjustSightBasedOnCoordinates(coordinates []*coordinate) int {

	actualSight := 0

	xDifference := determineDistanceBetweenHighAndLow(coordinates, func(aCoordinate *coordinate) int {
		return aCoordinate.x
	})

	yDifference := determineDistanceBetweenHighAndLow(coordinates, func(aCoordinate *coordinate) int {
		return aCoordinate.y
	})

	if xDifference > actualSight {
		actualSight = xDifference
	}

	if yDifference > actualSight {
		actualSight = yDifference
	}

	return actualSight
}

func populateAdjacentSeatsInGraph(coordinates []*coordinate, data [][]string, interestingCoordinate *coordinate) ([][]string, []string, []*picturePoint) {

	var strings []string
	var picture [][]string
	var picturePoints []*picturePoint
	picturePoints = make([]*picturePoint, 0)
	strings = make([]string, 0)

	actualSight := adjustSightBasedOnCoordinates(coordinates)

	if actualSight > 0 {
		picture = make([][]string, actualSight)
		for i := 0; i < len(picture); i++ {
			picture[i] = make([]string, actualSight)

		}
		pictureX := 0
		pictureY := 0
		for i := 0; i < len(coordinates); i++ {
			aCoordinate := coordinates[i]
			x := aCoordinate.x
			y := aCoordinate.y

			aString := " "

			centerIdentified := false

			//log.Println(fmt.Sprintf("X %d  Y %d" ,x, y))
			if x == interestingCoordinate.x && y == interestingCoordinate.y {
				//log.Println("We should skip this next one....")
				centerIdentified = true
				aString = data[y][x]

			} else {
				if x >= 0 && y >= 0 && y < len(data) && x < len(data[y]) {
					aString = data[y][x]
					strings = append(strings, aString)
				}
			}
			//fmt.Println(fmt.Sprintf("pictureX: %d pictureY:%d", pictureX, pictureY))

			aPicturePoint := &picturePoint{
				realX:    x,
				realY:    y,
				pictureX: pictureX,
				pictureY: pictureY,
				value:    aString,
				center:   centerIdentified,
			}

			picturePoints = append(picturePoints, aPicturePoint)

			picture[pictureX][pictureY] = aString
			pictureY++
			if pictureY == actualSight {
				pictureY = 0
				pictureX++
			}

		}
	}

	return picture, strings, picturePoints
}

func occupySeat(interestingCoordinate *coordinate, data [][]string) {
	data[interestingCoordinate.y][interestingCoordinate.x] = "#"
}

func emptySeat(interestingCoordinate *coordinate, data [][]string) {
	data[interestingCoordinate.y][interestingCoordinate.x] = "L"
}

func willASeatBecomeFilled(sight int, interestingCoordinate *coordinate, data [][]string) bool {
	available := false

	seat := data[interestingCoordinate.y][interestingCoordinate.x]
	available = isAvailable(seat)
	if available {
		var picture [][]string
		var results []string
		var picturePoints []*picturePoint
		var coordinates []*coordinate
		var directions []*direction

		coordinates = diagonalToDimensions(interestingCoordinate, sight)
		picture, results, picturePoints = populateAdjacentSeatsInGraph(coordinates, data, interestingCoordinate)

		sight, coordinates, picture, results, directions, picturePoints = lineOfSight(sight, interestingCoordinate, data)
		if len(directions) > 0 {

		}

		found := false
		//prettyPrint(picture)
		if len(picturePoints) > 0 {

		}

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
			log.Print(coordinates)
			log.Println("End coordinates.......")

			log.Println("Start results.......")
			prettyPrintSimple(results)
			log.Println("End results.......")

			log.Println("Start Picture.......")
			prettyPrint(picture)
			log.Println("End Picture.......")

			fmt.Println(fmt.Sprintf("Row: %d Column: %d", interestingCoordinate.y, interestingCoordinate.x))
			fmt.Println(fmt.Sprintf("Algorithm: %v Picture: %v", available, pictureSays))

			panic("Picture proof disagrees with other logic")
		}
	}

	return available
}

func prettyPrint(picture [][]string) {
	if len(picture) > 0 {
		fmt.Println("X <->")
		fmt.Print(" ")
		for i := 0; i < len(picture[0]); i++ {
			fmt.Print(fmt.Sprintf("%d ", i))
		}
		fmt.Print("\n")
	}

	for i := 0; i < len(picture); i++ {
		fmt.Print(picture[i])
		fmt.Print(fmt.Sprintf("%d\n", i))

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

func lineOfSight(sight int, interestingCoordinate *coordinate, data [][]string) (int, []*coordinate, [][]string, []string, []*direction, []*picturePoint) {
	var picture [][]string
	var results []string
	var coordinates []*coordinate
	var picturePoints []*picturePoint
	coordinates = diagonalToDimensions(interestingCoordinate, sight)
	picture, results, picturePoints = populateAdjacentSeatsInGraph(coordinates, data, interestingCoordinate)
	log.Println("Data")
	prettyPrint(data)
	log.Println("Picture")
	prettyPrint(picture)

	var directions []*direction
	directions = isSeatVisibleInAllDirections(sight, interestingCoordinate, picture, results, picturePoints)

	needToContinueLooking := false
	if len(directions) > 0 {
		for i := 0; !needToContinueLooking && i < len(directions); i++ {
			aDirection := directions[i]
			if !(aDirection.outOfBounds) {
				needToContinueLooking = (aDirection.occupied == 0 && aDirection.open == 0)
			}

		}
	}

	if needToContinueLooking {
		sight++
		sight, coordinates, picture, results, directions, picturePoints = lineOfSight(sight, interestingCoordinate, data)
	}

	return sight, coordinates, picture, results, directions, picturePoints

}

func findMiddle(picture [][]string, adjustment int) int {
	middle := 0
	if len(picture) == 1 {
		middle = 1
	} else {
		middle = (len(picture) / 2) + adjustment

		offsetXLeft := 0
		for x := 0; x < middle; x++ {
			offsetXLeft++
		}

		offsetXRight := 0
		for x := len(picture) - 1; x > middle; x-- {
			offsetXRight++
		}

		if offsetXLeft != offsetXRight {
			middle = findMiddle(picture, adjustment+1)
		}
	}

	return middle
}

func validateCoordinate(y int, x int) bool {
	/*
		if y, x are equal

		if y = 0 and x is a valid number
		if x = 0 and y is a valid number
	*/
	valid := y == x

	if !valid {
		valid = (x == 0 || (y == 0))
	}

	if !valid {
		if x == y {
			log.Fatal(fmt.Sprintf("Y: %d X: %d", y, x))
			panic("Assertion about validating cartesian coordinates is wrong")
		}
	}

	return valid

}

type direction struct {
	direction   int
	values      []string
	occupied    int
	open        int
	outOfBounds bool
}

func isSeatVisibleInAllDirections(sight int, interestingCoordinate *coordinate, picture [][]string, results []string, picturePoints []*picturePoint) []*direction {

	directions := make([]*direction, 8)
	for i := 0; i < 8; i++ {
		aDirection := &direction{
			direction:   i + 1,
			values:      make([]string, 0),
			occupied:    0,
			open:        0,
			outOfBounds: false,
		}
		directions[i] = aDirection
	}

	/*
		1
		y = 0, x = x + 1
	*/
	seekAlgorithm1 := func(x int, y int) (adjustedX int, adjustedY int) {
		adjustedX = x + 1
		adjustedY = y
		return
	}
	/*
		2
		y = y - 1 , x = x + 1
	*/
	seekAlgorithm2 := func(x int, y int) (adjustedX int, adjustedY int) {
		adjustedX = x + 1
		adjustedY = y - 1
		return
	}

	/*
		3
		y = x, y = y + 1
	*/
	seekAlgorithm3 := func(x int, y int) (adjustedX int, adjustedY int) {
		adjustedX = x
		adjustedY = y + 1
		return
	}

	/*
		4
		y = y - 1 , x = x - 1
	*/
	seekAlgorithm4 := func(x int, y int) (adjustedX int, adjustedY int) {
		adjustedX = x - 1
		adjustedY = y - 1
		return
	}

	/*
		5
		y = 0, x = x - 1
	*/
	seekAlgorithm5 := func(x int, y int) (adjustedX int, adjustedY int) {
		adjustedX = x - 1
		adjustedY = y
		return
	}

	/*
		6
		y = y + 1 , x = x - 1
	*/
	seekAlgorithm6 := func(x int, y int) (adjustedX int, adjustedY int) {
		adjustedX = x - 1
		adjustedY = y + 1
		return
	}

	/*
		7
		x = 0, y = y - 1
	*/
	seekAlgorithm7 := func(x int, y int) (adjustedX int, adjustedY int) {
		adjustedX = x
		adjustedY = y - 1
		return
	}

	/*
		8
		y = y + 1 , x = x + 1
	*/
	seekAlgorithm8 := func(x int, y int) (adjustedX int, adjustedY int) {
		adjustedX = x + 1
		adjustedY = y + 1
		return
	}
	log.Println("===========================================")
	prettyPrint(picture)
	log.Println("===========================================")
	if picture != nil {
		processAlgorithmToFindFirstSeatAndUpdatePicture(picture, interestingCoordinate, seekAlgorithm1, directions[0])
		processAlgorithmToFindFirstSeatAndUpdatePicture(picture, interestingCoordinate, seekAlgorithm2, directions[1])
		processAlgorithmToFindFirstSeatAndUpdatePicture(picture, interestingCoordinate, seekAlgorithm3, directions[2])
		processAlgorithmToFindFirstSeatAndUpdatePicture(picture, interestingCoordinate, seekAlgorithm4, directions[3])
		processAlgorithmToFindFirstSeatAndUpdatePicture(picture, interestingCoordinate, seekAlgorithm5, directions[4])
		processAlgorithmToFindFirstSeatAndUpdatePicture(picture, interestingCoordinate, seekAlgorithm6, directions[5])
		processAlgorithmToFindFirstSeatAndUpdatePicture(picture, interestingCoordinate, seekAlgorithm7, directions[6])
		processAlgorithmToFindFirstSeatAndUpdatePicture(picture, interestingCoordinate, seekAlgorithm8, directions[7])
	}

	//fmt.Println(fmt.Sprintf("Occupied %d Open: %d", occupied, open))
	//prettyPrint(picture)

	return directions
}

func processAlgorithmToFindFirstSeatAndUpdatePicture(picture [][]string, interestingCoordinate *coordinate, seekAlgorithm func(x int, y int) (adjustedX int, adjustedY int), aDirection *direction) *direction {
	var graphBoundsX, graphBoundsY int
	graphBoundsY = len(picture)
	graphBoundsX = len(picture[graphBoundsY-1])
	data := ""

	stopAlgorithm := func(graphBoundsX int, graphBoundsY int, x int, y int, value string) (stop bool) {
		stop = false
		stop = (x < 0 || x >= graphBoundsX)
		stop = stop || (y < 0 || y >= graphBoundsY)

		return stop
	}

	successAlgorithm := func(value string, valuesToFind ...string) (success bool) {
		success = false
		for i := 0; !success && i < len(valuesToFind); i++ {
			success = (value == valuesToFind[i])
		}
		return success
	}

	var adjustedX, adjustedY int
	adjustedX = interestingCoordinate.x
	adjustedY = interestingCoordinate.y

	occupied := 0
	open := 0
	outOfBounds := false
	for stop := false; !stop; stop = stopAlgorithm(graphBoundsX, graphBoundsY, adjustedX, adjustedY, data) {
		if !(adjustedX == interestingCoordinate.x && adjustedY == interestingCoordinate.y) {

			data = picture[adjustedY][adjustedX]
			aDirection.values = append(aDirection.values, data)

			if successAlgorithm(data, "#") {
				occupied++
				break
			} else if successAlgorithm(data, "L") {
				open++
				break
			} else if successAlgorithm(data, " ") {
				outOfBounds = true
				break
			}
		}
		adjustedX, adjustedY = seekAlgorithm(adjustedX, adjustedY)
	}
	aDirection.occupied = occupied
	aDirection.open = open
	aDirection.outOfBounds = outOfBounds
	return aDirection
}

func willASeatBecomeVacant(sight int, interestingCoordinate *coordinate, data [][]string, tolerance int) bool {
	willBecomeVacant := false
	seat := data[interestingCoordinate.y][interestingCoordinate.x]
	if !isAFloor(seat) {
		occupied := isOccupied(seat)

		if occupied {
			var picture [][]string
			var results []string
			var coordinates []*coordinate
			var picturePoints []*picturePoint
			coordinates = diagonalToDimensions(interestingCoordinate, sight)
			picture, results, picturePoints = populateAdjacentSeatsInGraph(coordinates, data, interestingCoordinate)

			if len(picturePoints) > 0 {

			}
			//prettyPrint(picture)

			occupiedCount := 0
			for i := 0; i < len(results); i++ {
				aPotentialSeat := results[i]
				if isOccupied(aPotentialSeat) {
					occupiedCount++
				}
			}
			willBecomeVacant = (occupiedCount >= tolerance)

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
			pictureSays = (occupiedCount >= tolerance)
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

func diagonalToDimensions(interestingCoordinate *coordinate, sight int) []*coordinate {
	results := make([]*coordinate, 0)

	for y := interestingCoordinate.y - sight; y <= (interestingCoordinate.y + sight); y++ {
		for x := interestingCoordinate.x - sight; x <= (interestingCoordinate.x + sight); x++ {

			coordinate := &coordinate{
				y: y,
				x: x,
			}
			results = append(results, coordinate)

		}
	}

	return results

}
