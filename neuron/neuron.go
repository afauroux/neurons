// Package neuron contains the class of all neuron like object
package neuron

import (
	"fmt"
	"math"

	NN "github.com/afauroux/neurons/network"
)

// Neuron is an exitable node with many input and output synapses
type Neuron struct {
	ID        int         // unique identifier
	Axones    []*Synapse  // output connections
	Dendrites []*Synapse  // input connections
	Potential float64     // excitation level
	Last      int         // Last tick when the Potential was updated
	Net       *NN.Network // a ref to the network in which this neuron resides

}

func (n *Neuron) String() string {
	return fmt.Sprintf("N%v(%.2f)", n.ID, n.Potential)
}

// total nb of neurons, used to generate unique IDs
var nbNeurones = 0

// generateID generate IDs for new neurons or new glias (only one count)
func generateID() int {
	nbNeurones++
	return nbNeurones
}

// NewNeuron generates a new neuron with default values
func NewNeuron(net *Network) *Neuron {
	return &Neuron{
		ID:  generateID(),
		Net: net,
	}
}

// UpdatePotential updates and retrieves the damped potential since the last tick we checked
func (n *Neuron) UpdatePotential(currenTick int) {
	dt := float64(currenTick - n.Last)
	n.Potential = n.Potential * math.Exp(-dt/DAMPING)
	n.Last = currenTick
}

// Fire send action potential throught all the axones (fire all axones)
// return all potentially excited Neurons
func (n *Neuron) Fire() {
	n.Net.excited[n.ID] = n
	n.Potential = LOWEND
	for _, s := range n.Axones {
		if s.Post.Potential < 0 {
			// the postsynaptic neuronjust fired and is now in recovery
			// the synapse is not in synchro it should weakens
			// according to https://neuronaldynamics.epfl.ch/online/Ch19.S2.html#SS1.p6
			s.ChangeWeight(LTD)
		} else {
			n.Net.excited[s.Post.ID] = s.Post
			s.Post.Potential += s.SignedWeight() * s.Sensitivity
			if s.Post.Potential > TRESH {
				s.Post.Fire()
			}
		}
	}

	for _, s := range n.Dendrites {
		if s.Pre.Potential < 0 {
			// presynaptic neurons participated in this one firing according
			// to HEBB's rule the associated synapses must be strenghtened
			s.ChangeWeight(LTP)
		}
	}

}
