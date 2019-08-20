package neuron

import "fmt"

// Glia emulate astrocyte glial cells
type Glia struct {
	ID        int             // a unique identifier (total number of neural cells)
	Input     chan int        // the input chanel through which all signal arrives
	Neurons   map[int]*Neuron // all neurons connected to this cell
	Glias     map[int]*Glia   // all glias connected to this cell
	Weights   map[int]int     // the weights associated with each connections
	Potential int             // the activity Potential of this cell
	Buffer    []int           // a list of all glias and neurons that sent signal since last firing
}

// NewGlia the default constructor
func NewGlia() *Glia {
	var g Glia
	return &g
}

// Fire is the function called when a glial cell fires
func (g *Glia) Fire() {

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
	return fmt.Sprintf("G%v", g.ID)
}

// Update is the goroutine associated with a Glia that will be executed
func (g *Glia) Update() {

}
