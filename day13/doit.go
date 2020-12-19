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
	id       int64
	position int
	times    []int64
}

type schedule struct {
	time   int64
	busses []bool
}

func (b *bus) doesBusLeaveAt(time int64) bool {
	var result int64
	if b.id == 0 {
		result = 0
	} else {
		result = time % b.id
	}

	return (result == 0)
}

func (b *bus) completeSchedule() {

	var nextTime int64
	nextTime = 0
	for i := 0; i < 1000; i++ {
		nextTime = nextTime + int64(b.id)
		b.times = append(b.times, nextTime)
	}
}

func (b *bus) closesTimeDifferenceGreaterThan(x int64) int64 {

	toFind := int64(0)
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

	var goalTimeToLeave int64
	goalTimeToLeave, _ = strconv.ParseInt(data1, 10, 64)
	log.Printf("%d\n", goalTimeToLeave)
	busses := make([]*bus, 0)

	for i, aString := range strings.Split(data2, ",") {
		var id int64
		if aString != "x" {
			id, _ = strconv.ParseInt(aString, 10, 64)
		}
		bus := &bus{
			id:       id,
			position: i,
			times:    make([]int64, 0),
		}
		busses = append(busses, bus)

	}

	part1(busses, goalTimeToLeave)

	var seed int64
	seed = 100000000000000
	//seed = 1
	part2(busses, seed)

	//table := tablewriter.NewWriter ( os.Stdout)
	//table.SetHeader([]string{"Date", "ActiveCases", "TotalPositiveCases", "ProbableCases", "ResolvedCases"})
	//table.SetBorder(false)       // Set Border to false
	//table.AppendBulk(busses) // Add Bulk Data
	//table.Render()

}

func compress(busses []*bus) []*bus {

	scrubbedBusses := make([]*bus, 0)
	//var lastKeep *bus
	for i := 0; i < len(busses); i++ {
		bus := busses[i]
		//
		//keep := lastKeep == nil
		//keep = keep || (bus.id != 0)
		//keep = keep || (lastKeep.id != 0 && bus.id == 0)
		//if keep {
		scrubbedBusses = append(scrubbedBusses, bus)
		//lastKeep = bus
		//}
	}

	for i := 0; i < len(scrubbedBusses); i++ {
		bus := scrubbedBusses[i]
		bus.position = i + 1

	}

	return scrubbedBusses

}

func part2(busses []*bus, seed int64) {

	busses = compress(busses)
	var schedules []*schedule
	var time int64

	schedules = make([]*schedule, len(busses))
	for x := 0; x < len(busses); x++ {
		aSchedule := &schedule{
			time:   time,
			busses: make([]bool, len(busses)),
		}
		schedules[x] = aSchedule

		for i := 0; i < len(busses); i++ {

			aSchedule.busses[i] = false
		}
	}

	/*
		advance until bus 1 wants to leave....
	*/

	bus1 := busses[0]
	found := false
	iterationCoount := 0
	for time = seed; !found; time++ {
		iterationCoount++
		if bus1.doesBusLeaveAt(time) {
			if (iterationCoount % 1000000) == 0 {
				log.Println(fmt.Sprintf("Iteration: %v", time))

			}
			for x := 0; x < len(busses); x++ {
				aSchedule := schedules[x]
				aSchedule.time = time

				for i := 0; i < len(busses); i++ {
					bus := busses[i]
					aSchedule.busses[i] = bus.doesBusLeaveAt(time)
				}
				time++
			}

			/*
				validate whether the schedule has a bus leaving at its position on each time that matches the bus.
			*/
			scheduleLadder := true
			for x := 0; scheduleLadder && x < len(schedules); x++ {
				schedule := schedules[x]
				busLeaving := schedule.busses[x]
				scheduleLadder = scheduleLadder && busLeaving
			}

			found = scheduleLadder
		}

	}

	if found {

		startingTime := schedules[0].time
		log.Println("Earliest Time: ", startingTime)

	}

}

func part1(busses []*bus, goalTimeToLeave int64) {

	scrubbedBusses := make([]*bus, 0)
	for i := 0; i < len(busses); i++ {
		bus := busses[i]
		if bus.id == 0 {

		} else {
			scrubbedBusses = append(scrubbedBusses, bus)
		}

	}

	for i := 0; i < len(scrubbedBusses); i++ {
		bus := scrubbedBusses[i]

		anInt64 := int64(bus.id)
		for x := (goalTimeToLeave / anInt64); x*anInt64 < (goalTimeToLeave + anInt64); x++ {
			aTime := x * (anInt64)
			bus.times = append(bus.times, int64(aTime))
		}
	}

	var closestBus *bus

	for i := 0; i < len(scrubbedBusses); i++ {
		bus := scrubbedBusses[i]

		difference := bus.closesTimeDifferenceGreaterThan(goalTimeToLeave)
		if closestBus == nil || closestBus.closesTimeDifferenceGreaterThan(goalTimeToLeave) > difference {
			closestBus = bus
		}

	}

	howLongToWait := closestBus.closesTimeDifferenceGreaterThan(goalTimeToLeave) - goalTimeToLeave
	log.Printf("Bus ID: %d Time to Leave: %d. Waiting: %d minutes... Magic number: %d", closestBus.id, 0, howLongToWait, howLongToWait*int64(closestBus.id))

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

/*
fun solvePart2(): Long {
    var stepSize = indexedBusses.first().bus
    var time = 0L
    indexedBusses.drop(1).forEach { (offset, bus) ->
        while ((time + offset) % bus != 0L) {
            time += stepSize
        }
        stepSize *= bus // New Ratio!
    }
    return time
}

*/
