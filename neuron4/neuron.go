package neuron4

import (
	"math"
	"time"
)

// Network is the base class representing a neural net
type Network struct {
	Nmap map[int]*Neuron // a map for easy acces of a neuron struct from its ID
	Net  [][]*Neuron     // neural net structure
}

// Neuron ...
type Neuron struct {
	Buffer    []*Synapse
	Axones    []*Synapse
	Potential float64
	Last      time.Time
}

// Synapse ...
type Synapse struct {
	Weight      float64 // the weight associated with this synapse
	Sensitivity float64 // the sensitivity multiplier contolled by glial cells
	Post        *Neuron
	Pre         *Neuron
	Glia        *Glia
}

// AdjustWeight ...
func (s *Synapse) AdjustWeight(dw float64) {
	if s.Weight > 0 {
		s.Weight = (MAXSIG - s.Weight) * dw
	} else {
		s.Weight = (-MAXSIG - s.Weight) * dw
	}
}

// Glia ...
type Glia struct {
	S         *Synapse
	Glias     []*Glia
	Store     float64
	Potential float64
	Last      time.Time
}

// Update ...
func (g *Glia) Update(a Activeable) {
	newP := g.GetActivation()
	if newP <= eps {
		g.Potential = 0
	}
	g.Potential = newP + a.GetActivation()*ACTIVATION
	g.Last = time.Now()
	for _, g2 := range g.Glias {
		go g2.Update(g)
	}
	g.S.Sensitivity = g.Potential / SENSITIVITY

}

// Update ...
func (n *Neuron) Update(s *Synapse) bool {
	n.Buffer = append(n.Buffer, s)
	newP := n.Potential * math.Exp(-time.Since(n.Last).Seconds()/DAMPING.Seconds())
	n.Last = time.Now()
	// the neuron was in recovery
	if n.Potential < 0 {
		if newP >= -eps {
			n.Potential = 0
		} else {
			n.Potential = newP
			return false // if in recovery it can't receive signal
		}
	}
	// the neuron was activated but its activity went to 0
	if newP <= eps && n.Potential > 0 {
		for _, s := range n.Buffer {
			s.AdjustWeight(LTD)
		}
		n.Buffer = []*Synapse{}
		n.Potential = 0
	}

	// the neuron is receiving signal from the synapse
	n.Potential += s.GetActivation()
	if n.Potential > TRESH {
		n.Fire()
		return true
	}
	return false
}

// Activeable ...
type Activeable interface {
	GetActivation() float64
}

// GetActivation ...
func (s *Synapse) GetActivation() float64 {
	return s.Weight * s.Sensitivity
}

// GetActivation ...
func (g *Glia) GetActivation() float64 {
	dt := time.Since(g.Last).Seconds()
	return g.Potential * math.Exp(-math.Pow(dt-DELAYGLIA, 2)/math.Pow(DAMPING.Seconds(), 2))
}

// Fire ...
func (n *Neuron) Fire() {
	for _, s := range n.Buffer {
		s.AdjustWeight(LTP)
	}
	n.Buffer = []*Synapse{}
	n.Potential = LOWEND // the neuron goes in recovery mode

	for _, s := range n.Axones {
		go s.Post.Update(s)
		go s.Glia.Update(s)
	}
}
