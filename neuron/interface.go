package neuron

import (
	"math/rand"
	"time"
)

// Neuron interface
type Neuron interface {
	Fire()   // Fire send a signal to all child neurons
	Update() // The main go routine that will run untill Neuron is killed
	GetID() int
	GetParent(int) Neuron
	GetParents() map[int]Neuron
	SetParent(int, Neuron)
	GetChild(int) Neuron
	GetChilds() map[int]Neuron
	SetChild(int, Neuron)
	SetWeight(int, int)
	GetPotential() int
	GetInput() chan int
	GetX() int
	GetY() int
	SetX(int)
	SetY(int)
	Kill() // to stop the update routine
}

// Connect two neurons together (weight==0 -> random value)
func Connect(pre, post Neuron, weight int) {
	post.SetParent(pre.GetID(), pre)
	if weight == 0 {
		weight = rand.Intn(MAXSIG*2) - MAXSIG
	}
	post.SetWeight(pre.GetID(), weight)
	pre.SetChild(post.GetID(), post)
}

// constantes and globals

// DT is the time between potential updates
const DT = 20 * time.Millisecond

// BUFFSIZE size of the input channels
const BUFFSIZE = 100

// TRESH potentials reaching this causes firing
const TRESH = 100

// DAMPING is the potential lost in one DT
const DAMPING = 5

// MAXSIG is the maximum signal strengh
const MAXSIG = 100

// LOWEND is the negative potential after firing
// that creates a cooldown time
const LOWEND = -3

// total nb of neurons, used to generate unique IDs
var nbNeurones = 0

// generateID generate IDs for new neurons
func generateID() int {
	nbNeurones++
	return nbNeurones
}
