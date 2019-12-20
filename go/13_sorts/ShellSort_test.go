package _3_sorts

import (
	_0_lib "algo_yr/go/00_lib"
	"testing"
	"time"
)

func TestShellSort(t *testing.T) {
	var err error
	var arr []int
	var start time.Time

	arr = []int{2, 7, 11, -5, 23, 165, -99}
	ShellSort1(arr)
	t.Log(arr)

	if arr, err = _0_lib.CreateRandArr(10000000, 1000000); err != nil {
		t.Fatal(err.Error())
	}
	start = time.Now()
	ShellSort1(arr)
	t.Log("ShellSort1 cost:", time.Since(start))

	if arr, err = _0_lib.CreateRandArr(10000000, 1000000); err != nil {
		t.Fatal(err.Error())
	}
	start = time.Now()
	ShellSort2(arr)
	t.Log("ShellSort2 cost:", time.Since(start))
}
