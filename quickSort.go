package main

import (
	"sync"
)

func inverse(myArray *[]int, pos1 int, pos2 int) {
	el := (*myArray)[pos1]
	(*myArray)[pos1] = (*myArray)[pos2]
	(*myArray)[pos2] = el
}

func sort(parentWg *sync.WaitGroup, myArray *[]int, begin int, end int) {

	if parentWg != nil {
		defer parentWg.Done()
	}

	initialPivotIndex := begin
	leftMarkIndex := initialPivotIndex + 1
	rightMarkIndex := end

	if leftMarkIndex > rightMarkIndex {
		return
	}

	pivotValue := (*myArray)[initialPivotIndex]
	// fmt.Printf("exec process on %v from %v to %v with pivot [%v]=%v\n", *myArray, begin, end, initialPivotIndex, pivotValue)

	for leftMarkIndex <= rightMarkIndex {
		for leftMarkIndex <= rightMarkIndex && pivotValue >= (*myArray)[leftMarkIndex] {
			leftMarkIndex++
		}

		for leftMarkIndex <= rightMarkIndex && pivotValue <= (*myArray)[rightMarkIndex] {
			rightMarkIndex--
		}

		if leftMarkIndex <= rightMarkIndex {
			inverse(myArray, rightMarkIndex, leftMarkIndex)
		}
	}

	// fmt.Printf(" inverse pivot [%v]=%v with [%v]=%v\n", initialPivotIndex, (*myArray)[initialPivotIndex], rightMarkIndex, (*myArray)[rightMarkIndex])
	inverse(myArray, initialPivotIndex, rightMarkIndex)

	// fmt.Printf("next sorts %v:%v and %v:%v\n", begin, rightMarkIndex-1, rightMarkIndex+1, end)
	var wg sync.WaitGroup
	wg.Add(2)

	// Apply same process on left array
	go sort(&wg, myArray, begin, rightMarkIndex-1)
	// Apply same process on right array
	go sort(&wg, myArray, rightMarkIndex+1, end)

	// Wait each Go routine end
	wg.Wait()
}

// QuickSort sorts an array using quick sort algorithm
func QuickSort(myArray *[]int) {
	sort(nil, myArray, 0, len(*myArray)-1)
}
