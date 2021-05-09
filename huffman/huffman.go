package huffman

import (
	"fmt"
	"sort"
)

type Huffman struct {
	sybomls []string
	weight  uint
	lnode   *Huffman
	rnode   *Huffman
}

type IHuffman interface {
	Decode([]byte) string
	Encode(string) []byte
	IsLeaf() bool
	// chose next node according to current bit
	Next(uint) IHuffman
}

type coder struct {
	buf []byte
	// [0, infinit)
	offset int
}

func (c *coder) read() int {
	if c.offset/8 > len(c.buf)-1 {
		return -1
	}
	defer func() {
		c.offset++
	}()
	n := c.offset / 8
	r := c.offset % 8
	mask := 1 << (7 - r)
	b := c.buf[n]
	if (b & uint8(mask)) == 0 {
		return 0
	}
	return 1
}

// New reutrn a Huffman node
func New(symbol string, weight uint) *Huffman {
	tmp := make([]string, 0, 1)
	tmp = append(tmp, symbol)
	return &Huffman{
		sybomls: tmp,
		weight:  weight,
	}
}

// Merge merge two Huffman nodes into one
// a is implicitly smaller or equal to b;
// a leaf node always on the left branch
func Merge(a, b *Huffman) *Huffman {
	tmp := make([]string, 0, len(a.sybomls)+len(b.sybomls))
	tmp = append(tmp, b.sybomls...)
	tmp = append(tmp, a.sybomls...)
	if a.IsLeaf() {
		return &Huffman{
			sybomls: tmp,
			weight:  a.weight + b.weight,
			lnode:   a,
			rnode:   b,
		}
	}
	if b.IsLeaf() {
		return &Huffman{
			sybomls: tmp,
			weight:  a.weight + b.weight,
			lnode:   b,
			rnode:   a,
		}
	}
	return &Huffman{
		sybomls: tmp,
		weight:  a.weight + b.weight,
		lnode:   a,
		rnode:   b,
	}
}

// NewHuffman construct a Huffman tree and reutrn the root
// of it from the input data
// the input data is a sequence of Huffman leaf node
func NewHuffman(set []*Huffman) *Huffman {
	if len(set) == 0 {
		return nil
	}
	if len(set) == 1 {
		return set[0]
	}

	sort.Slice(set, func(i, j int) bool {
		return set[i].weight > set[j].weight
	})
	var cursor *Huffman
	for len(set) > 1 {
		cursor = Merge(set[len(set)-1], set[len(set)-2])
		set = set[:len(set)-1]
		set[len(set)-1] = cursor
		sort.Slice(set, func(i, j int) bool {
			return set[i].weight > set[j].weight
		})
	}
	return cursor
}

// Decode
func (h *Huffman) Decode(b []byte) []string {
	c := new(coder)
	c.buf = make([]byte, len(b))
	copy(c.buf, b)
	var res []string
	current := h
	for i := 0; i < len(c.buf)*8; i++ {
		if current.IsLeaf() {
			res = append(res, current.sybomls...)
			current = h
			continue
		}
		bit := c.read()
		if bit == -1 {
			break
		}
		current = current.Next(bit)
	}
	return res
}

// Next return the next node accoriding to the current bit
func (h *Huffman) Next(b int) *Huffman {
	if b == 0 {
		return h.lnode
	}
	return h.rnode
}

// IsLeaf predicate if the current node if a leaf
func (h *Huffman) IsLeaf() bool {
	if len(h.sybomls) == 1 && h.lnode == nil && h.rnode == nil {
		return true
	}
	return false
}

func (h *Huffman) CodeTable() map[string][]string {
	res := make(map[string][]string)
	if h.IsLeaf() {
		return nil
	}
	hd := func(node *Huffman, bits int) {
		if node.IsLeaf() {
			res[fmt.Sprintf("%b", bits)] = node.sybomls
		}
	}
	tr := NewTraversal(hd, h)
	tr.PreOrder(h, 0)
	return res
}

func (h *Huffman) String() string {
	if h.IsLeaf() {
		return fmt.Sprintf("--Huffman Leaf(symbol: %v, weight: %v)", h.sybomls, h.weight)
	}
	return fmt.Sprintf("--Huffman Node(symbol: %v, weight: %v, \n -->left node(%v) -->right node(%v))", h.sybomls, h.weight, h.lnode, h.rnode)
}

func (h *Huffman) codeTable() map[string]string {
	return map[string]string{}
}

type Traversal struct {
	root    *Huffman
	handler func(node *Huffman, baseBits int)
}

func NewTraversal(handler func(*Huffman, int), root *Huffman) *Traversal {
	return &Traversal{
		root:    root,
		handler: handler,
	}

}

// PreOrder Traversal
func (tr *Traversal) PreOrder(node *Huffman, baseBits int) {
	if node == nil {
		return
	}
	tr.handler(node, baseBits)
	tr.PreOrder(node.lnode, baseBits<<1)
	tr.PreOrder(node.rnode, (baseBits<<1)+1)
}
