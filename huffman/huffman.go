package huffman

type Huffman struct {
	sybomls string
	weight  int
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
	if c.offset/8 > len(c.buf) {
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

func (h *Huffman) Decode(b []byte) string {
	c := new(coder)
	copy(c.buf, b)
	var res string
	current := h
	for i := 0; i < len(c.buf); i++ {
		if h.IsLeaf() {
			res += h.sybomls
			current = h
			continue
		}
		bit := c.read()
		if bit == -1 {
			break
		}
		current = h.Next(bit)
	}
	_ = current
	return res
}

func (h *Huffman) Next(b int) *Huffman {
	if b == 0 {
		return h.lnode
	}
	return h.rnode
}

func (h *Huffman) IsLeaf() bool {
	if len(h.sybomls) == 1 {
		return true
	}
	return false
}
