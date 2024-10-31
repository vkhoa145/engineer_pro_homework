package main

import (
	"fmt"

	"github.com/vkhoa145/engineer_pro_homework/day1"
)

func main() {
	fmt.Println("Hello, World!")
	area := day1.CalculateArea(10, 20)
	perimeter := day1.CalculatePerimeter(10, 20)
	fmt.Printf("area = %d\n", area)
	fmt.Printf("perimeter = %d\n", perimeter)

	isDividedByTwo := day1.IsDividedByTwo("hello0")
	fmt.Printf("isDividedByTwo: %v\n", isDividedByTwo)

	slice := []int{1, 2, 4, 3, 10, 5, 99, 7}
	day1.ModifySlice(slice)

	day1.BubbleSort(slice)
	fmt.Println("sort", slice)
}
