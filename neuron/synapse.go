package neuron

import (
	"fmt"
	"math/rand"
)

// Synapse is a weighted link bewteen two neurons
type Synapse struct {
	Pre, Post   *Neuron
	Weight      float64
	Sensitivity float64
	inib        bool
}

// SignedWeight is the negative weight if s.inib is true (inhibitory synapse)
func (s *Synapse) SignedWeight() float64 {
	var i float64 = 1
	if s.inib {
		i = -1
	}
	return s.Weight * i
}

func (s *Synapse) String() string {
	return fmt.Sprintf("S(%v->%v){%.2f,%.2f,%v}", s.Pre, s.Post, s.Weight, s.Sensitivity, s.inib)
}

// NewSynapse generates a new Synapse with random weight
func NewSynapse(pre, post *Neuron, weight float64, inib bool) *Synapse {
	if weight == 0 {
		weight = rand.Float64() * MAXSIG
	}
	return &Synapse{
		Pre:         pre,
		Post:        post,
		Weight:      weight,
		Sensitivity: 1,
		inib:        inib,
	}
}

// ChangeWeight changes the synaptic weight
func (s *Synapse) ChangeWeight(multiplicator float64) {
	s.Weight *= multiplicator
	if s.Weight > MAXSIG {
		s.Weight = MAXSIG
	}
	if s.Weight <= 0 {
		s.Weight = 0
	}
}
