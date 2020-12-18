package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	strings2 "strings"
)

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
