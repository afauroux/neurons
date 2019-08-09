package neuron

import (
	"math"
	"time"
)

// Neuron2 same as Neuron but without internal clock
type Neuron2 struct {
	ID        int            // id (total number)
	Input     chan int       // every presynaptic Neuron sends his id
	Weights   map[int]int    // from the id we know the associated
	Parents   map[int]Neuron // Neuron and weight
	Childs    map[int]Neuron // input chanels of all Neurons listening to this one
	Last      time.Time
	Potential int
	Log       chan string
	Alive     bool
	X, Y      int // used for ploting (see gui.go)

}

// NewNeuron2 creates a neuron with default values
func NewNeuron2() (n *Neuron2) {
	//var last [10]time.Time
	//last[0] = time.Now()
	n = &Neuron2{
		ID:        generateID(),
		Input:     make(chan int, BUFFSIZE),
		Weights:   make(map[int]int),
		Parents:   make(map[int]Neuron),
		Childs:    make(map[int]Neuron),
		Potential: 0,
		Log:       nil,
		Last:      time.Now(),
		Alive:     true,
		X:         0,
		Y:         0,
	}
	go n.Update()
	return n
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
