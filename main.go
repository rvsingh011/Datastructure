package main

import (
	"fmt"

	"github.com/rvsingh011/Datastructure/quick_sort"
)

// Main driver program to import and run various ds programs

func main() {
	data := []int{10, 20, 1, 0, 5}
	// Sorting

	// Quick Sort
	quick_sort.QuickSort(data)
	fmt.Println(data)

	// Heap Sort

}
