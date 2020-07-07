package gosort

func sort(myArray *[]int, leftIndex int, middleIndex int, rightIndex int) {
	sortedArray := make([]int, rightIndex-leftIndex+1)

	i := leftIndex
	j := middleIndex
	k := 0

	for i < middleIndex || j <= rightIndex {
		switch {

		// No more value left; copy remaining values from right array part
		case i == middleIndex:
			sortedArray[k] = (*myArray)[j]
			j++

		// No more value right; copy remaining values from left array part
		case j > rightIndex:
			sortedArray[k] = (*myArray)[i]
			i++

		// Still value on the right and left parts: choose the smallest
		default:
			if (*myArray)[i] < (*myArray)[j] {
				sortedArray[k] = (*myArray)[i]
				i++
			} else {
				sortedArray[k] = (*myArray)[j]
				j++
			}
		}
		k++
	}
	for index, value := range sortedArray {
		(*myArray)[leftIndex+index] = value
	}
}

func merge(myArray *[]int, leftIndex int, rightIndex int) {
	if leftIndex != rightIndex {
		var middleIndex int = (rightIndex + leftIndex) / 2
		merge(myArray, leftIndex, middleIndex)
		merge(myArray, middleIndex+1, rightIndex)
		sort(myArray, leftIndex, middleIndex+1, rightIndex)
	}
}

// MergeSort sorts an array using merge sort algorithm
func MergeSort(myArray *[]int) {
	if len(*myArray) == 0 {
		return
	}
	merge(myArray, 0, len(*myArray)-1)
}
