package main

import (
	"fmt"
	"log"
	"math"
)

func toRadians(degrees int) float64 {
	/*
		to radians
	*/
	aFloat := math.Pi / 180

	radians := float64(degrees) * aFloat

	return radians
}

func toDegrees(radians int) float64 {
	/*
		to radians
	*/
	aFloat := 180 / math.Pi

	degrees := float64(radians) * aFloat

	return degrees
}

func main() {

	x := 4
	y := 10
	degrees := 90

	radians := toRadians(degrees)
	log.Println(fmt.Sprintf("Radians: %f", radians))

	rotatedX := (math.Cos(radians) * float64(x)) - (math.Sin(radians) * float64(y))
	rotatedY := (math.Cos(radians) * float64(x)) + (math.Sin(radians) * float64(y))
	log.Println(fmt.Sprintf("X: %f", rotatedX))
	log.Println(fmt.Sprintf("Y: %f", rotatedY))

}
