// 希尔排序

package _3_sorts

func ShellSort1(arr []int) {
	length := len(arr)
	gap := 1
	for gap < length/3 {
		gap = gap*3 + 1
		//fmt.Println("gap:", gap)
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
}

func ShellSort2(nums []int) {
	length := len(nums)
	// 外层步长控制
	for step := length / 2; step > 0; step /= 2 {
		// 开始插入排序
		for i := step; i < length; i++ {
			// 满足条件则插入
			for j := i - step; j >= 0 && nums[j+step] < nums[j]; j -= step {
				nums[j], nums[j+step] = nums[j+step], nums[j]
			}
		}
	}
}
