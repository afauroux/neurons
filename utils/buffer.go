// Package utils is regrouping several structures and methods that are not
// directly related to neural networks
package utils

// Buffer is a simple lifo list type buffer
type Buffer []int

// Push to a buffer in last position
func (b *Buffer) Push(v int) {
	*b = append(*b, v)
}

// Pop the first item of a buffer
func (b *Buffer) Pop() int {
	ret := (*b)[0]
	*b = (*b)[1:]
	return ret
}

// Empty returns true if the Buffer is empty
func (b *Buffer) Empty() bool {
	return len(*b) == 0
}

// BufferCount is a less simple lifo list type buffer with multiplicity
// it get filled with cells containing an ID and the associate value
type BufferCount [][2]int

// Push to a buffer in last position
func (b *BufferCount) Push(ID, value int) {
	*b = append(*b, [2]int{ID, value})
}

// Pop the first item of a buffer count
// if the first cell is now empty then we remove it
func (b *BufferCount) Pop() int {
	v := (*b)[0][1] - 1
	if v <= 0 {
		*b = (*b)[1:]
	}
	return (*b)[0][0]
}
