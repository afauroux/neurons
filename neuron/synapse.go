package neuron

import (
	"fmt"
	"math/rand"
)

// Synapse the link between all neurons cells
type Synapse struct {
	C         chan int //the receiving chanel of the post synaptic cell
	Weight    int      // the weight associated with this synapse
	pre, post int      // the id of the pre and post synaptic cells
}

func (s *Synapse) String() string {
	return fmt.Sprintf("s[%v:%v]: %v", s.pre, s.post, s.Weight)
}

// Connectable an interface to connect together different types of neurons, glias, sensors etc...
type Connectable interface {
	GetID() int                      // unique identifier among all cell types
	GetInput() chan int              // the Input chanel on which every cell listens
	GetAxons() *map[int]*Synapse     // the output Synapses
	GetDendrites() *map[int]*Synapse // the input Synapses
}

// Connect two connectable together (pre and post synaptic)
func Connect(pre, post Connectable, minweight, maxweight int) {
	s := Synapse{
		C:      post.GetInput(),
		Weight: minweight + int(float32(maxweight-minweight)*rand.Float32()),
	}

	(*post.GetDendrites())[pre.GetID()] = &s
	(*pre.GetAxons())[post.GetID()] = &s
}

// Strengthen or weaken a bunch of dendrites
func Strengthen(c Connectable, dweight int, ids ...int) {
	for id := range ids {
		s := (*c.GetDendrites())[id]
		s.Weight += dweight
		if s.Weight > MAXSIG {
			s.Weight = MAXSIG
		}
		if s.Weight < -MAXSIG {
			s.Weight = -MAXSIG
		}
	}
}
