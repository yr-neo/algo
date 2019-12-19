// 计数排序

package _4_sorts

import (
	"algo_yr/go/00_lib"
)

func CountingSort(a []int, n int) {
	if n <= 1 {
		return
	}

	max, min := _0_lib.GetMaxMin(a)
	limit := max - min

	c := make([]int, limit+1)
	for i := range a {
		a[i] -= min
		c[a[i]]++
	}
	for i := 1; i <= limit; i++ {
		c[i] += c[i-1]
	}

	r := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		index := c[a[i]] - 1
		r[index] = a[i] + min
		c[a[i]]--
	}

	copy(a, r)
}
