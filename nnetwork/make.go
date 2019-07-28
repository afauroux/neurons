package nnetwork

// MakeNeuralNetwork creates a fully connected network with
// layer sizes defined by shape
func MakeNeuralNetwork(shape []int, loop bool) [][]*Neuron {
	n := make([][]*Neuron, len(shape))
	for i, s := range shape {
		n[i] = make([]*Neuron, s)
		for j := 0; j < s; j++ {
			n[i][j] = NewNeuron()
			n[i][j].X = i - len(shape)/2
			n[i][j].Y = j - s/2
			if i >= 1 {
				for k := 0; k < shape[i-1]; k++ {
					Connect(n[i-1][k], n[i][j])
				}
			}
		}
	}
	if loop {
		for j, pre := range n[len(shape)-1] {
			Connect(pre, n[0][j])
		}
	}
	return n
}
