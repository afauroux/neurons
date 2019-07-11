// Package neurons is an attempt at a neural network simulation
// with each neurons beeing a goroutine that receive signals of others
// via channels in a LTP style backpropagation of discrete signal
package neurons

import (
	"fmt"
	"math/rand"
	"time"
)

// constantes and globals
const dtNeurones = 50 * time.Millisecond  // time between updates
const dtSensors = 1000 * time.Millisecond // time between sensors firing
const buffSize = 100                      // buffer size of the input chanel
const tresh = 100                         // potential reaching this causes firing
const damping = 1                         // potential lost in one DT
const maxSig = 50                         // maximum signal strengh
var nbNeurones = 0                        // total nb of neurons, used to generate unique ids

// generateID generate IDs for new neurons
func generateID() int {
	nbNeurones++
	return nbNeurones
}

// Neuron the default brick
type Neuron struct {
	id             int             // id (total number)
	input          chan int        // every presynaptic Neuron sends his id
	weights        map[int]int     // from the id we know the associated
	parents        map[int]*Neuron // Neuron and weight
	childs         map[int]*Neuron // input chanels of all Neurons listening to this one
	clock          *time.Ticker    //internal clock for potential updates
	thresh         int
	damping        int
	potential      int
	verbose, alive bool
}

// NewNeuron creates a neuron with default values
func NewNeuron() *Neuron {
	n := &Neuron{
		id:        nbNeurones,
		input:     make(chan int, buffSize),
		weights:   make(map[int]int),
		parents:   make(map[int]*Neuron),
		childs:    make(map[int]*Neuron),
		clock:     time.NewTicker(dtNeurones),
		potential: 0,
		verbose:   true,
		alive:     true,
	}
	go n.Update()
	nbNeurones++
	return n
}

//Connect two neurons together (pre and post synaptic)
func Connect(pre, post *Neuron) {
	post.parents[pre.id] = pre
	post.weights[pre.id] = rand.Intn(maxSig)

	pre.childs[post.id] = post

}

// Fire a neuron when its potential is above the threshold
func (n *Neuron) Fire() {
	for _, post := range n.childs {
		post.input <- n.id
	}
}

// Update a neuron potential whenever it receive a msg from a dendrite
func (n *Neuron) Update() {
	for n.alive {
		select {
		case <-n.clock.C:
			n.potential -= damping
			if n.potential <= 0 {
				n.potential = 0
			}
			if n.verbose {
				logchan <- fmt.Sprintf("Nr%v: %v", n.id, n.potential)
			}
		case id := <-n.input:
			n.potential += n.weights[id]
			if n.potential >= tresh {
				n.Fire()
			}
		}
	}
	close(n.input) //closing input should be catched by parents
}

// Sensor are sensitive neurons that fires at regular rates
type Sensor struct {
	potential int
	ticker    *time.Ticker
}
