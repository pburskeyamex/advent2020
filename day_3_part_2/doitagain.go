package main

import "fmt"
import . "../day3"

type slope struct{
	right int
	down int
	treeCount uint
	missCount uint
}


func main() {

	var runningCount uint

	slope := &slope{
		right: 1,
		down:  1,
	}
	slope.treeCount, slope.missCount = Algorithm(slope.right,slope.down, 5)
	runningCount = slope.treeCount
	fmt.Printf("Tree count: %d Miss count: %d Running Count %d\n", slope.treeCount, slope.missCount, runningCount)

	slope.right = 3
	slope.down = 1
	slope.treeCount, slope.missCount = Algorithm(slope.right,slope.down, 5)
	runningCount = runningCount * slope.treeCount
	fmt.Printf("Tree count: %d Miss count: %d Running Count %d\n", slope.treeCount, slope.missCount, runningCount)

	slope.right = 5
	slope.down = 1
	slope.treeCount, slope.missCount = Algorithm(slope.right,slope.down, 5)
	runningCount = runningCount * slope.treeCount
	fmt.Printf("Tree count: %d Miss count: %d Running Count %d\n", slope.treeCount, slope.missCount, runningCount)

	slope.right = 7
	slope.down = 1
	slope.treeCount, slope.missCount = Algorithm(slope.right,slope.down, 5)
	runningCount = runningCount * slope.treeCount
	fmt.Printf("Tree count: %d Miss count: %d Running Count %d\n", slope.treeCount, slope.missCount, runningCount)


	slope.right = 1
	slope.down = 2
	slope.treeCount, slope.missCount = Algorithm(slope.right,slope.down, 5)
	runningCount = runningCount * slope.treeCount
	fmt.Printf("Tree count: %d Miss count: %d Running Count %d\n", slope.treeCount, slope.missCount, runningCount)

}
