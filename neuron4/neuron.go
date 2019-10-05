package neuron4

import (
	"fmt"
	"math"
	"time"
)

// log ...
func (n *Neuron) log(str string, values ...interface{}) {
	if n.Log != nil {
		n.Log <- fmt.Sprintf(str, values...)
	}
}

// Neuron ...
type Neuron struct {
	Input     chan *Synapse
	Buffer    []*Synapse
	Axones    []*Synapse
	Potential float64
	Last      time.Time
	Active    bool
	Log       chan string
}

// NewNeuron creates a Neuron struct with default values
func NewNeuron() *Neuron {
	return &Neuron{
		Input:     make(chan *Synapse),
		Buffer:    make([]*Synapse, 100),
		Axones:    make([]*Synapse, 100),
		Potential: 0.0,
		Last:      time.Now(),
		Active:    false,
	}
}

// Connect ...
func Connect(n1, n2 *Neuron, w1, w2 float64) {
	n1.

}

// Update ...
func (n *Neuron) Update() {
	for s := range n.Input {
		newP := n.GetPotential()

		// ** The neuron is in recovery **
		if n.Potential < 0 {
			if newP >= -EPS {
				n.Potential = 0
			} else {
				n.Potential = newP
				continue // if in recovery it can't receive signal
			}
		}

		// ** the neuron was activated but its activity went to 0 **
		if newP <= EPS && n.Potential > 0 {
			for _, s := range n.Buffer {
				s.AdjustWeight(LTD)
			}
			n.Buffer = []*Synapse{}
			n.Potential = 0
		}

		// ** the neuron is receiving signal from a synapse **
		n.Buffer = append(n.Buffer, s) // we buffer the responsible synapse
		n.Potential += s.GetPotential()
		n.Last = time.Now()
		if n.Potential > TRESH {
			n.Fire()
		}
	}
}

// Fire ...
func (n *Neuron) Fire() {
	for _, s := range n.Buffer {
		s.AdjustWeight(LTP)
	}
	n.Buffer = []*Synapse{}
	n.Potential = LOWEND // the neuron goes in recovery mode (LOWEND is < 0)

	for _, s := range n.Axones {
		s.Fire()
	}
}

// GetPotential ...
func (n *Neuron) GetPotential() float64 {
	dt := time.Since(n.Last).Seconds()
	return n.Potential * math.Exp(-dt/DAMPING.Seconds())
}
