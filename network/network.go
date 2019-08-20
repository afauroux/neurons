package network

import (
	"errors"
	"fmt"
	"math/rand"

	N "github.com/afauroux/neurons/neuron"
)

// Network is the base class representing a neural net
type Network struct {
	Nmap map[int]*N.Neuron // a map for easy acces of a neuron struct from its ID
	Net  [][]*N.Neuron     // neural net structure
}

func (nn *Network) String() string {
	return fmt.Sprintf("%v", nn.Net)
}

// New create a new neural net with given shape
func New(shape []int) *Network {
	nn := &Network{
		Nmap: make(map[int]*N.Neuron),
		Net:  make([][]*N.Neuron, len(shape)),
	}
	for i, s := range shape {
		nn.Net[i] = make([]*N.Neuron, s)
		for j := 0; j < s; j++ {
			n := N.NewNeuron()
			nn.Nmap[n.ID] = n
			nn.Net[i][j] = n
		}
	}
	nn.DefaultConnect()
	return nn
}

// DefaultConnect will fully connect every neurons to all the previous layers ones
// and to all the others in the same layer but with inhibitory connections
func (nn *Network) DefaultConnect() {
	nn.Connect(nn.PeerConnect(-N.MAXSIG, N.MAXSIG, 100), nn.PeerConnect(-N.MAXSIG, 0, 0))
}

// PeerConnect is an helper functions that return a function
// that connect 2 neurons together with probability 'proba' (an int from 0 to 100)
// and with a weight choosen randomly between wmin a wmax
// it is used to provide argument for the network.Connect function
func (nn *Network) PeerConnect(wmin, wmax, proba int) func(*N.Neuron, *N.Neuron) {
	return func(pre, post *N.Neuron) {
		if rand.Intn(100) <= proba {
			N.Connect(nn.Nmap[pre.ID], nn.Nmap[post.ID])
			nn.Nmap[post.ID].Weights[pre.ID] = wmin + int(float64(wmax-wmin)*rand.Float64())
		}
	}
}

// Connect is creating all the links between neurons from this neural network.
// *funcParents* and *funcNeighbours* are two functions that will be executed
// between each neurons and the ones in respectively the previous or same layer.
// Those functions should take as imput two neurons *ID* (pre and post synaptic)
func (nn *Network) Connect(funcParents, funcNeighbours func(*N.Neuron, *N.Neuron)) {
	for i, layer := range nn.Net {
		for _, n := range layer {
			if i > 0 {
				for _, n2 := range nn.Net[i-1] {
					funcParents(n2, n)
				}
			}
			for _, n2 := range nn.Net[i] {
				if n2.ID != n.ID {
					funcNeighbours(n2, n)
				}
			}
		}
	}
}

// AddNeuron allows adding a neuron to a neural network
func (nn *Network) AddNeuron(n *N.Neuron, layer int) error {
	if layer == len(nn.Net)+1 {
		nn.Net = append(nn.Net, []*N.Neuron{n})
	} else if layer < len(nn.Net) {
		nn.Net[layer] = append(nn.Net[layer], n)
	} else {
		return errors.New("Can't add neuron to a layer which is non existing and not the next unfilled one")
	}
	nn.Nmap[n.ID] = n
	return nil
}
