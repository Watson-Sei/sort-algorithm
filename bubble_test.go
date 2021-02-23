package main

import (
	"testing"
)

func Sum(numbers []int) (sum int) {
	for i := 0; i < len(numbers); i ++ {
		sum += numbers[i]
	}
	return
}

func TestBubble(t *testing.T) {
	expect := []int{2, 7, 12, 34, 99}
	actual := bubble([]int{99, 12, 2, 34, 7})

	got := Sum(actual)
	want := Sum(expect)

	// 配列と配列の比較ができないので、Sumの合計を数値で出し比較する
	if want != got {
		t.Errorf("expect: %v, actual: %v", expect, actual)
	}
}