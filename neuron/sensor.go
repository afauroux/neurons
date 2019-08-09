package neuron

// Sensor neurons
type Sensor struct {
	Neurons  []*Neuron
	Data     [][]bool
	Expected [][]bool
}

// MakeXOR make a XOR sensor for testing purposes
func MakeXOR() *Sensor {
	var neurons = []*Neuron{
		NewNeuron(),
		NewNeuron(),
	}
	sensors := &Sensor{
		Neurons: neurons,
		Data: [][]bool{
			[]bool{false, true},
			[]bool{false, false},
			[]bool{true, true},
			[]bool{true, false},
		},
		Expected: [][]bool{ // expected XOR notXOR
			[]bool{true, false},
			[]bool{false, true},
			[]bool{false, true},
			[]bool{true, false},
		},
	}
	return sensors
}
