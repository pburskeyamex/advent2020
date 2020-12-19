package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bus struct {
	id    int
	times []int
}

func (b *bus) closesTimeDifferenceGreaterThan(x int) int {

	toFind := 0
	for i := 0; toFind <= x && i < len(b.times); i++ {
		value := b.times[i]
		if value >= x {
			toFind = value
		}
	}
	return toFind
}

func main() {

	data := Parse("day_13_data.txt")
	data1 := data[0]
	data2 := data[1]

	goalTimeToLeave, _ := strconv.Atoi(data1)
	log.Printf("%d\n", goalTimeToLeave)
	busses := make([]*bus, 0)

	for _, aString := range strings.Split(data2, ",") {
		if aString != "x" {
			id, _ := strconv.Atoi(aString)
			bus := &bus{
				id:    id,
				times: make([]int, 0),
			}
			bus.times = append(bus.times, 0)
			bus.times = append(bus.times, id)
			busses = append(busses, bus)
		}
	}

	for i := 0; i < len(busses); i++ {
		bus := busses[i]

		for x := (goalTimeToLeave / bus.id); x*bus.id < (goalTimeToLeave + bus.id); x++ {
			aTime := x * (bus.id)
			bus.times = append(bus.times, aTime)
		}
	}

	var closestBus *bus

	for i := 0; i < len(busses); i++ {
		bus := busses[i]

		difference := bus.closesTimeDifferenceGreaterThan(goalTimeToLeave)
		if closestBus == nil || closestBus.closesTimeDifferenceGreaterThan(goalTimeToLeave) > difference {
			closestBus = bus
		}

	}

	howLongToWait := closestBus.closesTimeDifferenceGreaterThan(goalTimeToLeave) - goalTimeToLeave

	log.Printf("Bus ID: %d Time to Leave: %d. Waiting: %d minutes... Magic number: %d", closestBus.id, 0, howLongToWait, howLongToWait*closestBus.id)
	//data := Parse("day_12_data.txt")

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
