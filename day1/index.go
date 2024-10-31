package day1

import "fmt"

func CalculateArea(a int, b int) int {
	return a * b
}

func CalculatePerimeter(a int, b int) int {
	return (a + b) * 2
}

func IsDividedByTwo(str string) bool {
	if len(str)%2 == 0 {
		return true
	} else {
		return false
	}
}

func ModifySlice(slice []int) {
	var total int
	var max int
	min := slice[0]
	for _, ele := range slice {
		total += ele

		if ele > max {
			max = ele
		}

		if ele < min {
			min = ele
		}
	}
	fmt.Printf("total = %d\n", total)
	fmt.Printf("max = %d\n", max)
	fmt.Printf("min = %d\n", min)
	fmt.Printf("average = %.2f\n", float64(total)/float64(len(slice)))
}

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, ele := range nums {
		complete := target - ele
		if j, found := numMap[complete]; found {
			return []int{i, j}
		}

		numMap[ele] = i
	}

	return nil
}
