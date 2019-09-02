package neuron

import (
	"fmt"
	"math/rand"
)

// Connectable an interface to connect together different types of neurons, glias, sensors etc...
type Connectable interface {
	Connect(*Synapse)
}

// Synapse the link between all neurons cells
type Synapse struct {
	C           chan int //the receiving chanel of the post synaptic cell
	Weight      int      // the weight associated with this synapse
	Sensitivity int      // the sensitivity multiplier contolled by glial cells
}

// NewSynapse create a new excitatory or inhibitory synapse
func NewSynapse(input chan *Synapse) *Synapse {
	return &Synapse{
		C:           input,
		Weight:      rand.Intn(2*MAXSIG) - MAXSIG,
		Sensitivity: 1,
	}
}

func (s *Synapse) String() string {
	return fmt.Sprintf("s: %v x %v", s.Weight, s.Sensitivity)
}

// Potentiate or weaken a synapse by dw
func (s *Synapse) Potentiate(dw int) {
	s.Weight += dw
}
