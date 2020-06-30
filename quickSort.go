package main

func inverse(myArray *[]int, pos1 int, pos2 int) {
	el := (*myArray)[pos1]
	(*myArray)[pos1] = (*myArray)[pos2]
	(*myArray)[pos2] = el
}

func sort(myArray *[]int, begin int, end int) {
	// fmt.Printf("exec process on %v from %v to %v\n", *myArray, begin, end)
	initialPivotIndex := begin
	leftMarkIndex := initialPivotIndex + 1
	rightMarkIndex := end

	if leftMarkIndex > rightMarkIndex {
		return
	}

	pivotEl := (*myArray)[initialPivotIndex]

	for leftMarkIndex <= rightMarkIndex && pivotEl >= (*myArray)[leftMarkIndex] {
		leftMarkIndex++
	}

	for leftMarkIndex <= rightMarkIndex && pivotEl <= (*myArray)[rightMarkIndex] {
		rightMarkIndex--
	}

	if leftMarkIndex <= rightMarkIndex {
		inverse(myArray, rightMarkIndex, leftMarkIndex)
	} else {
		inverse(myArray, initialPivotIndex, rightMarkIndex)
	}

	// Apply same process on left array
	sort(myArray, begin, rightMarkIndex-1)

	// Apply same process on right array
	sort(myArray, rightMarkIndex+1, end)
}

// QuickSort sorts an array using quick sort algorithm
// Follow https://runestone.academy/runestone/books/published/pythonds/SortSearch/TheQuickSort.html
func QuickSort(myArray *[]int) {
	sort(myArray, 0, len(*myArray)-1)
}
