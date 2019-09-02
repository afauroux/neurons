package neuron4

import (
	"time"
)

// ---------------------- Constantes and globals ------------------------------
const eps = 0.05

// DT is the time between potential updates
const DT = 20 * time.Millisecond

// TRESH potentials reaching this causes firing
const TRESH = 100.0

// BUFFSIZE size of the input channels
const BUFFSIZE = TRESH

// DAMPING is the potential lost (in percentage) in one DT
const DAMPING = 3 * DT

// LTP is the gain or loss in weight (in percentage) associated with an event causing long term potentiation
const LTP = 0.05

// LTD is the gain or loss in weight (in percentage) associated with an event causing long term potentiation
const LTD = 0.05

// MAXSIG is the maximum signal strengh
const MAXSIG = 100.0

// LOWEND is the negative potential after firing
// that creates a cooldown time
const LOWEND = -10.0

// ----------------- Glial cells specific constants ---------------------------

// FOODREWARD is the reward associated with firing that represent the general activity of one neuron
// e.g. if FOODREWARD * DT = 20 seconds, and if no firing occur during thos 20s then the neuron
// will have to make new connections to augment its probability to get triggered, but if
// a lot of firing occured and n.Food > FOODSPLIT then the neuron will multiply
const FOODREWARD = 1000.0

// FOODBASE is the base level for food
const FOODBASE = 2 * FOODREWARD

// FOODSPLIT is the amount of food which will cause a neuron to split
const FOODSPLIT = 10000.0

// TRESHGLIA is the treshold above which a glia fires (starting a calcium like wave)
const TRESHGLIA = 100.0

// DELAYGLIA is the time taken to raise the glia Potential to the maximum
const DELAYGLIA = 1.0

// DAMPINGGLIA is the damping half life of the glia potential
const DAMPINGGLIA = DAMPING

// ACTIVATION ...
const ACTIVATION = 0.80

// SENSITIVITY ...
const SENSITIVITY = 100

// ----------------- Globals and helper functions -----------------------------

// total nb of neurons, used to generate unique IDs
var nbNeurones = 0

// generateID generate IDs for new neurons or new glias (only one count)
func generateID() int {
	nbNeurones++
	return nbNeurones
}
