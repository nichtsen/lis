package deriv

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	a := "日本語"
	fmt.Println(len(a))
	var r = []rune{'0', '1', '2', '9'}
	fmt.Println(r)
}
