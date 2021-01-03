package main

import (
	"bufio"
	"fmt"
	"os"
)

type point3d struct {
	x int
	y int
	z int
}

func (self *point3d) relativePoints(depth int) []*point3d {

	relativePoints := make([]*point3d, 0)

	startingZ := self.z - depth
	endingZ := self.z + depth
	startingX := self.x - depth
	endingX := self.x + depth
	startingY := self.y - depth
	endingY := self.y + depth

	for z := startingZ; z <= endingZ; z++ {

		for x := startingX; x <= endingX; x++ {

			for y := startingY; y <= endingY; y++ {

				aPoint := &point3d{
					x: x,
					y: y,
					z: z,
				}
				relativePoints = append(relativePoints, aPoint)

			}

		}

	}

	return relativePoints

}

type cubes struct {
	stateMap map[*point3d]*cubeState
}

func (self *cubes) cycle() {

	var aCubeState *cubeState
	for _, aCubeState = range self.stateMap {
		aCubeState.cycleWithin(self)
	}
}

type cubeState struct {
	active bool
	point  *point3d
}

func (self *cubeState) cycleWithin(cubes *cubes) {
	if self.active {
		neighborActiveCount := self.countActiveNeighbors(cubes)
		if neighborActiveCount == 2 || neighborActiveCount == 3 {

		} else {
			self.active = false
		}
	} else {
		neighborActiveCount := self.countActiveNeighbors(cubes)
		if neighborActiveCount == 3 {
			self.active = true
		}
	}
}

func (self *cubeState) countActiveNeighbors(cubes *cubes) int {
	activeCount := 0
	var virtualCubes [][][]*cubeState
	virtualCubes = cubes.plot(self)
	if len(virtualCubes) > 0 {

	}
	return activeCount
}

func main() {

	expectation := 112
	if result := part1("day_17_sample_data.txt"); result != expectation {
		panic(fmt.Sprintf("Expected: %d", expectation))
	}

}

func part1(aFileName string) (result int) {

	var initialGraphState []*cubeState
	initialGraphState = Parse(aFileName)

	cubes := &cubes{
		stateMap: make(map[*point3d]*cubeState, 0),
	}

	for _, aCube := range initialGraphState {
		cubes.stateMap[aCube.point] = aCube
	}

	cubes.cycle()

	if len(initialGraphState) > 0 {

	}

	return result
}

func (self *cubes) getCubeStateAtPoint(point *point3d) *cubeState {

	var aCubeState *cubeState
	var ok bool
	if aCubeState, ok = self.stateMap[point]; !ok {

		aCubeState = &cubeState{
			active: false,
			point:  point,
		}
		self.stateMap[point] = aCubeState
	}

	return aCubeState
}

func (self *cubes) plot(aCube *cubeState) [][][]*cubeState {

	relativePoints := aCube.point.relativePoints(1)

	depth := 3

	graph := make([][][]*cubeState, 0)

	var z, x, y int
	startingZ := 0
	endingZ := z + depth
	startingX := 0
	endingX := x + depth
	startingY := 0
	endingY := y + depth

	for z := startingZ; z < endingZ; z++ {
		graph = append(graph, make([][]*cubeState, 0))
		for x := startingX; x < endingX; x++ {
			graph[z] = append(graph[z], make([]*cubeState, 0))
			for y := startingY; y < endingY; y++ {
				graph[z][x] = append(graph[z][x], nil)
			}
		}
	}

	for _, aPoint := range relativePoints {
		currentCubeState := self.getCubeStateAtPoint(aPoint)
		if currentCubeState != nil {
			z := aPoint.z + 1
			x := aPoint.x + 1
			y := aPoint.y + 1
			if currentCubeState.active {
				graph[z][x][y] = currentCubeState
			} else {
				graph[z][x][y] = currentCubeState
			}

		}

	}

	return graph
}

func Parse(aFilePart string) (initialGraphState []*cubeState) {
	filename := fmt.Sprintf("data/%s", aFilePart)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	initialGraphState = make([]*cubeState, 0)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for i := 0; fileScanner.Scan(); i++ {
		aString := fileScanner.Text()
		runes := []rune(aString)

		for j, aRune := range runes {

			active := (aRune == '#')
			aPoint := &point3d{
				x: j,
				y: i,
				z: 0,
			}

			aCube := &cubeState{
				active: active,
				point:  aPoint,
			}

			initialGraphState = append(initialGraphState, aCube)

		}

	}

	file.Close()

	return initialGraphState
}
