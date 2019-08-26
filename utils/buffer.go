// Package utils is regrouping several structures and methods that are not
// directly related to neural networks
package utils

// Buffer is a simple lifo list type buffer
type Buffer []int

// Push to a buffer in last position
func (b *Buffer) Push(v int) int {
	*b = append(*b, v)
	return v
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
