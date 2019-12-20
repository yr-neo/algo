package _0_lib

import (
	"errors"
	"math/rand"
)

func CreateRandArr(length, ranges int) ([]int, error) {
	if length <= 0 || ranges <= 0 {
		return nil, errors.New("length <= 0 || ranges <= 0")
	}

	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(ranges)
	}
	return arr, nil
}

func GetMaxMin(arr []int) (max int, min int, err error) {
	if len(arr) <= 0 {
		return 0, 0, errors.New("len(arr) <= 0")
	}

	max, min = arr[0], arr[0]
	for i := range arr {
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
	}
	return max, min, nil
}
