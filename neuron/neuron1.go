package neuron

import (
	"fmt"
	"time"
)

// Neuron the default brick
type Neuron struct {
	ID        int             // id (total number)
	Input     chan int        // every presynaptic Neuron sends his id
	Weights   map[int]int     // from the id we know the associated
	Parents   map[int]*Neuron // Neuron and weight
	Childs    map[int]*Neuron // all Neurons listening to this one
	Clock     *time.Ticker    //internal clock for potential updates
	Potential int
	Log       chan string
	Alive     bool
	X, Y      int // used for ploting (see gui.go)
}

// NewNeuron creates a neuron with default values
func NewNeuron() (n *Neuron) {
	n = &Neuron{
		ID:        generateID(),
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
	return n
}

// Fire a neuron when its potential is above the threshold
func (n *Neuron) Fire() {
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
