// Package nnetwork is an attempt at a neural network simulation
// with each neurons beeing a goroutine that receive signals of others
// via channels in a LTP style backpropagation of discrete signal
package nnetwork

import (
	"math"
	"time"
)

// Neuron2 same as Neuron but without internal clock
type Neuron2 struct {
	ID        int              // id (total number)
	Input     chan int         // every presynaptic Neuron sends his id
	Weights   map[int]int      // from the id we know the associated
	Parents   map[int]*Neuron2 // Neuron and weight
	Childs    map[int]*Neuron2 // input chanels of all Neurons listening to this one
	Thresh    int
	Damping   int
	Last      time.Time
	Potential int
	Log       chan string
	Alive     bool
	X, Y      int // used for ploting (see gui.go)

}

// NewNeuron2 creates a neuron with default values
func NewNeuron2() *Neuron2 {
	n := &Neuron2{
		ID:        nbNeurones,
		Input:     make(chan int, BUFFSIZE),
		Weights:   make(map[int]int),
		Parents:   make(map[int]*Neuron2),
		Childs:    make(map[int]*Neuron2),
		Potential: 0,
		Log:       nil,
		Last:      time.Now(),
		Alive:     true,
		X:         nbNeurones, // by default its a linear system
		Y:         0,
	}
	go n.Update()
	nbNeurones++
	return n
}

//Connect two neurons together (pre(n) and post synaptic)
func (n *Neuron2) Connect(post *Neuron2) {
	post.Parents[n.ID] = n
	post.Weights[n.ID] = MAXSIG //rand.Intn(MAXSIG*2) - MAXSIG
	n.Childs[post.ID] = post
}

// Fire a neuron when its potential is above the threshold
func (n *Neuron2) Fire() {
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
func (n *Neuron2) Update() {
	for n.Alive {
		ID := <-n.Input

		if n.Potential < 0 {
			continue //can't receive signal during recovery
		}
		n.Potential -= int(math.Trunc(time.Since(n.Last).Seconds() / DT.Seconds() * float64(DAMPING)))
		n.Potential += n.Weights[ID]
		// an inibitory neuron cannot
		// make potential go lower than 0
		if n.Potential < 0 {
			n.Potential = 0
		} else if n.Potential >= TRESH {
			n.Fire()
		}

	}
	close(n.Input) //closing input should be catched by parents
}

// Sensor2 are sensitive neurons that fires at regular rates
type Sensor2 struct {
	N      []*Neuron
	values [][]int // excitation  represents the potential
	// that will be added to this sensor every tick
}
