package neuron

import (
	"fmt"
	"time"
)

// Neuron1 the default brick
type Neuron1 struct {
	ID        int            // id (total number)
	Input     chan int       // every presynaptic Neuron sends his id
	Weights   map[int]int    // from the id we know the associated
	Parents   map[int]Neuron // Neuron and weight
	Childs    map[int]Neuron // input chanels of all Neurons listening to this one
	Clock     *time.Ticker   //internal clock for potential updates
	Thresh    int
	Damping   int
	Potential int
	Log       chan string
	Alive     bool
	X, Y      int // used for ploting (see gui.go)

}

// NewNeuron1 creates a neuron with default values
func NewNeuron1() (n Neuron1) {
	n = Neuron1{
		ID:        nbNeurones,
		Input:     make(chan int, BUFFSIZE),
		Weights:   make(map[int]int),
		Parents:   make(map[int]Neuron),
		Childs:    make(map[int]Neuron),
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

// Fire a neuron when its potential is above the threshold
func (n Neuron1) Fire() {
	n.Potential = LOWEND // negative potential so neuron can fire only after some regeneration
	for _, post := range n.Childs {
		post.GetInput() <- n.ID
	}
	for id, p := range n.Parents {
		if p.GetPotential() < 0 { //fired recently
			n.Weights[id]++ // Long Term Potentiation
			if n.Weights[id] > MAXSIG {
				n.Weights[id] = MAXSIG
			}
		}
	}

}

// Update a neuron potential whenever it receive a msg from a dendrite
func (n Neuron1) Update() {
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

// GetID in order to implement Neuron interface
func (n Neuron1) GetID() int {
	return n.ID
}

// SetParent in order to implement Neuron interface
func (n Neuron1) SetParent(ID int, parent Neuron) {
	n.Parents[ID] = parent
}

// SetChild in order to implement Neuron interface
func (n Neuron1) SetChild(ID int, child Neuron) {
	n.Childs[ID] = child
}

// SetWeight in order to implement Neuron interface
func (n Neuron1) SetWeight(ID int, weight int) {
	n.Weights[ID] = weight
}

// GetParent in order to implement Neuron interface
func (n Neuron1) GetParent(ID int) (parent Neuron) {
	return n.Parents[ID]
}

// GetParents in order to implement Neuron interface
func (n Neuron1) GetParents() (parents map[int]Neuron) {
	return n.Parents
}

// GetChilds in order to implement Neuron interface
func (n Neuron1) GetChilds() (childs map[int]Neuron) {
	return n.Childs
}

// GetChild in order to implement Neuron interface
func (n Neuron1) GetChild(ID int) (child Neuron) {
	return n.Childs[ID]
}

// GetWeight in order to implement Neuron interface
func (n Neuron1) GetWeight(ID int) (weight int) {
	return n.Weights[ID]
}

// GetPotential in order to implement Neuron interface
func (n Neuron1) GetPotential() (Potential int) {
	return n.Potential
}

// GetInput in order to implement Neuron interface
func (n Neuron1) GetInput() (Input chan int) {
	return n.Input
}

// GetX return corresponding value
func (n Neuron1) GetX() int {
	return n.X
}

// GetY return corresponding value
func (n Neuron1) GetY() int {
	return n.Y
}

// SetX sets corresponding value
func (n Neuron1) SetX(X int) {
	n.X = X
}

// SetY sets corresponding value
func (n Neuron1) SetY(Y int) {
	n.Y = Y
}

// Kill stop the neuron activity
func (n Neuron1) Kill() {
	n.Alive = false
}
