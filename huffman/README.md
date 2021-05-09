# Huffman encoding tree

## constructor 
contrusct a Huffman tree from a set of pairs of symbols and weight.
Note that input data of paris is also represented as a set of Huffman
nodes each of which has nil pointers as their left and right barnch until a final huffman tree is constructed will they have proper barnch.

## bit stream 
we represent the bit stream as sequense of bytes which is 8 bit long AKA. unint8, hence we need a coder to decode bit stream from sequence
of bytes  