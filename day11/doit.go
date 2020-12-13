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

func proof(aFileName string) {
	data := Parse(aFileName)

	var realData [][]string
	realData = make([][]string, 0)

	for _, aString := range data {
		realData = append(realData, strings.Split(aString, ""))
	}

	interestingCoordinate := &coordinate{
		y: 4,
		x: 3,
	}

	//sight, coordinates, picture, results, directions, picturePoints := lineOfSight(0, interestingCoordinate, realData)
	_, _, picture, _, directions, _ := lineOfSight(0, 1, interestingCoordinate, realData)
	prettyPrint(picture)
	log.Print(directions)
}

func main() {

	proof("day_11_sample_data_2.txt")

	data := Parse("day_11_sample_data.txt")

	var realData [][]string
	realData = make([][]string, 0)

	for _, aString := range data {
		realData = append(realData, strings.Split(aString, ""))
	}

	//prettyPrint(realData)

	//var picture [][]string
	//var results []string
	//var coordinates [][]int
	//x := 0
	//y := 0
	sight := 0
	limitSightTo := 0
	tolerance := 5
	//coordinates = diagonalToDimensions(x, y, sight)
	//picture, results = populateAdjacentSeatsInGraph( coordinates, realData, x, y)
	//prettyPrint(picture)
	//log.Print(results)

	//var picture [][]string
	//var results []string
	//var coordinates []*coordinate
	//var directions []*direction
	//var picturePoints []*picturePoint
	//interestingCoordinate := &coordinate{
	//	y: 4,
	//	x: 3,
	//}
	//
	//assertEquals(realData, interestingCoordinate.y, interestingCoordinate.x, "L")
	//sight, coordinates, picture, results, directions, picturePoints = lineOfSight(sight, interestingCoordinate, realData)
	//prettyPrint(picture)
	//prettyPrintSimple(results)
	//log.Print(coordinates)
	//if len(directions) > 0 || len(picturePoints) > 0 {
	//
	//}

	changing := true
	available, occupied := count(realData)
	dataToConsider := realData
	for i := 0; changing; i++ {
		_, _, emptySeatPhaseData := adjustSeating(sight, limitSightTo, dataToConsider, tolerance)
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

func adjustSeating(sight int, limitSightTo int, originalData [][]string, tolerance int) ([][]string, [][]string, [][]string) {

	occupyPhaseData := deepCopy(originalData)

	//fmt.Println("Start..............")
	//prettyPrint(occupyPhaseData)

	for y := 0; y < len(originalData); y++ {
		for x := 0; x < len(originalData[y]); x++ {
			interestingCoordinate := &coordinate{
				y: y,
				x: x,
			}
			if willASeatBecomeFilled(sight, limitSightTo, interestingCoordinate, originalData) {
				occupySeat(interestingCoordinate, occupyPhaseData)
			}
		}
	}

	//fmt.Println("Filled Seats")
	//prettyPrint(occupyPhaseData)

	emptySeatPhaseData := deepCopy(occupyPhaseData)

	//log.Println("Adjusting seats")

	for y := 0; y < len(occupyPhaseData); y++ {
		for x := 0; x < len(occupyPhaseData[y]); x++ {
			interestingCoordinate := &coordinate{
				y: y,
				x: x,
			}
			if willASeatBecomeVacant(sight, limitSightTo, interestingCoordinate, occupyPhaseData, tolerance) {
				emptySeat(interestingCoordinate, emptySeatPhaseData)
			}
		}

	}
	fmt.Println("Emptied Seats")
	prettyPrint(emptySeatPhaseData)
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

func willASeatBecomeFilled(sight int, limitSightTo int, interestingCoordinate *coordinate, data [][]string) bool {
	available := false

	seat := data[interestingCoordinate.y][interestingCoordinate.x]
	available = isAvailable(seat)
	if available {
		var picture [][]string
		var results []string
		var picturePoints []*picturePoint
		var coordinates []*coordinate
		var directions []*direction

		//coordinates = diagonalToDimensions(interestingCoordinate, sight)
		//picture, results, picturePoints = populateAdjacentSeatsInGraph(coordinates, data, interestingCoordinate)

		sight, coordinates, picture, results, directions, picturePoints = lineOfSight(sight, limitSightTo, interestingCoordinate, data)
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
		for y := 0; pictureSays && y < len(picture); y++ {
			for x := 0; pictureSays && x < len(picture[y]); x++ {
				aDot := picture[y][x]

				aPicturePoint := picturePointHavingPictureCoordinate(&coordinate{
					y: y,
					x: x,
				}, picturePoints)

				if !aPicturePoint.center {
					occupied := isOccupied(aDot)
					pictureSays = !occupied
				}

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

func picturePointHavingRealCoordinate(aCoordinate *coordinate, picturePoints []*picturePoint) *picturePoint {
	var aPicturePoint *picturePoint

	for i := 0; aPicturePoint == nil && i < len(picturePoints); i++ {
		candidate := picturePoints[i]
		if candidate.realX == aCoordinate.x && candidate.realY == aCoordinate.y {
			aPicturePoint = candidate
		}
	}

	return aPicturePoint

}

func picturePointHavingPictureCoordinate(aCoordinate *coordinate, picturePoints []*picturePoint) *picturePoint {
	var aPicturePoint *picturePoint

	for i := 0; aPicturePoint == nil && i < len(picturePoints); i++ {
		candidate := picturePoints[i]
		if candidate.pictureX == aCoordinate.x && candidate.pictureY == aCoordinate.y {
			aPicturePoint = candidate
		}
	}

	return aPicturePoint

}

func lineOfSight(sight int, limitSightTo int, interestingCoordinate *coordinate, data [][]string) (int, []*coordinate, [][]string, []string, []*direction, []*picturePoint) {
	var picture [][]string
	var results []string
	var coordinates []*coordinate
	var picturePoints []*picturePoint
	coordinates = diagonalToDimensions(interestingCoordinate, sight)
	picture, results, picturePoints = populateAdjacentSeatsInGraph(coordinates, data, interestingCoordinate)
	//log.Println("Data")
	//prettyPrint(data)
	//log.Println("Picture")
	//prettyPrint(picture)

	var directions []*direction
	anInterestingPicturePoint := picturePointHavingRealCoordinate(interestingCoordinate, picturePoints)
	directions = peakInAllDirections(anInterestingPicturePoint, picture)

	needToContinueLooking := false
	if len(directions) > 0 {
		for i := 0; !needToContinueLooking && i < len(directions); i++ {
			aDirection := directions[i]
			if !(aDirection.outOfBounds) {
				needToContinueLooking = (aDirection.occupied == 0 && aDirection.open == 0)
			}

		}
	}

	if needToContinueLooking && limitSightTo == 0 || (limitSightTo > 0 && sight < limitSightTo) {
		sight++
		sight, coordinates, picture, results, directions, picturePoints = lineOfSight(sight, limitSightTo, interestingCoordinate, data)
	}

	return sight, coordinates, picture, results, directions, picturePoints

}

type direction struct {
	direction   int
	values      []string
	occupied    int
	open        int
	outOfBounds bool
}

func peakInAllDirections(anInterestingPicturePoint *picturePoint, picture [][]string) []*direction {

	interestingCoordinate := &coordinate{
		y: anInterestingPicturePoint.pictureY,
		x: anInterestingPicturePoint.pictureX,
	}

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
	seekAlgorithm1 := func(aCoordinate *coordinate) *coordinate {
		adjustedCoordinate := &coordinate{
			y: aCoordinate.y,
			x: aCoordinate.x + 1,
		}
		return adjustedCoordinate

	}
	/*
		2
		y = y - 1 , x = x + 1
	*/
	seekAlgorithm2 := func(aCoordinate *coordinate) *coordinate {

		adjustedCoordinate := &coordinate{
			y: aCoordinate.y + 1,
			x: aCoordinate.x + 1,
		}
		return adjustedCoordinate
	}

	/*
		3
		y = x, y = y + 1
	*/
	seekAlgorithm3 := func(aCoordinate *coordinate) *coordinate {

		adjustedCoordinate := &coordinate{
			y: aCoordinate.y + 1,
			x: aCoordinate.x,
		}
		return adjustedCoordinate
	}

	/*
		4
		y = y - 1 , x = x - 1
	*/
	seekAlgorithm4 := func(aCoordinate *coordinate) *coordinate {

		adjustedCoordinate := &coordinate{
			y: aCoordinate.y + 1,
			x: aCoordinate.x - 1,
		}
		return adjustedCoordinate
	}

	/*
		5
		y = 0, x = x - 1
	*/
	seekAlgorithm5 := func(aCoordinate *coordinate) *coordinate {

		adjustedCoordinate := &coordinate{
			y: aCoordinate.y,
			x: aCoordinate.x - 1,
		}
		return adjustedCoordinate
	}

	/*
		6
		y = y + 1 , x = x - 1
	*/
	seekAlgorithm6 := func(aCoordinate *coordinate) *coordinate {

		adjustedCoordinate := &coordinate{
			y: aCoordinate.y - 1,
			x: aCoordinate.x - 1,
		}
		return adjustedCoordinate
	}

	/*
		7
		x = 0, y = y - 1
	*/
	seekAlgorithm7 := func(aCoordinate *coordinate) *coordinate {

		adjustedCoordinate := &coordinate{
			y: aCoordinate.y - 1,
			x: aCoordinate.x,
		}
		return adjustedCoordinate
	}

	/*
		8
		y = y + 1 , x = x + 1
	*/
	seekAlgorithm8 := func(aCoordinate *coordinate) *coordinate {

		adjustedCoordinate := &coordinate{
			y: aCoordinate.y - 1,
			x: aCoordinate.x + 1,
		}
		return adjustedCoordinate
	}
	//log.Println("===========================================")
	//prettyPrint(picture)
	//log.Println("===========================================")
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

func processAlgorithmToFindFirstSeatAndUpdatePicture(picture [][]string, interestingCoordinate *coordinate, seekAlgorithm func(aCoordinate *coordinate) *coordinate, aDirection *direction) *direction {

	graphBounds := &coordinate{
		y: len(picture),
		x: len(picture[len(picture)-1]),
	}

	data := ""

	stopAlgorithm := func(graphBoundsCoordinate *coordinate, aCoordinate *coordinate, value string) (stop bool) {
		stop = false
		stop = (aCoordinate.x < 0 || aCoordinate.x >= graphBoundsCoordinate.x)
		stop = stop || (aCoordinate.y < 0 || aCoordinate.y >= graphBoundsCoordinate.y)

		return stop
	}

	successAlgorithm := func(value string, valuesToFind ...string) (success bool) {
		success = false
		for i := 0; !success && i < len(valuesToFind); i++ {
			success = (value == valuesToFind[i])
		}
		return success
	}

	adjustedCoordinate := &coordinate{
		y: interestingCoordinate.y,
		x: interestingCoordinate.x,
	}

	occupied := 0
	open := 0
	outOfBounds := false
	for stop := false; !stop; stop = stopAlgorithm(graphBounds, adjustedCoordinate, data) {
		if !(adjustedCoordinate.x == interestingCoordinate.x && adjustedCoordinate.y == interestingCoordinate.y) {

			data = picture[adjustedCoordinate.y][adjustedCoordinate.x]
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

		adjustedCoordinate = seekAlgorithm(adjustedCoordinate)
	}
	aDirection.occupied = occupied
	aDirection.open = open
	aDirection.outOfBounds = outOfBounds
	return aDirection
}

func willASeatBecomeVacant(sight int, limitSightTo int, interestingCoordinate *coordinate, data [][]string, tolerance int) bool {
	willBecomeVacant := false
	seat := data[interestingCoordinate.y][interestingCoordinate.x]
	if !isAFloor(seat) {
		occupied := isOccupied(seat)

		if occupied {
			var picture [][]string
			var results []string
			var coordinates []*coordinate
			var picturePoints []*picturePoint
			var directions []*direction
			//coordinates = diagonalToDimensions(interestingCoordinate, sight)
			//picture, results, picturePoints = populateAdjacentSeatsInGraph(coordinates, data, interestingCoordinate)

			sight, coordinates, picture, results, directions, picturePoints = lineOfSight(sight, limitSightTo, interestingCoordinate, data)

			if len(coordinates) > 0 {

			}
			if len(directions) > 0 {

			}
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
			for y := 0; y < len(picture); y++ {
				for x := 0; x < len(picture[y]); x++ {
					aDot := picture[y][x]

					aPicturePoint := picturePointHavingPictureCoordinate(&coordinate{
						y: y,
						x: x,
					}, picturePoints)

					if !aPicturePoint.center {
						if isOccupied(aDot) {
							occupiedCount++
						}
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
