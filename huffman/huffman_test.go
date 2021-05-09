package huffman

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Println(0b100)
	fmt.Printf("%b\n", 0b110)
	fmt.Println(32 & 01 << 2)
	var str = "abc"
	fmt.Println(len(str))
	tmp := make([]int, 10)
	fmt.Println("---")
	for len(tmp) > 1 {
		tmp = tmp[:len(tmp)-1]
		fmt.Println(len(tmp))
	}
}

func TestConsturctor(t *testing.T) {
	input := map[string]uint{
		"A": 8,
		"B": 3,
		"C": 1,
		"D": 1,
		"E": 1,
		"F": 1,
		"G": 1,
		"H": 1,
	}
	var set []*Huffman
	for symbol, weight := range input {
		set = append(set, New(symbol, weight))
	}
	// fmt.Println(set)
	h := NewHuffman(set)
	fmt.Printf("root: %v\n", h)
	tab := h.CodeTable()
	fmt.Println(tab)
	bee := []byte{0b11010011, 0b00100000}

	str := h.Decode(bee)
	fmt.Println(str)
}

func TestC(t *testing.T) {
	bee := []byte{0b11010011, 0b00100000}
	coder := new(coder)
	coder.buf = bee
	fmt.Println(len(coder.buf))
	for {
		i := coder.read()
		fmt.Printf("%d", i)
		if i == -1 {
			break
		}
	}
}

func TestCoder(t *testing.T) {
	bee := int64(0b11011111)
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, bee)
	fmt.Println(n)
	bin := binary.LittleEndian.Uint64(buf)
	buf = buf[:2]
	fmt.Printf("%b\n", bee)
	fmt.Printf("%b\n", bin)
	fmt.Printf("%b | %b\n", buf[0], buf[1])
	coder := new(coder)
	coder.buf = buf
	fmt.Println(len(coder.buf))

	for {
		i := coder.read()
		fmt.Printf("%d", i)
		if i == -1 {
			break
		}
	}
}
