package _4_sorts

import "testing"

func TestRadixSort(t *testing.T) {
	arr := []int{1, 34, 546, 23, 465}
	RadixSort(arr)
	t.Log(arr)

	arr = []int{1, 34, -546, -23, 465, 34, -1234}
	RadixSort(arr)
	t.Log(arr)
}
