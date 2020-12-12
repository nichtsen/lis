package lis

import (
	"fmt"
	"testing"
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

func BenchmarkLIS(b *testing.B) {
	dyp := NewDyp([]int{5, 7, 8, 9, 1, 2, 3})
	for i := 0; i < b.N; i++ {
		_ = dyp.LIS()
	}
}

func BenchmarkLISdynamic(b *testing.B) {
	dyp := NewDyp([]int{5, 7, 8, 9, 1, 2, 3})
	for i := 0; i < b.N; i++ {
		_ = dyp.LISdynamic()
	}
}
