package neuron

import (
	"time"
)

// ---------------------- Constantes and globals ------------------------------

// DT is the time between potential updates
const DT = 20 * time.Millisecond

// TRESH potentials reaching this causes firing
const TRESH = 100

// BUFFSIZE size of the input channels
const BUFFSIZE = TRESH

// DAMPING is the potential lost in one DT
const DAMPING = 1

// LTP is the gain or loss in weight associated with an event causing long term potentiation
const LTP = 5

// MAXSIG is the maximum signal strengh
const MAXSIG = 100

// LOWEND is the negative potential after firing
// that creates a cooldown time
const LOWEND = -10

// ----------------- Glial cells specific constants ---------------------------

// FOODREWARD is the reward associated with firing that represent the general activity of one neuron
// e.g. if FOODREWARD * DT = 20 seconds, and if no firing occur during thos 20s then the neuron
// will have to make new connections to augment its probability to get triggered, but if
// a lot of firing occured and n.Food > FOODSPLIT then the neuron will multiply
const FOODREWARD = 1000

// FOODBASE is the base level for food
const FOODBASE = 2 * FOODREWARD

// FOODSPLIT is the amount of food which will cause a neuron to split
const FOODSPLIT = 10000

// TRESHGLIA is the treshold above which a glia fires (starting a calcium like wave)
const TRESHGLIA = 100

// SPEEDGLIA is the time taken to raise the glia Potential by one (will dictate the wave propagation speed)
const SPEEDGLIA = 1

// ----------------- Globals and helper functions -----------------------------

// total nb of neurons, used to generate unique IDs
var nbNeurones = 0

// generateID generate IDs for new neurons or new glias (only one count)
func generateID() int {
	nbNeurones++
	return nbNeurones
}
