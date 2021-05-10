# Huffman encoding tree

## constructor 
contrusct a Huffman tree from a set of pairs of symbols and weight.
Note that input data of paris is also represented as a set of Huffman
nodes each of which has nil pointers as their left and right barnch until a final huffman tree is constructed will they have proper barnch.

## bit stream 
we represent the bit stream as sequense of bytes which is 8 bit long AKA. unint8, hence we need a coder to decode bit stream from sequence
of bytes

## example
```shell
root: --Huffman Node(symbol: [B C D G H E F A], weight: 17, 
 -->left node(--Huffman Leaf(symbol: [A], weight: 8)) -->right node(--Huffman Node(symbol: [B C D G H E F], weight: 9, 
 -->left node(--Huffman Node(symbol: [G H E F], weight: 4, 
 -->left node(--Huffman Node(symbol: [E F], weight: 2, 
 -->left node(--Huffman Leaf(symbol: [F], weight: 1)) -->right node(--Huffman Leaf(symbol: [E], weight: 1)))) -->right node(--Huffman Node(symbol: [G H], weight: 2, 
 -->left node(--Huffman Leaf(symbol: [H], weight: 1)) -->right node(--Huffman Leaf(symbol: [G], weight: 1)))))) -->right node(--Huffman Node(symbol: [B C D], weight: 5, 
 -->left node(--Huffman Leaf(symbol: [B], weight: 3)) -->right node(--Huffman Node(symbol: [C D], weight: 2, 
 -->left node(--Huffman Leaf(symbol: [D], weight: 1)) -->right node(--Huffman Leaf(symbol: [C], weight: 1)))))))))
map[0:[A] 1000:[F] 1001:[E] 1010:[H] 1011:[G] 110:[B] 1110:[D] 1111:[C]]
```

## variable-long code

## TODO
0 prefix bits