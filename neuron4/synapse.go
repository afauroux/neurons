package neuron4

// Synapse ...
type Synapse struct {
	Weight      float64 // the weight associated with this synapse
	Sensitivity float64 // the sensitivity multiplier contolled by glial cells
	Post        *Neuron
	Pre         *Neuron
	Glia        *Glia
}

// Fire ...
func (s *Synapse) Fire() {
	s.Post.Input <- s
	s.Glia.Input <- s.GetPotential()
}

// GetPotential ...
func (s *Synapse) GetPotential() float64 {
	return s.Weight * s.Sensitivity
}

// AdjustWeight ...
func (s *Synapse) AdjustWeight(dw float64) {
	if s.Weight > 0 {
		s.Weight = (MAXSIG - s.Weight) * dw
	} else {
		s.Weight = (-MAXSIG - s.Weight) * dw
	}
}
