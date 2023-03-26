package main

import (
	"fmt"
)

func main() {
	data1 := [][][]int{
		{
			{1, 2, 3}, // 6
			{1, 2, 3}, // 6
		},
		{
			{1, 2, 3}, // 6
			{1, 2},    // 3
			{1},       // 1
		},
	} // 22
	deepSum(data1)

	data2 := [][][]int{
		{
			{1, 1, 1}, // 3
			{1},       // 1
		},
		{
			{1, 1, 2}, // 4
			{3, 2, 1}, // 6
		},
		{{3}}, // 3
	} // 17
	deepSum(data2)
}
func deepSum(data [][][]int) {
	// logic to sum all the int from multidimensional array
	result := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			for k := 0; k < len(data[i][j]); k++ {
				result += data[i][j][k]
			}
		}
	}
	fmt.Println(result)
	return
}

// func deepSum(arr interface{}) int {
// 	switch arr.(type) {
// 	case int:
// 		return arr.(int)
// 	case []int:
// 		sum := 0
// 		for _, v := range arr.([]int) {
// 			sum += v
// 		}
// 		return sum
// 	case []interface{}:
// 		sum := 0
// 		for _, v := range arr.([]interface{}) {
// 			sum += deepSum(v)
// 		}
// 		return sum
// 	default:
// 		return 0
// 	}
// 	fmt.Println(deepSum(data1))
// 	fmt.Println(deepSum(data2))
// }
