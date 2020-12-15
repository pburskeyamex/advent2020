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

type vision struct {
	sight                   int
	allowedToIncreaseVision bool
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

	sight := &vision{
		sight:                   0,
		allowedToIncreaseVision: true,
	}

	//sight, coordinates, picture, results, directions, picturePoints := lineOfSight(0, interestingCoordinate, realData)
	_, _, _, directions := lineOfSight(sight, interestingCoordinate, realData)
	//prettyPrint(picture)
	log.Print(directions)
}

func main() {

	//proof("day_11_sample_data_2.txt")

	//data := Parse("day_11_sample_data.txt")
	data := Parse("day_11_data.txt")

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
	debug := false

	//sight := &vision{
	//	sight:                   1,
	//	allowedToIncreaseVision: false,
	//}
	//tolerance := 4

	sight := &vision{
		sight:                   0,
		allowedToIncreaseVision: true,
	}
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
		_, _, emptySeatPhaseData := adjustSeating(debug, sight, dataToConsider, tolerance)
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

	if occupied != 1862 {
		//	if occupied != 37{
		panic("should be 1862")
	}

}

func adjustSeating(debug bool, sight *vision, originalData [][]string, tolerance int) ([][]string, [][]string, [][]string) {

	occupyPhaseData := deepCopy(originalData)

	if debug {
		fmt.Println("Start..............")
		prettyPrint(occupyPhaseData)
	}

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
	if debug {
		fmt.Println("Filled Seats")
		prettyPrint(occupyPhaseData)
	}

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
	if debug {
		fmt.Println("Emptied Seats")
		prettyPrint(emptySeatPhaseData)
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

	for i := 0; i < len(coordinates); i++ {
		aCoordinate := coordinates[i]
		value := valueFunction(aCoordinate)
		aMap[value] = value

	}

	return len(aMap)
}

func determineDistanceBetweenHighAndLowPoints(points []*picturePoint, valueFunction func(aPoint *picturePoint) int) int {

	aMap := make(map[int]int)

	for i := 0; i < len(points); i++ {
		aPoint := points[i]
		value := valueFunction(aPoint)
		aMap[value] = value

	}

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

func adjustSightBasedOnPicturePoints(points []*picturePoint) int {

	actualSight := 0

	xDifference := determineDistanceBetweenHighAndLowPoints(points, func(aCoordinate *picturePoint) int {
		return aCoordinate.pictureX
	})

	yDifference := determineDistanceBetweenHighAndLowPoints(points, func(aCoordinate *picturePoint) int {
		return aCoordinate.pictureY
	})

	if xDifference > actualSight {
		actualSight = xDifference
	}

	if yDifference > actualSight {
		actualSight = yDifference
	}

	return actualSight
}

func generatePictureFrom(points []*picturePoint) [][]string {
	var picture [][]string

	coordinates := make([]*coordinate, 0)
	for i := 0; i < len(points); i++ {

	}

	actualSight := adjustSightBasedOnCoordinates(coordinates)

	picture = make([][]string, actualSight)
	panic("Not yet implemented")
	return picture

}

func populateAdjacentSeatsInGraph(coordinates []*coordinate, data [][]string, interestingCoordinate *coordinate) []*picturePoint {

	//var picture [][]string
	var picturePoints []*picturePoint
	picturePoints = make([]*picturePoint, 0)
	//strings = make([]string, 0)

	actualSight := adjustSightBasedOnCoordinates(coordinates)

	if actualSight > 0 {
		//picture = make([][]string, actualSight)
		//for i := 0; i < len(picture); i++ {
		//	picture[i] = make([]string, actualSight)
		//
		//}
		pictureX := 0
		pictureY := 0
		for i := 0; i < len(coordinates); i++ {
			aCoordinate := coordinates[i]
			realX := aCoordinate.x
			realY := aCoordinate.y

			aString := ""

			centerIdentified := false

			//log.Println(fmt.Sprintf("X %d  Y %d" ,x, y))
			if realX == interestingCoordinate.x && realY == interestingCoordinate.y {
				//log.Println("We should skip this next one....")
				centerIdentified = true
				aString = data[realY][realX]

			} else {
				if realX >= 0 && realY >= 0 && realY < len(data) && realX < len(data[realY]) {
					aString = data[realY][realX]
					//strings = append(strings, aString)
				}
			}
			//fmt.Println(fmt.Sprintf("pictureX: %d pictureY:%d", pictureX, pictureY))

			aPicturePoint := &picturePoint{
				realX:    realX,
				realY:    realY,
				pictureX: pictureX,
				pictureY: pictureY,
				value:    aString,
				center:   centerIdentified,
			}

			picturePoints = append(picturePoints, aPicturePoint)

			//picture[pictureY][pictureX] = aString
			pictureY++
			if pictureY == actualSight {
				pictureY = 0
				pictureX++
			}

		}
	}

	return picturePoints
}

func occupySeat(interestingCoordinate *coordinate, data [][]string) {
	data[interestingCoordinate.y][interestingCoordinate.x] = "#"
}

func emptySeat(interestingCoordinate *coordinate, data [][]string) {
	data[interestingCoordinate.y][interestingCoordinate.x] = "L"
}

func willASeatBecomeFilled(sight *vision, interestingCoordinate *coordinate, data [][]string) bool {
	available := false

	seat := data[interestingCoordinate.y][interestingCoordinate.x]
	available = isAvailable(seat)
	if available {
		//var picture [][]string
		var picturePoints []*picturePoint
		var coordinates []*coordinate
		var directions []*direction

		//coordinates = diagonalToDimensions(interestingCoordinate, sight)
		//picture, results, picturePoints = populateAdjacentSeatsInGraph(coordinates, data, interestingCoordinate)

		sight, coordinates, directions, picturePoints = lineOfSight(sight, interestingCoordinate, data)
		if len(directions) > 0 {

		}

		//found := false
		//prettyPrint(picture)
		if len(picturePoints) > 0 {

		}

		//if len(picture) > 0 {
		//
		//}

		if len(coordinates) > 0 {

		}

		occupiedCount := 0
		for i := 0; i < len(directions); i++ {
			aDirection := directions[i]
			occupiedCount = occupiedCount + aDirection.occupied

		}

		//for i := 0; !found && i < len(results); i++ {
		//	aPotentialSeat := results[i]
		//	found = !isAFloor(aPotentialSeat) && !isAvailable(aPotentialSeat)
		//}
		//available = !found

		available = occupiedCount == 0
		//pictureSays := true
		//for y := 0; pictureSays && y < len(picture); y++ {
		//	for x := 0; pictureSays && x < len(picture[y]); x++ {
		//		aDot := picture[y][x]
		//
		//		aPicturePoint := picturePointHavingPictureCoordinate(&coordinate{
		//			y: y,
		//			x: x,
		//		}, picturePoints)
		//
		//		if !aPicturePoint.center {
		//			occupied := isOccupied(aDot)
		//			pictureSays = !occupied
		//		}
		//
		//	}
		//}
		//if pictureSays != available {
		//	log.Println("Start Data.......")
		//	prettyPrint(data)
		//	log.Println("End Data.......")
		//
		//	log.Println("Start coordinates.......")
		//	log.Print(coordinates)
		//	log.Println("End coordinates.......")
		//
		//	log.Println("Start results.......")
		//	prettyPrintSimple(results)
		//	log.Println("End results.......")
		//
		//	log.Println("Start Picture.......")
		//	prettyPrint(picture)
		//	log.Println("End Picture.......")
		//
		//	fmt.Println(fmt.Sprintf("Row: %d Column: %d", interestingCoordinate.y, interestingCoordinate.x))
		//	fmt.Println(fmt.Sprintf("Algorithm: %v Picture: %v", available, pictureSays))
		//
		//	panic("Picture proof disagrees with other logic")
		//}
	}

	return available
}

func prettyPrintDirection(directions []*direction) {

	for i := 0; i < len(directions); i++ {
		aDirection := directions[i]
		fmt.Println(fmt.Sprintf("Direction: %d Occupied: %d Open: %d Out of Bounds: %v", aDirection.direction, aDirection.occupied, aDirection.open, aDirection.outOfBounds))
	}
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

func lineOfSight(sight *vision, interestingCoordinate *coordinate, data [][]string) (*vision, []*coordinate, []*direction, []*picturePoint) {

	var coordinates []*coordinate
	var picturePoints []*picturePoint
	coordinates = diagonalToDimensions(interestingCoordinate, sight)
	picturePoints = populateAdjacentSeatsInGraph(coordinates, data, interestingCoordinate)
	//log.Println("Data")
	//prettyPrint(data)
	//log.Println("Picture")
	//prettyPrint(picture)

	var directions []*direction
	anInterestingPicturePoint := picturePointHavingRealCoordinate(interestingCoordinate, picturePoints)
	directions = peakInAllDirections(anInterestingPicturePoint, picturePoints)

	needToContinueLooking := false
	if len(directions) > 0 {
		for i := 0; !needToContinueLooking && i < len(directions); i++ {
			aDirection := directions[i]
			if !(aDirection.outOfBounds) {
				needToContinueLooking = (aDirection.occupied == 0 && aDirection.open == 0)
			}

		}
	}

	increaseVision := needToContinueLooking && (sight.allowedToIncreaseVision)
	if increaseVision {
		sight.sight++
		sight, coordinates, directions, picturePoints = lineOfSight(sight, interestingCoordinate, data)
	}

	return sight, coordinates, directions, picturePoints

}

type direction struct {
	direction   int
	values      []string
	occupied    int
	open        int
	outOfBounds bool
}

func peakInAllDirections(anInterestingPicturePoint *picturePoint, picturePoints []*picturePoint) []*direction {

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
	graphBounds := graphBoundsFrom(picturePoints)

	//if picture != nil {
	if len(picturePoints) > 0 {
		processAlgorithmToFindFirstSeatAndUpdatePicture(graphBounds, picturePoints, interestingCoordinate, seekAlgorithm1, directions[0])
		processAlgorithmToFindFirstSeatAndUpdatePicture(graphBounds, picturePoints, interestingCoordinate, seekAlgorithm2, directions[1])
		processAlgorithmToFindFirstSeatAndUpdatePicture(graphBounds, picturePoints, interestingCoordinate, seekAlgorithm3, directions[2])
		processAlgorithmToFindFirstSeatAndUpdatePicture(graphBounds, picturePoints, interestingCoordinate, seekAlgorithm4, directions[3])
		processAlgorithmToFindFirstSeatAndUpdatePicture(graphBounds, picturePoints, interestingCoordinate, seekAlgorithm5, directions[4])
		processAlgorithmToFindFirstSeatAndUpdatePicture(graphBounds, picturePoints, interestingCoordinate, seekAlgorithm6, directions[5])
		processAlgorithmToFindFirstSeatAndUpdatePicture(graphBounds, picturePoints, interestingCoordinate, seekAlgorithm7, directions[6])
		processAlgorithmToFindFirstSeatAndUpdatePicture(graphBounds, picturePoints, interestingCoordinate, seekAlgorithm8, directions[7])
	}

	//fmt.Println(fmt.Sprintf("Occupied %d Open: %d", occupied, open))
	//prettyPrint(picture)

	return directions
}

func graphBoundsFrom(picturePoints []*picturePoint) *coordinate {

	length := adjustSightBasedOnPicturePoints(picturePoints)

	y := length
	x := length

	graphBounds := &coordinate{
		y: y,
		x: x,
	}
	return graphBounds
}

func processAlgorithmToFindFirstSeatAndUpdatePicture(graphBounds *coordinate, picturePoints []*picturePoint, interestingCoordinate *coordinate, seekAlgorithm func(aCoordinate *coordinate) *coordinate, aDirection *direction) *direction {

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

			//data = picture[adjustedCoordinate.y][adjustedCoordinate.x]

			aPictureCoordinate := &coordinate{
				y: adjustedCoordinate.y,
				x: adjustedCoordinate.x,
			}
			aPicturePointForData := picturePointHavingPictureCoordinate(aPictureCoordinate, picturePoints)

			//if data != aPicturePointForData.value{
			//	log.Printf("Picture :y %d :x %d value: %s mismatch picture point :y %d :x %d value: %s", adjustedCoordinate.y, adjustedCoordinate.x, data, aPictureCoordinate.y, aPictureCoordinate.x, aPicturePointForData.value)
			//}

			data = aPicturePointForData.value
			aDirection.values = append(aDirection.values, data)

			if successAlgorithm(data, "#") {
				occupied++
				break
			} else if successAlgorithm(data, "L") {
				open++
				break
			} else if successAlgorithm(data, "") {
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

func willASeatBecomeVacant(sight *vision, interestingCoordinate *coordinate, data [][]string, tolerance int) bool {
	willBecomeVacant := false
	seat := data[interestingCoordinate.y][interestingCoordinate.x]
	if !isAFloor(seat) {
		occupied := isOccupied(seat)

		if occupied {
			//var picture [][]string
			var coordinates []*coordinate
			var picturePoints []*picturePoint
			var directions []*direction
			//coordinates = diagonalToDimensions(interestingCoordinate, sight)
			//picture, results, picturePoints = populateAdjacentSeatsInGraph(coordinates, data, interestingCoordinate)

			sight, coordinates, directions, picturePoints = lineOfSight(sight, interestingCoordinate, data)

			if len(coordinates) > 0 {

			}
			if len(directions) > 0 {

			}
			if len(picturePoints) > 0 {

			}

			//if len(picture) > 0 {
			//
			//}

			//prettyPrint(picture)

			occupiedCount := 0
			//for i := 0; i < len(results); i++ {
			//	aPotentialSeat := results[i]
			//	if isOccupied(aPotentialSeat) {
			//		occupiedCount++
			//	}
			//}
			//willBecomeVacant = (occupiedCount >= tolerance)

			//pictureSays := false
			//occupiedCount = 0
			for i := 0; i < len(directions); i++ {
				aDirection := directions[i]
				occupiedCount = occupiedCount + aDirection.occupied

			}
			willBecomeVacant = (occupiedCount >= tolerance)
			//pictureSays = (occupiedCount >= tolerance)
			//if pictureSays != willBecomeVacant {
			//	panic("Picture proof disagrees with other logic")
			//}
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

func diagonalToDimensions(interestingCoordinate *coordinate, sight *vision) []*coordinate {
	results := make([]*coordinate, 0)

	for y := interestingCoordinate.y - sight.sight; y <= (interestingCoordinate.y + sight.sight); y++ {
		for x := interestingCoordinate.x - sight.sight; x <= (interestingCoordinate.x + sight.sight); x++ {

			coordinate := &coordinate{
				y: y,
				x: x,
			}
			results = append(results, coordinate)

		}
	}

	return results

}
