package network

import (
	"math/rand"

	N "github.com/afauroux/neurons/neuron"
)

// MakeNeuralNetwork creates a fully connected network with
// layer sizes defined by shape
func MakeNeuralNetwork(shape []int, loop, fullyconnected bool, linkproba float64) [][]N.Neuron1 {
	n := make([][]N.Neuron1, len(shape))
	for i, s := range shape {
		n[i] = make([]N.Neuron1, s)
		for j := 0; j < s; j++ {
			n[i][j] = N.NewNeuron1()
			n[i][j].X = i - len(shape)/2
			n[i][j].Y = j - s/2
			if i >= 1 {
				if fullyconnected {
					for k := 0; k < shape[i-1]; k++ {
						N.Connect(n[i-1][k], n[i][j], 0)
					}
				} else {
					for k := 0; k < shape[i-1]; k++ {
						if rand.Float64() >= linkproba { // flip a coin
							N.Connect(n[i-1][k], n[i][j], 0)
						}
					}
				}
			}
		}
	}
	if loop {
		for j, pre := range n[len(shape)-1] {
			N.Connect(pre, n[0][j], 0)
		}
	}
	return n
}
