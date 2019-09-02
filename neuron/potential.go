package neuron

type Excitable interface {
	Update()
	RaisePotential(*Synapse) bool
	LowerPotential(int) bool
	Fire()
}
