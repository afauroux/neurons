package neuron

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/afauroux/neurons/utils"
)

func (n Neuron) log(str string, values ...interface{}) {
	if n.Log != nil {
		n.Log <- fmt.Sprintf(str, values...)
	}
}

// Neuron the default brick
type Neuron struct {
	ID        int              // id (total number of neurons)
	Input     chan int         // Neuron signal receiver
	Dendrites map[int]*Synapse // all channels (Neurons or Glia) listening to this neuron activity
	Axones    map[int]*Synapse // all channels (Neurons or Glia) listening to this neuron activity
	Potential int              // the action potential
	Clock     *time.Ticker     // internal clock for potential updates
	Food      int              // the reward associated with beeing often firing
	Buffer    utils.Buffer     // a buffer containing the received IDs with multiplicity corresponding to weight
	Log       chan string      // a log chanel used to print out info
	Alive     bool             // boolean used to kill the never ending update goroutine
}

func (n *Neuron) String() string {
	return fmt.Sprintf("n%v(%3d)", n.ID, n.Potential)
}

// NewNeuron creates a neuron with default values
func NewNeuron() (n *Neuron) {
	n = &Neuron{
		ID:        generateID(),
		Input:     make(chan int, BUFFSIZE),
		Dendrites: make(map[int]*Synapse),
		Axones:    make(map[int]*Synapse),
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
	n.log("n%v: -- fires --", n.ID)
	n.Potential = LOWEND // negative potential so neuron can fire only after some regeneration
	for _, synapse := range n.Axones {
		synapse.C <- n.ID // we send our ID to all listening chanels (Axones)
	}
	inbuff := make(map[int]bool)
	for !n.Buffer.Empty() {
		inbuff[n.Buffer.Pop()] = true
	}

	for id := range inbuff {
		n.Dendrites[id].Weight += LTP // Long Term Potentiation
		if n.Dendrites[id].Weight > MAXSIG {
			n.Dendrites[id].Weight = MAXSIG
		}
		n.log("n%v: new weight[%v]=%v", n.ID, id, n.Dendrites[id].Weight)

	}
	//time.AfterFu

}

// Starve ...
func (n *Neuron) Starve() {
	if n.Food < 0 {
		n.Food = 0
		randID := rand.Intn(n.ID) // ID of the potential should be smaller
		fmt.Println(randID)
		//Connect(nmap[randID], n)
	}
}

//RaisePotential is used to raise the potential and fill the buffer
func (n *Neuron) RaisePotential(amount, ID int) {
	n.Potential += amount
	if ID > 0 {
		for i := 0; i < amount; i++ {
			n.Buffer.Push(ID)
		}
	}
}

// LowerPotential is used to lower the potential and empty the buffer
// return false if the buffer is totally depleted
func (n *Neuron) LowerPotential(amount int) bool {
	for i := amount; i > 0; i-- {
		if n.Buffer.Empty() {
			return false
		}
		n.Potential--
		n.Buffer.Pop()
	}
	return true
}

// Update a neuron potential whenever it receive a msg from a dendrite (parent neuron)
func (n *Neuron) Update() {
	for n.Alive {
		select {
		case <-n.Clock.C:
			n.log("n%v: %v", n.ID, n.Potential)

			if n.Potential == 0 {
				n.Clock.Stop() // no need to do anything in between excitation so we turn the clock off
				// it will be restarted the next time the Input chanel will receive a signal
			} else if n.Potential < 0 {
				n.RaisePotential(1, -1) // slowly recovering
				break
			} else {
				n.LowerPotential(DAMPING)
			}

		case ID := <-n.Input:

			if n.Potential < 0 {
				break //can't receive signal during recovery
			}

			oldPot := n.Potential
			signal := n.Dendrites[ID].Weight

			n.log("n%v: n%v -> %v", n.ID, ID, signal)

			if signal < 0 { // inhibitory neuron case
				if oldPot > 0 { // good job, it is actually fighting an excitation
					n.Dendrites[ID].Weight -= LTP
				}
				n.LowerPotential(signal)
			} else { // excitatory case
				n.RaisePotential(signal, ID)
				if n.Potential >= 100 {
					n.Fire()
				} else if oldPot == 0 && n.Potential > 0 {
					// If the neuron was previously dormant we need to reset its ticker
					// we create a new ticker for triggering updates until the potential
					// is damped out or the neuron is fired
					n.Clock = time.NewTicker(DT)
					c := time.NewTimer(DT)
					c.Stop()

				}

			}
		}
	}
}
