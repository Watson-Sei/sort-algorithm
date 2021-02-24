package main

func selection(array []int) []int {
	for i := 0; i < len(array) - 1; i++ {
		minindex := i
		for j := i + 1; j < len(array); j++ {
			if array[minindex] > array[j] {
				minindex = j
			}
		}
		array[i],array[minindex] = array[minindex],array[i]
	}
	return array
}