package network

import (
	N "github.com/afauroux/neurons/neuron"
)

// Network is tha base class representing a neural net
type Network struct {
	Nmap  map[int]*N.Neuron // a map for easy acces of a neuron struct from its ID
	NN    [][]int           // neural net structure
	Glias map[int]*N.Glia
}

func addNeuron(n)
