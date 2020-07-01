package main

import "fmt"

func main() {
	fmt.Println("start")
	entry := []int{42, 1, 2, 3, 2, 4, 7, 18, 1, 26, 37, 49, 5, 34}
	QuickSort(&entry)
	fmt.Println("end", entry)
}
