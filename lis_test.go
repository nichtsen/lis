package lis

import (
	"fmt"
)

func Example_lis() {
	dyp := NewDyp([]int{5, 7, 8, 9, 1, 2, 3})
	ans := dyp.LIS()
	fmt.Printf("%v", ans)
	len := dyp.LISdynamic()
	fmt.Println(len)
	// Output:
	// 44
}
