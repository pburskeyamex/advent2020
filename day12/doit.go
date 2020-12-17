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

func main() {

	//proof("day_11_sample_data_2.txt")

	data := Parse("day_12_sample_data.txt")
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
				direction: nil,
				value:     0,
			},
			y: &compassValue{
				direction: nil,
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

			/*
				fun rotateLeft(): Point2D =
				    Point2D(x = y * -1, y = x)

				fun rotateRight(): Point2D =
				    Point2D(x = y, y = x * -1)
			*/

			x := nextPosition.waypoint.y.value
			if x < 0 {
				nextPosition.waypoint.x.direction = WEST
			} else {
				nextPosition.waypoint.x.direction = EAST
			}

			y := nextPosition.waypoint.x.value * -1
			if y < 0 {
				nextPosition.waypoint.y.direction = SOUTH
			} else {
				nextPosition.waypoint.y.direction = NORTH
			}
			nextPosition.waypoint.y.value = Abs(y)
			nextPosition.waypoint.x.value = Abs(x)

			//
			//aNewDirection := rotate(aMovement.value, anOrientationCompassDirection)
			//anOrientationCompassDirection = aNewDirection
			//
			//aMovementValue = 0
			//aMovementCompassDirection = anOrientationCompassDirection

			//var radians float64
			/*x rads in degrees - > x*180/pi
			x degrees in rads -> x*pi/180

			*/

			//aFloat := 180 / math.Pi
			//
			//radians = float64(aMovement.value) * aFloat
			//
			//rotatedX := (math.Cos(radians) * float64(last.waypoint.x.value)) - (math.Sin(radians) * float64(last.waypoint.y.value))
			//rotatedY := (math.Cos(radians) * float64(last.waypoint.x.value)) + (math.Sin(radians) * float64(last.waypoint.y.value))
			//log.Println(rotatedX)
			//log.Println(rotatedY)

		} else if aMovement.direction.value == "L" {

			aNewDirection := rotate((aMovement.value * -1), anOrientationCompassDirection)
			anOrientationCompassDirection = aNewDirection

			aMovementValue = 0
			aMovementCompassDirection = anOrientationCompassDirection

			/*
				fun rotateLeft(): Point2D =
				    Point2D(x = y * -1, y = x)

				fun rotateRight(): Point2D =
				    Point2D(x = y, y = x * -1)
			*/

			x := nextPosition.waypoint.y.value * -1
			if x < 0 {
				nextPosition.waypoint.x.direction = WEST
			} else {
				nextPosition.waypoint.x.direction = EAST
			}

			y := nextPosition.waypoint.x.value
			if y < 0 {
				nextPosition.waypoint.y.direction = SOUTH
			} else {
				nextPosition.waypoint.y.direction = NORTH
			}
			nextPosition.waypoint.y.value = Abs(y)
			nextPosition.waypoint.x.value = Abs(x)

		} else {
			log.Panicf("Direction: %s not understood", aMovement.direction.value)
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
