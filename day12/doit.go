package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	strings2 "strings"
)

var (
	EAST = &compassDirection{
		value:       1,
		description: "EAST",
	}

	SOUTHEAST = &compassDirection{
		value:       2,
		description: "SOUTHEAST",
	}
	SOUTH = &compassDirection{
		value:       3,
		description: "SOUTH",
	}
	SOUTHWEST = &compassDirection{
		value:       4,
		description: "SOUTHWEST",
	}
	WEST = &compassDirection{
		value:       5,
		description: "WEST",
	}
	NORTHWEST = &compassDirection{
		value:       6,
		description: "NORTHWEST",
	}
	NORTH = &compassDirection{
		value:       7,
		description: "NORTH",
	}
	NORTHEAST = &compassDirection{
		value:       8,
		description: "NORTHEAST",
	}

	directionMap = map[int]*compassDirection{
		EAST.value:      EAST,
		SOUTHEAST.value: SOUTHEAST,
		SOUTH.value:     SOUTH,
		SOUTHWEST.value: SOUTHWEST,
		WEST.value:      WEST,
		NORTHWEST.value: NORTHWEST,
		NORTH.value:     NORTH,
		NORTHEAST.value: NORTHEAST,
	}
)

/*
Action N means to move north by the given value.
Action S means to move south by the given value.
Action E means to move east by the given value.
Action W means to move west by the given value.
Action L means to turn left the given number of degrees.
Action R means to turn right the given number of degrees.
Action F means to move forward by the given value in the direction the ship is currently facing.
*/
type direction struct {
	value string
}

type waypoint struct {
	direction *compassDirection
	value     int
}

type movement struct {
	direction *direction
	value     int
}

type compassDirection struct {
	value       int
	description string
}

type position struct {
	orientation       *compassDirection
	bearingMovement   *bearing
	manhattanDistance *manhattanDistance
}

type bearing struct {
	direction *compassDirection
	distance  int
}

//
//func proof(aFileName string) {
//	data := Parse(aFileName)
//
//	var realData [][]string
//	realData = make([][]string, 0)
//
//	for _, aString := range data {
//		realData = append(realData, strings.Split(aString, ""))
//	}
//
//	interestingCoordinate := &coordinate{
//		y: 4,
//		x: 3,
//	}
//
//	sight := &vision{
//		sight:                   0,
//		allowedToIncreaseVision: true,
//	}
//
//	//sight, coordinates, picture, results, directions, picturePoints := lineOfSight(0, interestingCoordinate, realData)
//	_, _, _, directions := lineOfSight(sight, interestingCoordinate, realData)
//	//prettyPrint(picture)
//	log.Print(directions)
//}

func parse(aString string) *movement {

	strings := strings2.Split(aString, "")
	aFunction := strings[0]
	chars := strings[1:]
	movementsString := strings2.Join(chars, "")
	aMovement, _ := strconv.Atoi(movementsString)

	aMove := &movement{
		direction: &direction{value: aFunction},
		value:     aMovement,
	}

	return aMove
}

func main() {

	//proof("day_11_sample_data_2.txt")

	//data := Parse("day_11_sample_data.txt")
	data := Parse("day_12_data.txt")

	movements := make([]*movement, 0)

	for _, aString := range data {
		aMovement := parse(aString)
		movements = append(movements, aMovement)
	}

	positions := make([]*position, 0)

	startingPosition := &position{
		orientation: EAST,
		bearingMovement: &bearing{
			direction: EAST,
			distance:  0,
		},
	}
	positions = append(positions, startingPosition)

	eastWaypoint := &waypoint{
		direction: EAST,
		value:     10,
	}
	northWaypoint := &waypoint{
		direction: NORTH,
		value:     1,
	}
	westWaypoint := &waypoint{
		direction: WEST,
		value:     0,
	}
	southWaypoint := &waypoint{
		direction: SOUTH,
		value:     0,
	}

	for i := 0; i < len(movements); i++ {

		aMovement := movements[i]

		last := lastPosition(positions)
		var aMovementCompassDirection *compassDirection
		var anOrientationCompassDirection *compassDirection

		anOrientationCompassDirection = last.orientation
		aMovementValue := aMovement.value
		if aMovement.direction.value == "F" {
			aMovementCompassDirection = last.orientation
		} else if aMovement.direction.value == "E" {
			aMovementCompassDirection = EAST
		} else if aMovement.direction.value == "W" {
			aMovementCompassDirection = WEST
		} else if aMovement.direction.value == "S" {
			aMovementCompassDirection = SOUTH
		} else if aMovement.direction.value == "N" {
			aMovementCompassDirection = NORTH
		} else if aMovement.direction.value == "R" {

			aNewDirection := rotate(aMovement.value, anOrientationCompassDirection)
			anOrientationCompassDirection = aNewDirection

			aMovementValue = 0
			aMovementCompassDirection = anOrientationCompassDirection

		} else if aMovement.direction.value == "L" {

			aNewDirection := rotate((aMovement.value * -1), anOrientationCompassDirection)
			anOrientationCompassDirection = aNewDirection

			aMovementValue = 0
			aMovementCompassDirection = anOrientationCompassDirection

		} else {
			log.Panicf("Direction: %s not understood", aMovement.direction.value)
		}

		nextPosition := &position{
			orientation: anOrientationCompassDirection,
			bearingMovement: &bearing{
				direction: aMovementCompassDirection,
				distance:  aMovementValue,
			},
		}
		nextPosition.manhattanDistance = &manhattanDistance{}
		nextPosition.manhattanDistance.collect(nextPosition)

		positions = append(positions, nextPosition)

	}

	aManhattanDistance := &manhattanDistance{
		ew: 0,
		ns: 0,
	}
	aManhattanDistance.collectAll(positions)

	fmt.Println(fmt.Sprintf("Manhattan Distance... East / West: %d North / South: %d .... Total: %d", aManhattanDistance.ew, aManhattanDistance.ns, aManhattanDistance.distance()))

}

type manhattanDistance struct {
	ew int
	ns int
}

func (md *manhattanDistance) distance() int {
	y := Abs(md.ns)
	x := Abs(md.ew)

	return y + x
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return (x * -1)
	}
	return x
}

func (md *manhattanDistance) collectAll(positions []*position) {
	for i := 0; i < len(positions); i++ {
		aPosition := positions[i]
		md.collect(aPosition)
	}
}

func (md *manhattanDistance) collect(aPosition *position) {
	aDirection := aPosition.bearingMovement.direction
	if aDirection == EAST {
		md.ew += aPosition.bearingMovement.distance
	} else if aDirection == WEST {
		md.ew -= aPosition.bearingMovement.distance
	} else if aDirection == NORTH {
		md.ns += aPosition.bearingMovement.distance
	} else if aDirection == SOUTH {
		md.ns -= aPosition.bearingMovement.distance
	}
}

func rotate(degrees int, start *compassDirection) *compassDirection {

	aDirectionValue := start.value

	count := Abs(degrees) / 45

	degreeIndex := aDirectionValue
	for i := 0; i < count; i++ {
		if degrees > 0 {
			degreeIndex++
		} else {
			degreeIndex--
		}

		if degreeIndex == 0 {
			degreeIndex = 8
		} else if degreeIndex == 9 {
			degreeIndex = 1
		}

	}

	aDirection := directionMap[degreeIndex]

	return aDirection

}

func lastPosition(positions []*position) *position {
	return positions[len(positions)-1]
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
