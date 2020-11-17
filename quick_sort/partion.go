package quick_sort

func swap(arr []int, firstIndex, secondIndex int) {
	var temp int
	temp = arr[firstIndex]
	arr[firstIndex] = arr[secondIndex]
	arr[secondIndex] = temp
}

func Partion(list []int) int {
	pivot := list[len(list)-1]
	i, scanner := -1, 0
	for scanner < len(list) {
		if list[scanner] < pivot {
			i++
			swap(list, i, scanner)
		}
		scanner++
	}
	i++
	swap(list, i, scanner-1)
	return i
}

func QuickSort(data []int) {
	if len(data) < 2 {
		return
	}
	changeIndex := Partion(data)
	QuickSort(data[:changeIndex])
	QuickSort(data[changeIndex+1:])
}
