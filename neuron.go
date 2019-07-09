package neuron

const dt int = 1

// Axone ...
type Axone struct {
	pre, pos *Neuron
	msg      chan int
}

// Dendrite ...
type Dendrite struct {
	pre, pos *Neuron
	msg      chan int
	synapse  int
}

// Neuron ...
type Neuron struct {
	axone          Axone
	dendrites      []Dendrite
	thresh         int
	damping        int
	potential      int
	verbose, alive bool
}

// Fire a neuron when its potential is above the threshold
func (n *Neuron) Fire() {}

// Update a neuron potential whenever it receive a msg from a dendrite
func (n *Neuron) Update() {}

// New creates a neuron with default values
func New() (n *Neuron) {

	return &Neuron{
		axone:     Axone{},
		dendrites: make([]Dendrite, 0),
		thresh:    90,
		damping:   10 * dt,
		potential: 0,
		verbose:   false,
		alive:     true,
	}
}
