package _4_sorts

import "testing"

func TestCountingSort(t *testing.T) {
	arr := []int{5, 4}
	if err := CountingSort(arr, len(arr)); err != nil {
		t.Fatal(err.Error())
	}
	t.Log(arr)

	arr = []int{-5, 4, 3, -2, 1}
	if err := CountingSort(arr, len(arr)); err != nil {
		t.Fatal(err.Error())
	}
	t.Log(arr)
}
