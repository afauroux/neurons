package neuron

import (
	"fmt"
	"math/rand"
	"time"
)

// Neuron the default brick
type Neuron struct {
	ID        int             // id (total number of neurons)
	Input     chan int        // every presynaptic Neuron sends his id
	Weights   map[int]int     // from the id we know the corresponding
	Parents   map[int]*Neuron // Neuron and its associated synaptic weight
	Childs    map[int]*Neuron // all Neurons listening to this one
	Clock     *time.Ticker    // internal clock for potential updates
	Potential int             // the action potential
	Food      int             // the reward function associated with beeing often firing
	Log       chan string     // a log chanel used to print out info
	Alive     bool            // boolean used to kill the never ending update goroutine
	layer     int             // layer number if the neurone is in a multilayer network
}

// Connect two neurons together (pre and post synaptic)
func Connect(pre, post *Neuron) {
	post.Parents[pre.ID] = pre
	post.Weights[pre.ID] = rand.Intn(2*MAXSIG) - MAXSIG
	pre.Childs[post.ID] = post
}

// New creates a neuron with default values
func New() (n *Neuron) {
	n = &Neuron{
		ID:        generateID(),
		Input:     make(chan int, BUFFSIZE),
		Weights:   make(map[int]int),
		Parents:   make(map[int]*Neuron),
		Childs:    make(map[int]*Neuron),
		Clock:     time.NewTicker(DT),
		Potential: 0,
		Food:      0,
		Log:       nil,
		Alive:     true,
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
			n.Weights[id] += LTP // Long Term Potentiation
			if n.Weights[id] > MAXSIG {
				n.Weights[id] = MAXSIG
			}
		}
	}
	n.Food += FOODREWARD
	time.AfterFunc(time.Duration(n.Food)*DT, n.Starve)

}

// Starve ...
func (n *Neuron) Starve() {
	if n.Food < 0 {
		n.Food = 0
		randID := rand.Intn(n.ID) // ID of the potential should be smaller
		Connect(nmap[randID], n)
	}
}

// Update a neuron potential whenever it receive a msg from a dendrite
func (n *Neuron) Update() {
	for n.Alive {
		select {
		case <-n.Clock.C:
			if n.Potential == 0 {
				n.Clock.Stop() // no need to do anything in between excitation so we turn the clock off
				// it will be restarted the next time the Input chanel will receive a signal
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

			oldPot := n.Potential
			if ID < 0 {
				// to excite a neuron artificially a negative iD could be sent
				// this is used in gui.go to allow user to trigger neurons
				n.Potential += MAXSIG
			} else {
				// action of a pre-synaptic neuron onto this one throught a weighted synapse
				n.Potential += n.Weights[ID]

				if n.Weights[ID] < 0 { // inhibitory neuron case
					if oldPot > 0 { // good job, it is actually fighting an excitation
						n.Weights[ID] -= LTP
					}
				}
			}

			// an inibitory neuron cannot
			// make potential go lower than 0
			if n.Potential <= 0 {
				n.Potential = 0
			} else {
				// a new ticker for triggering updates until the potential
				// is damped out or the neuron is fired
				n.Clock = time.NewTicker(DT)
			}
		}
	}
}
