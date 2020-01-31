package neuron

import (
	"time"
)

// ---------------------- Constantes and globals ------------------------------

// DT is the time between potential updates
const DT = 20 * time.Millisecond

// TRESH potentials reaching this causes firing
const TRESH float64 = 100

// DAMPING is the exponential decrease time constant of neuron potentials
const DAMPING float64 = 10

// LTP is the gain in weight associated with an event causing long term potentiation
const LTP float64 = 1.2

// LTD is the loss in weight associated with an event causing long term depotentiation
const LTD float64 = 0.8

// MAXSIG is the maximum signal strengh
const MAXSIG float64 = 100

// LOWEND is the negative potential after firing
// that creates a cooldown time
const LOWEND float64 = -0.02

const eps = 0.01

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
