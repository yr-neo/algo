// 基数排序

package _4_sorts

import (
	_0_lib "algo_yr/go/00_lib"
	"errors"
)

// 十进制
const base = 10

// 获取最大元素的位数
func getMaxBit(max int) int {
	digit := 1
	for max >= base {
		max /= base
		digit++
	}
	return digit
}

func RadixSort(a []int) error {
	length := len(a)
	if length <= 1 {
		return errors.New("length <= 1")
	}

	max, min, err := _0_lib.GetMaxMin(a)
	if err != nil {
		return err
	}

	radix := 1
	digit := getMaxBit(max - min)
	tmp := make([]int, length)

	var i, j, k int
	var count []int

	for i = range a {
		a[i] -= min
	}

	for i = 0; i < digit; i++ {
		count = make([]int, 10)
		for j = range a {
			k = a[j] / radix % base
			count[k]++
		}
		for j = 1; j < 10; j++ {
			count[j] += count[j-1]
		}
		for j = length - 1; j >= 0; j-- {
			k = a[j] / radix % base
			tmp[count[k]-1] = a[j]
			count[k]--
		}
		copy(a, tmp)
		radix *= base
	}

	for i = range a {
		a[i] += min
	}
	return nil
}
