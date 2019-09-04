package neuron

import (
	"github.com/afauroux/neurons/utils"
)

// ActionPotential is representing the action potential of a neuron as an integer value from 0 to TRESH
// also there is a buffer that helps keeping track of which neurons sended their signals
// in order to get the current potential
type ActionPotential struct {
	Value  int          // the action potential
	Buffer utils.Buffer // a buffer containing the received IDs with multiplicity corresponding to weight
}

// NewActionPotential create a new ActionPotential struct
func NewActionPotential() ActionPotential {
	return &ActionPotential{
		Value:  0,
		Buffer: []int{},
	}
}

// Recover is used to recover the potential after firing
// returns true is the potential is back at 0
func (p *ActionPotential) Recover() bool {
	p.Value++
	return p.Value == 0
}

// Raise is used to raise the potential and fill the buffer
func (p *ActionPotential) Raise(amount, ID int) {
	p.Value += amount
	for i := 0; i < amount; i++ {
		p.Buffer.Push(ID)
	}
}

func (p *ActionPotential) GetInBuffer(){
	inbuff := make(map[int]bool)
	
	for !n.Buffer.Empty() {
		inbuff[n.Buffer.Pop()] = true
	}

	return 
}

// Lower is used to lower the potential and empty the buffer
// return false if the buffer is totally depleted
func (p *ActionPotential) Lower(amount int) int {
	for i := amount; i > 0; i-- {
		if p.Buffer.Empty() {
			return -1
		}
		p.Value--
		p.Buffer.Pop()
	}
	return 1
}

// CalciumPotential is representing the Calcium potential of a neuron as an integer value from 0 to TRESH
// also there is a buffer that helps keeping track of which neurons sended their signals
// in order to get the current potential
type CalciumPotential struct {
	Value  int
	Buffer utils.Buffer
}
