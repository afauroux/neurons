package neuron4

import (
	"math"
	"time"
)

// Glia ...
type Glia struct {
	Input     chan float64
	Synapses  []*Synapse
	Glias     []*Glia
	Potential float64
	Store     float64
	Last      time.Time
	Active    bool
}

// NewGlia creates a Glia struct with default values
func NewGlia() *Glia {
	return &Glia{
		Input:     make(chan float64),
		Synapses:  make([]*Synapse, 100),
		Glias:     make([]*Glia, 100),
		Potential: 0.0,
		Store:     0.0,
		Last:      time.Now(),
		Active:    false,
	}
}

// GetPotential ...
func (g *Glia) GetPotential() float64 {
	dt := time.Since(g.Last).Seconds()
	return g.Store * math.Exp(-math.Pow(dt-DELAYGLIA, 2)/math.Pow(DAMPING.Seconds(), 2))
}

// Update ...
func (g *Glia) Update() {
	g.Active = true
	for signal := range g.Input {
		g.Store += signal
		g.Potential = g.GetPotential()
		for _, g2 := range g.Glias {
			if !g2.Active {
				go g2.Update()
			}
			g2.Input <- g.Potential

		}
		for _, s := range g.Synapses {
			s.Sensitivity = 1 + g.Potential/GLIAMEANPOT
		}
	}

}
