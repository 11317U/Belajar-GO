package main

import (
	"fmt"
)

func main() {
	// Loop from 0 to 4
	for i := 0; i <= 4; i++ {
		// Output the value of i
		fmt.Printf("Nilai i = %d\n", i)
	}
	for j := 0; j <= 10; j++ {
		if j <= 4 {
			fmt.Printf("Nilai j = %d\n", j)
		} else {
			switch j {
			case 5:
				fmt.Printf("character U+0421 'С' starts at byte position %d\n", (j-5)*2)
			case 6:
				fmt.Printf("character U+0410̀ 'А̀' starts at byte position %d\n", (j-5)*2)
			case 7:
				fmt.Printf("character U+0428 'Ш' starts at byte position %d\n", (j-5)*2)
			case 8:
				fmt.Printf("character U+0410̀ 'А̀' starts at byte position %d\n", (j-5)*2)
			case 9:
				fmt.Printf("character U+0420 'Р' starts at byte position %d\n", (j-5)*2)
			case 10:
				fmt.Printf("character U+0412 'В' starts at byte position %d\n", (j-5)*2)
				fmt.Printf("character U+041E 'О' starts at byte position %d\n", (j-5)*2+2)
				for j := 0; j <= 10; j++ {
					if j > 5 {
						fmt.Printf("Nilai j = %d\n", j)
					}
				}
			}
		}
	}

	// Loop from 0 to 10
	// for j := 0; j <= 10; j++ {
	// 	// 	// Output the value of j
	// 	fmt.Printf("Nilai j = %d\n", j)

	// if j < 6 {
	// 		fmt.Printf("Nilai j = %d\n", j)
	// 	} else if j == 6 {
	// 		fmt.Printf("character U+%04X 'С' strats at byte position %d\n", '\u0421', j*2)
	// 	} else if j == 7 {
	// 		fmt.Printf("character U+%04X 'А' strats at byte position %d\n", '\u0410', j*2)
	// 	} else if j == 8 {
	// 		fmt.Printf("character U+%04X 'Ш' strats at byte position %d\n", '\u0428', j*2)
	// 	} else if j == 9 {
	// 		fmt.Printf("character U+%04X 'А' strats at byte position %d\n", '\u0410', j*2)
	// 	} else if j == 10 {
	// 		fmt.Printf("character U+%04X 'Р' strats at byte position %d\n", '\u0420', j*2)
	// 	} else {
	// 		fmt.Printf("character U+%04X 'В' strats at byte position %d\n", '\u0412', j*2)
	// 		break
	// }
	// }
}
