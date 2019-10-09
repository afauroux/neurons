package main

import "math"

// isIdle evaluates to true if there is no neuron firing and no excited neurons
func (net *Network) isIdle() bool {
	return len(net.excited) == 0
}

// Update general update method for networks
func (net *Network) Update() {
	net.UpdatePotentials()
	net.currenTick++
}

// UpdatePotentials update the potentials of all excited neurons and update the list of excited
func (net *Network) UpdatePotentials() {
	stillexcited := make(map[int]*Neuron)
	for _, n := range net.excited {
		n.UpdatePotential(net.currenTick)
		if math.Abs(n.Potential) > eps { // it could be negative when the neuron is recovering
			stillexcited[n.ID] = n
		} else {
			n.Potential = 0
			for _, s := range n.Dendrites {
				s.ChangeWeight(LTP)
			}
			// augment sensibilities/weights ?
			// diminishe weight of bad excitatory neurons ?
		}
	}
	net.excited = stillexcited
}
