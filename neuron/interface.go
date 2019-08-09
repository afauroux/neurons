package neuron

import (
	"time"
)

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
