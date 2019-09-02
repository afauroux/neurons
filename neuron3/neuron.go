package neuron3

import (
	"math"
	"time"
)

// Activable is the interface that is used for neurons
// and glial cells Activation Potential handling
type Activable interface {
	RaisePotential(*Synapse) bool
	Update()
}

// NeuroPotential an implementation of the actveable interface in the case of neuron potential
type NeuroPotential struct {
	P      float64
	Last   time.Time
	Buffer []*Synapse
}

//GetPotentialAt ...
func (p *NeuroPotential) GetPotentialAt(time.Time) {

}

//Update ...
func (p *NeuroPotential) Update() {
	dt := time.Since(p.Last)
	p.P += p.P * math.Exp(-dt/DAMPING)

}

// RaisePotential ...
func (p *NeuroPotential) RaisePotential(s *Synapse) bool {
	p.Update()
	p.Buffer = append(p.Buffer, s)
	//p.P += s.GetActivity()
	p.Last = time.Now()
	if p.P >= TRESH {
		return true
	}
	//p.AfterDamping = time.AfterFunc(time.Second, p.DampedOut)
	p.GetPotentialAt = func() {}
	return false
}
