// Package nnetwork is an attempt at a neural network simulation
// with each neurons beeing a goroutine that receive signals of others
// via channels in a LTP style backpropagation of discrete signal
package nnetwork

import (
	"fmt"
	"time"
)

// constantes and globals

// DT is the time between updates
const DT = 20 * time.Millisecond

// BUFFSIZE size of the input channels
const BUFFSIZE = 100

// TRESH potentials reaching this causes firing
const TRESH = 100

// DAMPING is the potential lost in one DT
const DAMPING = 5

// MAXSIG is the maximum signal strengh
const MAXSIG = 100

// LOWEND is the negative potential after firing
// that creates a cooldown time
const LOWEND = -3

// total nb of neurons, used to generate unique IDs
var nbNeurones = 0

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
	X, Y      int // used for ploting (see gui.go)

}

// NewNeuron creates a neuron with default values
func NewNeuron() *Neuron {
	n := &Neuron{
		ID:        nbNeurones,
		Input:     make(chan int, BUFFSIZE),
		Weights:   make(map[int]int),
		Parents:   make(map[int]*Neuron),
		Childs:    make(map[int]*Neuron),
		Clock:     time.NewTicker(DT),
		Potential: 0,
		Log:       nil,
		Alive:     true,
		X:         nbNeurones, // by default its a linear system
		Y:         0,
	}
	go n.Update()
	nbNeurones++
	return n
}

//Connect two neurons together (pre and post synaptic)
func Connect(pre, post *Neuron) {
	post.Parents[pre.ID] = pre
	post.Weights[pre.ID] = MAXSIG //rand.Intn(MAXSIG*2) - MAXSIG
	pre.Childs[post.ID] = post
}

// Fire a neuron when its potential is above the threshold
func (n *Neuron) Fire() {
	time.Sleep(DT)
	n.Potential = LOWEND // negative potential so neuron can fire only after some regeneration
	for _, post := range n.Childs {
		post.Input <- n.ID
	}
	for id, p := range n.Parents {
		if p.Potential < 0 { //fired recently
			n.Weights[id]++ // Long Term Potentiation
			if n.Weights[id] > MAXSIG {
				n.Weights[id] = MAXSIG
			}
		}
	}

}

// Update a neuron potential whenever it receive a msg from a dendrite
func (n *Neuron) Update() {
	for n.Alive {
		select {
		case <-n.Clock.C:
			if n.Potential == 0 {
				break
			} else if n.Potential < 0 {
				n.Potential++ // slowly recovering
				break
			} else if n.Potential >= TRESH {
				n.Fire()
			} else {
				n.Potential -= DAMPING
			}

			if n.Log != nil {
				n.Log <- fmt.Sprintf("Nr%v: %v", n.ID, n.Potential)
			}

		case ID := <-n.Input:
			if n.Potential < 0 {
				break //can't receive signal during recovery
			}
			if ID < 0 { // for testing purposes we allow negative id
				n.Potential += MAXSIG
				break
			}

			n.Potential += n.Weights[ID]
			// an inibitory neuron cannot
			// make potential go lower than 0
			if n.Potential < 0 {
				n.Potential = 0
			}
		}
	}
	close(n.Input) //closing input should be catched by parents
}

// Sensor are sensitive neurons that fires at regular rates
type Sensor struct {
	N      []*Neuron
	values [][]int // excitation  represents the potential
	// that will be added to this sensor every tick
}
