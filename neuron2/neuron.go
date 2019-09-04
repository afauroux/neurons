package neuron2

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Connectable an interface to connect together different types of neurons, glias, sensors etc...
type Connectable interface {
	Connect(Connectable)
	MakeSynapse() Synapse
}

// Fireable is the common interface for all fireable cells
type Fireable interface {
	Fire()
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

// Neuron is the default neuron implementation
type Neuron struct {
	Dendrites
	Axones
}

// Dendrites store the input synapses
type Dendrites struct {
	Input     chan *Synapse
	Buffer    []*Synapse
	Potential int
}

// Axones store the output synapses
type Axones struct {
	Output []Synapse
}

// getDampedValue returns a damped value after 'damping' was removed from it every DT
// since last
// returns min(0, value - damping * time.Since(last)/ DT)
func getDampedValue(value, damping int, last time.Time) int {
	return int(math.Min(0, DAMPING*float64(time.Since(last)/DT)))
}

// Update is called whenever the potential should be updated
func (n *Neuron) Update(tick time.Time) {

}

// Receive signal from its input chanel and update its potential
func (n *Neuron) Receive(synapse Synapse) {
	for synapse := range n.Input {
		// we calculate the new potential based on the time since the last
		// signal was received and the damping that occured
		if len(n.Buffer) > 0 && n.Potential > 0 {
			n.Potential = getDampedValue(n.Potential, DAMPING, n.Buffer[len(n.Buffer)-1].LastFired)
		}

		// We append to the buffer and add the synapse firing to the buffer
		//synapse.AfterDamped
		t := time.AfterFunc(DT, func() { n.Buffer = []*Synapse{} })
		n.Buffer = append(n.Buffer, synapse)

		// The potential is updated
		n.Potential += synapse.GetActivity()

		// if we reached the treshold this neuron fires
		if n.Potential > TRESH {
			n.Fire()
		}
	}
}

// Fire a neuron consist of firing all its Output synapses
// and potentiating all synapses that contributed in elaborating
// this action potential
func (n *Neuron) Fire() {
	for _, s := range n.Output {
		s.Fire()
	}
	for _, s := range n.Buffer {
		if getDampedValue(s.GetActivity(), DAMPING, s.LastFired) > 0 {

		}
	}
}

// Connect a neuron to any connectable
func (n *Neuron) Connect(c Connectable) {
	n.Output = append(n.Output, c.MakeSynapse())
}

// MakeSynapse ...
func (n *Neuron) MakeSynapse(inhibitory bool) *Synapse {
	var w int
	if inhibitory {
		w = -rand.Intn(MAXSIG)
	} else {
		w = rand.Intn(2*MAXSIG) - MAXSIG
	}
	return &Synapse{
		C:           n.Input,
		Weight:      w,
		Sensitivity: 1,
	}
}

// Update is the main goroutine to update a neuron's activity
func (n *Neuron) Update() {
	//for n.Alive {

	//}
}
