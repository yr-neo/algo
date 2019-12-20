package _4_sorts

import "testing"

func TestRadixSort(t *testing.T) {
	arr := []int{1, 34, 546, 23, 465}
	if err := RadixSort(arr); err != nil {
		t.Fatal(err.Error())
	}
	t.Log(arr)

	arr = []int{1, 34, -546, -23, 465, 34, -1234}
	if err := RadixSort(arr); err != nil {
		t.Fatal(err.Error())
	}
	t.Log(arr)
}
