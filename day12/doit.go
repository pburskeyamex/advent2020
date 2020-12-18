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

type compassValue struct {
	direction *compassDirection
	value     int
}
type ddPoint struct {
	x *compassValue
	y *compassValue
}

type movement struct {
	direction *direction
	value     int
}

type compassDirection struct {
	value       int
	description string
}

type vessel struct {
	orientation       *compassDirection
	position          *ddPoint
	manhattanDistance *manhattanDistance
	waypoint          *ddPoint
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

//func rotateLeft(degrees int, x int, y int) (int, *compassDirection, int, *compassDirection){
//	count := Abs(degrees) / 90
//	var newX, newY int
//	var xDirection, yDirection *compassDirection
//
//	for rotationCount := 0; rotationCount < count; rotationCount++ {
//		newX = y * -1
//		if newX < 0 {
//			xDirection = WEST
//		} else {
//			xDirection = EAST
//		}
//
//		newY = x
//		if newY < 0 {
//			yDirection = SOUTH
//		} else {
//			yDirection = NORTH
//		}
//		newY = Abs(newY)
//		newX = Abs(newX)
//	}
//	return newX, xDirection, newY, yDirection
//}

func rotateRight(degrees int, x int, xDirection *compassDirection, y int, yDirection *compassDirection) (int, *compassDirection, int, *compassDirection) {
	count := Abs(degrees) / 90
	var currentX, currentY, newX, newY int
	var currentXDirection, newXDirection, currentYDirection, newYDirection *compassDirection

	currentX = x
	currentXDirection = xDirection
	currentY = y
	currentYDirection = yDirection
	for rotationCount := 0; rotationCount < count; rotationCount++ {
		newX, newXDirection, newY, newYDirection = rotateRightOnce(currentX, currentXDirection, currentY, currentYDirection)
		fmt.Printf("Rotation Right: %d\n", rotationCount+1)
		fmt.Printf("%d %s %d %s", currentX, currentXDirection.description, currentY, currentYDirection.description)
		fmt.Printf(" ----> ")
		fmt.Printf("%d %s %d %s", newX, newXDirection.description, newY, newYDirection.description)
		fmt.Printf("\n")

		currentX = newX
		currentXDirection = newXDirection
		currentY = newY
		currentYDirection = newYDirection

	}
	return newX, newXDirection, newY, newYDirection
}

func rotateLeft(degrees int, x int, xDirection *compassDirection, y int, yDirection *compassDirection) (int, *compassDirection, int, *compassDirection) {
	count := Abs(degrees) / 90
	var currentX, currentY, newX, newY int
	var currentXDirection, newXDirection, currentYDirection, newYDirection *compassDirection

	currentX = x
	currentXDirection = xDirection
	currentY = y
	currentYDirection = yDirection
	for rotationCount := 0; rotationCount < count; rotationCount++ {
		newX, newXDirection, newY, newYDirection = rotateLeftOnce(currentX, currentXDirection, currentY, currentYDirection)
		fmt.Printf("Rotation Left: %d\n", rotationCount+1)
		fmt.Printf("%d %s %d %s", currentX, currentXDirection.description, currentY, currentYDirection.description)
		fmt.Printf(" ----> ")
		fmt.Printf("%d %s %d %s", newX, newXDirection.description, newY, newYDirection.description)
		fmt.Printf("\n")

		currentX = newX
		currentXDirection = newXDirection
		currentY = newY
		currentYDirection = newYDirection

	}
	return newX, newXDirection, newY, newYDirection
}

func rotateRightOnce(x int, xDirection *compassDirection, y int, yDirection *compassDirection) (int, *compassDirection, int, *compassDirection) {

	var newX, newY, realX, realY int
	var newXDirection, newYDirection *compassDirection

	realX = x
	if xDirection == WEST {
		realX = realX * -1
	}
	realY = y
	if yDirection == SOUTH {
		realY = realY * -1
	}

	newX = realY
	newY = realX * -1

	if newX < 0 {
		newXDirection = WEST
	} else {
		newXDirection = EAST
	}

	if newY < 0 {
		newYDirection = SOUTH
	} else {
		newYDirection = NORTH
	}
	newY = Abs(newY)
	newX = Abs(newX)
	//fmt.Printf("%d %s %d %s", x, xDirection.description, y, yDirection.description)
	//fmt.Printf(" --------> ")
	//fmt.Printf("%d %s %d %s", newX, newXDirection.description, newY, newYDirection.description)
	//fmt.Printf("\n")

	return newX, newXDirection, newY, newYDirection
}

func rotateLeftOnce(x int, xDirection *compassDirection, y int, yDirection *compassDirection) (int, *compassDirection, int, *compassDirection) {

	var newX, newY, realX, realY int
	var newXDirection, newYDirection *compassDirection

	realX = x
	if xDirection == WEST {
		realX = realX * -1
	}
	realY = y
	if yDirection == SOUTH {
		realY = realY * -1
	}

	newY = realX
	newX = realY * -1

	if newX < 0 {
		newXDirection = WEST
	} else {
		newXDirection = EAST
	}

	if newY < 0 {
		newYDirection = SOUTH
	} else {
		newYDirection = NORTH
	}
	newY = Abs(newY)
	newX = Abs(newX)
	//fmt.Printf("%d %s %d %s", x, xDirection.description, y, yDirection.description)
	//fmt.Printf(" --------> ")
	//fmt.Printf("%d %s %d %s", newX, newXDirection.description, newY, newYDirection.description)
	//fmt.Printf("\n")

	return newX, newXDirection, newY, newYDirection
}

func addX(start int, startDirection *compassDirection, value int, valueDirection *compassDirection) (int, *compassDirection) {

	adjustedStart := start
	if startDirection == WEST {
		adjustedStart *= -1
	}

	adjustedValue := value
	if valueDirection == WEST {
		adjustedValue *= -1
	}

	newValue := adjustedStart + adjustedValue
	direction := EAST
	if newValue < 0 {
		direction = WEST
		newValue = Abs(newValue)
	}
	return newValue, direction

}

func addY(start int, startDirection *compassDirection, value int, valueDirection *compassDirection) (int, *compassDirection) {

	adjustedStart := start
	if startDirection == SOUTH {
		adjustedStart *= -1
	}

	adjustedValue := value
	if valueDirection == SOUTH {
		adjustedValue *= -1
	}

	newValue := adjustedStart + adjustedValue
	direction := NORTH
	if newValue < 0 {
		direction = SOUTH
		newValue = Abs(newValue)
	}
	return newValue, direction

}

func main() {

	//newX, xDirection, newY, yDirection := rotateRight(360, 10, EAST, 4, NORTH)
	//log.Print(newX)
	//log.Print(xDirection)
	//log.Print(newY)
	//log.Print(yDirection)
	//
	//newX, xDirection, newY, yDirection = rotateLeft(360, 10, EAST, 4, NORTH)
	//log.Print(newX)
	//log.Print(xDirection)
	//log.Print(newY)
	//log.Print(yDirection)

	//proof("day_11_sample_data_2.txt")

	data := Parse("day_12_data.txt")
	//data := Parse("day_12_data.txt")

	movements := make([]*movement, 0)

	for _, aString := range data {
		aMovement := parse(aString)
		movements = append(movements, aMovement)
	}

	positions := make([]*vessel, 0)

	startingPosition := &vessel{
		orientation: EAST,
		position: &ddPoint{
			x: &compassValue{
				direction: EAST,
				value:     0,
			},
			y: &compassValue{
				direction: NORTH,
				value:     0,
			},
		},
		waypoint: &ddPoint{
			x: &compassValue{
				direction: EAST,
				value:     10,
			},
			y: &compassValue{
				direction: NORTH,
				value:     1,
			},
		},
	}
	positions = append(positions, startingPosition)

	for i := 0; i < len(movements); i++ {

		aMovement := movements[i]

		last := lastPosition(positions)

		nextPosition := &vessel{
			orientation: last.orientation,
			position: &ddPoint{
				x: &compassValue{
					direction: last.position.x.direction,
					value:     last.position.x.value,
				},
				y: &compassValue{
					direction: last.position.y.direction,
					value:     last.position.y.value,
				},
			},
			waypoint: &ddPoint{
				x: &compassValue{
					direction: last.waypoint.x.direction,
					value:     last.waypoint.x.value,
				},
				y: &compassValue{
					direction: last.waypoint.y.direction,
					value:     last.waypoint.y.value,
				},
			},
		}
		positions = append(positions, startingPosition)

		if aMovement.direction.value == "F" {

			/*
				plot where we are starting....
			*/
			startingX := nextPosition.position.x.value
			if nextPosition.position.x.direction == WEST {
				startingX = startingX * -1
			}
			startingY := nextPosition.position.y.value
			if nextPosition.position.y.direction == SOUTH {
				startingY = startingY * -1
			}

			/*
				find the adjustment to the position by moving to the waypoint n times.....
			*/
			wayPoint := nextPosition.waypoint
			wayPointAdjustmentX := 0
			wayPointAdjustmentY := 0
			for i := 0; i < aMovement.value; i++ {
				wayPointAdjustmentX += (wayPoint.x.value)
				wayPointAdjustmentY += (wayPoint.y.value)
			}

			if wayPoint.x.direction == WEST {
				wayPointAdjustmentX = wayPointAdjustmentX * -1
			}
			if wayPoint.y.direction == SOUTH {
				wayPointAdjustmentY = wayPointAdjustmentY * -1
			}

			nextPosition.position.x.value = (startingX + wayPointAdjustmentX)
			nextPosition.position.y.value = (startingY + wayPointAdjustmentY)

			if nextPosition.position.x.value < 0 {
				nextPosition.position.x.value *= -1
				nextPosition.position.x.direction = WEST
			} else {
				nextPosition.position.x.direction = EAST
			}

			if nextPosition.position.y.value < 0 {
				nextPosition.position.y.value *= -1
				nextPosition.position.y.direction = SOUTH
			} else {
				nextPosition.position.y.direction = NORTH
			}

		} else if aMovement.direction.value == "E" {
			adjustedValue, adjustedDirection := addX(nextPosition.waypoint.x.value, nextPosition.waypoint.x.direction, aMovement.value, EAST)
			nextPosition.waypoint.x.value = adjustedValue
			nextPosition.waypoint.x.direction = adjustedDirection
		} else if aMovement.direction.value == "W" {
			adjustedValue, adjustedDirection := addX(nextPosition.waypoint.x.value, nextPosition.waypoint.x.direction, aMovement.value, WEST)
			nextPosition.waypoint.x.value = adjustedValue
			nextPosition.waypoint.x.direction = adjustedDirection
		} else if aMovement.direction.value == "S" {
			adjustedValue, adjustedDirection := addY(nextPosition.waypoint.y.value, nextPosition.waypoint.y.direction, aMovement.value, SOUTH)
			nextPosition.waypoint.y.value = adjustedValue
			nextPosition.waypoint.y.direction = adjustedDirection
		} else if aMovement.direction.value == "N" {
			adjustedValue, adjustedDirection := addY(nextPosition.waypoint.y.value, nextPosition.waypoint.y.direction, aMovement.value, NORTH)
			nextPosition.waypoint.y.value = adjustedValue
			nextPosition.waypoint.y.direction = adjustedDirection
		} else if aMovement.direction.value == "R" {

			newX, xDirection, newY, yDirection := rotateRight(aMovement.value, nextPosition.waypoint.x.value, nextPosition.waypoint.x.direction, nextPosition.waypoint.y.value, nextPosition.waypoint.y.direction)
			nextPosition.waypoint.x.direction = xDirection
			nextPosition.waypoint.x.value = newX

			nextPosition.waypoint.y.direction = yDirection
			nextPosition.waypoint.y.value = newY

		} else if aMovement.direction.value == "L" {

			newX, xDirection, newY, yDirection := rotateLeft(aMovement.value, nextPosition.waypoint.x.value, nextPosition.waypoint.x.direction, nextPosition.waypoint.y.value, nextPosition.waypoint.y.direction)
			nextPosition.waypoint.x.direction = xDirection
			nextPosition.waypoint.x.value = newX

			nextPosition.waypoint.y.direction = yDirection
			nextPosition.waypoint.y.value = newY

		} else {
			log.Panicf("Direction: %s not understood", aMovement.direction.value)
		}

		nextPosition.manhattanDistance = &manhattanDistance{}
		nextPosition.manhattanDistance.collect(nextPosition)

		positions = append(positions, nextPosition)

	}

	aManhattanDistance := &manhattanDistance{
		x: 0,
		y: 0,
	}
	aManhattanDistance.collectAll(positions)

	fmt.Println(fmt.Sprintf("Manhattan Distance... East / West: %d North / South: %d .... Total: %d", aManhattanDistance.x, aManhattanDistance.y, aManhattanDistance.distance()))

}

type manhattanDistance struct {
	x int
	y int
}

func (md *manhattanDistance) distance() int {
	y := Abs(md.y)
	x := Abs(md.x)

	return y + x
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return (x * -1)
	}
	return x
}

func (md *manhattanDistance) collectAll(positions []*vessel) {
	for i := 0; i < len(positions); i++ {
		aPosition := positions[i]
		md.collect(aPosition)
	}
}

func (md *manhattanDistance) collect(aPosition *vessel) {

	md.y = aPosition.position.y.value
	md.x = aPosition.position.x.value

}

func rotateWaypointRight(degrees int, waypoint **ddPoint) {

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

func lastPosition(positions []*vessel) *vessel {
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
