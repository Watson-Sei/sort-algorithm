package main

import "testing"

func SelectionSum(numbers []int) (sum int) {
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return
}

func TestSelection(t *testing.T) {
	expect := []int{2, 7, 12, 34, 99}
	actual := selection([]int{99, 12, 2, 34, 7})

	got := SelectionSum(actual)
	want := SelectionSum(expect)

	// 配列と配列の比較ができないので、Sumの合計を数値で出し比較する
	if want != got {
		t.Errorf("expect: %v, actual: %v", expect, actual)
	}
}
