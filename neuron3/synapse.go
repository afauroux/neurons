package neuron3

import (
	"fmt"
	"math"
	"time"
)

// Connectable an interface to connect together different types of neurons, glias, sensors etc...
type Connectable interface {
	Connect(Connectable)
	MakeSynapse() Synapse
}

// Synapse the link between all neurons cells
type Synapse struct {
	C           chan *Synapse // the receiving chanel of the post synaptic cell
	Weight      int           // the weight associated with this synapse
	Sensitivity float32       // the sensitivity multiplier contolled by glial cells
	AfterDamped time.Timer    // after this synapse fire it will start a timer for when the potential
	// raise that it generated will be damped out
}

func (s *Synapse) String() string {
	return fmt.Sprintf("s: %v x %v", s.Weight, s.Sensitivity)
}

// Potentiate or weaken a synapse by m% of the difference between MAXSIG and its weight
func (s *Synapse) Potentiate(m float64) {
	if s.Weight > 0 {
		s.Weight += int(math.Ceil(float64(MAXSIG-s.Weight) * m))
	} else {
		s.Weight -= int(math.Ceil(float64(MAXSIG+s.Weight) * m))
	}

}

// Fire send its own address to the listening channel
func (s *Synapse) Fire() {
	s.C <- s
}

// GetActivity return the weighted signal that will
// be used to modify the listening neuron's potential
func (s *Synapse) GetActivity() int {
	return int(float32(s.Weight) * s.Sensitivity)
}
