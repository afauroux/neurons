package neuron

import (
	"fmt"

	"github.com/afauroux/neurons/utils"
)

// Glia emulate astrocyte glial cells
type Glia struct {
	ID        int              // a unique identifier (total number of neural cells)
	Input     chan int         // the input chanel through which all signal arrives
	Dendrites map[int]*Synapse // all parents synapse
	Axones    map[int]*Synapse // all channels (Neurons or Glia) listening to this neuron activity
	Potential int              // the activity Potential of this cell
	Buffer    utils.Buffer     // a buffer containing the received IDs with multiplicity corresponding to weight
	Store     int              // signals comming in are first stored and potential slowly rise
	Food      int              // glias food level (firing too often: splits, starving: make new connections)
}

// NewGlia the default constructor
func NewGlia() *Glia {
	return &Glia{
		ID:        generateID(),
		Input:     make(chan int),
		Axones:    make(map[int]*Synapse),
		Dendrites: make(map[int]*Synapse),
		Potential: 0,
		Buffer:    []int{},
		Food:      FOODBASE,
	}
}

// Fire is the function called when a glial cell fires
func (g *Glia) Fire() {
	for _, syn := range g.Axones {
		syn.C <- g.Potential
	}
}

// Feed is a function that takes care of feeding (rewarding) neurons that
// contributed to this cell firing
func (g *Glia) Feed() {

}

// ConnectGlias two Glial cell together
func (g *Glia) ConnectGlias(g2 *Glia) {

}

// InterConnect a Glial cell with a Neuron
func (g *Glia) InterConnect(n *Neuron) {

}

// CheckFood checks the food level of each neurons and:
// - if higher than FOODSPLIT split it
// - if lower than 0 create new dendrites (parent connection)
func (g *Glia) CheckFood(n *Neuron) {

}

// CheckSynchroNeuro is a function that check if another neuron or Glia was fired
// in synchro with the neurons or glias that caused this ones firing
// and either connect to it if it was not connected or reenforce the link
func (g *Glia) CheckSynchroNeuro() {

}

// String return a string representation
func (g *Glia) String() string {
	return fmt.Sprintf("g%v(%3d)", g.ID, g.Potential)
}

// Update is the goroutine associated with a Glia that will be executed
func (g *Glia) Update() {

}

// GetID unique identifier among all cell types
func (g *Glia) GetID() int {
	return g.ID
}

// GetInput the Input chanel on which every cell listens
func (g *Glia) GetInput() chan int {
	return g.Input
}

// GetAxons the output Synapses
func (g *Glia) GetAxons() map[int]*Synapse {
	return g.Axones
}

// GetDendrites the input Synapses
func (g *Glia) GetDendrites() map[int]*Synapse {
	return g.Dendrites
}
