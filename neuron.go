// Package neurons is an attempt at a neural network simulation
// with each neurons beeing a goroutine that receive signals of others
// via channels in a LTP style backpropagation of discrete signal
package neurons

import (
	"fmt"
	"sync"
	"time"
)

var mainclock = time.NewTicker(500 * time.Millisecond)

// Axone neurons output chanel
type Axone struct {
	pre, pos *Neuron
	C        chan int
}

// Synapse neurons imput chanels
type Synapse struct {
	pre, post *Neuron
	C         chan int
	weight    int
}

// Neuron the default brick
type Neuron struct {
	id      int             // id (total number)
	input   chan<- int      // every presynaptic Neuron sends his id
	outputs []chan int      // input chanels of all Neurons listening to this one
	weights map[int]int     // from the id we know the associated
	parents map[int]*Neuron // Neuron and weight

	clock          time.Ticker //internal clock for potential updates
	thresh         int
	damping        int
	potential      int
	verbose, alive bool
}

// NewNeuron creates a neuron with default values
func NewNeuron() *Neuron {
	return &Neuron{}
}

// Fire a neuron when its potential is above the threshold
func (n *Neuron) Fire() {
	for _, out := range n.outputs {
		out <- n.id
	}
}

// Update a neuron potential whenever it receive a msg from a dendrite
func (n *Neuron) Update() {
	for t := range n.clock.C {
		// lower the potential according to dunping and reads
		// inputs from the input chanel buffer
		fmt.Println(t)
	}
}

// Sensor are sensitive neurons that fires at regular rates
type Sensor struct {
	axone     *Axone
	potential int
	ticker    *time.Ticker
}

// NewAxone creates an axone with default values
func NewAxone(pre, pos *Neuron) *Axone {
	C := make(chan int)
	return &Axone{
		pre: pre,
		pos: pos,
		C:   C,
	}

}

// NewSensor creates a sensor with default values
func NewSensor() *Sensor {
	var dt = 500 * time.Millisecond
	var ticker = time.NewTicker(dt)
	return &Sensor{
		axone:     NewAxone(nil, nil),
		potential: 10,
		ticker:    ticker,
	}
}

func combine(inputs []<-chan int, output chan<- int) {
	var group sync.WaitGroup
	for i := range inputs {
		group.Add(1)
		go func(input <-chan int) {
			for val := range input {
				output <- val
			}
			group.Done()
		}(inputs[i])
	}
	go func() {
		group.Wait()
		close(output)
	}()
}
