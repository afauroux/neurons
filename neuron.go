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
var nbNeurones = 0                        // total nb of neurons, used to generate unique IDs

// generateID generate IDs for new neurons
func generateID() int {
	nbNeurones++
	return nbNeurones
}

// Neuron the default brick
type Neuron struct {
	ID        int             // id (total number)
	Input     chan int        // every presynaptic Neuron sends his id
	Weights   map[int]int     // from the id we know the associated
	Parents   map[int]*Neuron // Neuron and weight
	Childs    map[int]*Neuron // input chanels of all Neurons listening to this one
	Clock     *time.Ticker    //internal clock for potential updates
	Thresh    int
	Damping   int
	Potential int
	Log       chan string
	Alive     bool
}

// NewNeuron creates a neuron with default values
func NewNeuron() *Neuron {
	n := &Neuron{
		ID:        nbNeurones,
		Input:     make(chan int),
		Weights:   make(map[int]int),
		Parents:   make(map[int]*Neuron),
		Childs:    make(map[int]*Neuron),
		Clock:     time.NewTicker(dtNeurones),
		Potential: 0,
		Log:       nil,
		Alive:     true,
	}
	go n.Update()
	nbNeurones++
	return n
}

//Connect two neurons together (pre and post synaptic)
func Connect(pre, post *Neuron) {
	post.Parents[pre.ID] = pre
	post.Weights[pre.ID] = rand.Intn(maxSig)

	pre.Childs[post.ID] = post

}

// Fire a neuron when its potential is above the threshold
func (n *Neuron) Fire() {
	for _, post := range n.Childs {
		post.Input <- n.ID
	}
}

// Update a neuron potential whenever it receive a msg from a dendrite
func (n *Neuron) Update() {
	for n.Alive {
		select {
		case <-n.Clock.C:
			n.Potential -= damping
			if n.Potential <= 0 {
				n.Potential = 0
			}
			if n.Potential >= tresh {
				n.Potential = tresh
			}
			if n.Log != nil {
				n.Log <- fmt.Sprintf("Nr%v: %v", n.ID, n.Potential)
			}
		case ID := <-n.Input:
			n.Potential += n.Weights[ID]
			if n.Potential >= tresh {
				n.Fire()
			}
		}
	}
	close(n.Input) //closing input should be catched by parents
}

// Sensor are sensitive neurons that fires at regular rates
type Sensor struct {
	potential int
	ticker    *time.Ticker
}
