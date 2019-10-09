package main

import (
	"fmt"
	"math/rand"
)

// Network grouping of neurons accessible by arrays or map
type Network struct {
	NN         [][]*Neuron
	Nmap       map[int]*Neuron
	excited    map[int]*Neuron
	currenTick int
}

func (net *Network) String() string {
	return fmt.Sprintf("%v", net.NN)
}

// NewNetwork generates a neural network
func NewNetwork(shape []int) *Network {
	net := &Network{
		NN:      make([][]*Neuron, len(shape)),
		Nmap:    make(map[int]*Neuron),
		excited: make(map[int]*Neuron),
	}

	for i, s := range shape {
		net.NN[i] = make([]*Neuron, s)
		for j := 0; j < s; j++ {
			n := NewNeuron(net)
			net.NN[i][j] = n
			net.Nmap[n.ID] = n
		}
	}
	return net
}

// RandomConnectLayers connects neurons within layers and to previous layers in a random way
func (net *Network) RandomConnectLayers(probainter, probalayer, weight float64) {
	for i, layer := range net.NN {
		for j, post := range layer {
			for k, pre := range layer {
				if j == k {
					continue
				}
				if rand.Float64() <= probalayer {
					Connect(pre, post, weight, true)
				}
			}
			if i > 0 {
				for _, pre := range net.NN[i-1] {
					if rand.Float64() <= probainter {
						Connect(pre, post, weight, false)
					}
				}
			}
		}
	}
}

// Connect two neurons together
func Connect(pre, post *Neuron, weight float64, inib bool) {
	s := NewSynapse(pre, post, weight, inib)
	pre.Axones = append(pre.Axones, s)
	post.Dendrites = append(post.Dendrites, s)
}
