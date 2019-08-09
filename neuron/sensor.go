package neuron

// Sensor neurons
type Sensor struct {
	Neurons  []Neuron
	Data     [][]bool
	Expected []bool
}

// MakeXOR make a XOR sensor for testing purposes
func MakeXOR() *Sensor {
	var neurons = []Neuron{
		NewNeuron1(),
		NewNeuron1(),
	}
	sensors := &Sensor{
		Neurons: neurons,
		Data: [][]bool{
			[]bool{false, true, true, false},
			[]bool{true, false, true, false},
		},
		Expected: []bool{true, true, false, false},
	}
	return sensors
}
