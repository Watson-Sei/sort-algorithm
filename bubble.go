package main

// バブルソート
func bubble(array []int) []int {
	for i := 0; i < len(array) - 1; i++ {
		for j := 0; j < len(array) - i - 1; j++ {
			if array[j] > array[j + 1] {
				// 配列を入れ替える
				array[j], array[j + 1] = array[j + 1], array[j]
			}
		}
	}
	return array
}