package neuron4

import (
	"testing"
	"time"
)

func printlog(t *testing.T, log chan string) {
	for msg := range log {
		t.Log(msg)
	}
}
func make2connectedNeurons(t *testing.T, weight int) (n1, n2 *Neuron) {
	logchan := make(chan string)
	n1 = NewNeuron()
	n2 = NewNeuron()
	n1.Log = logchan
	n2.Log = logchan
	Connect(n1, n2, weight, weight)
	go printlog(t, logchan)
	return n1, n2
}

func waitAllNeuronsOff(t *testing.T, timeout time.Duration, neurons ...*Neuron) {
	var delay time.Duration
	for _, n := range neurons {
		for n.Potential != 0 {
			time.Sleep(1 * time.Second)
			delay += 1 * time.Second
			if delay > timeout {
				t.Error("Error: time out")
			}
		}
	}
}

// TestFires, basic test n1 fires and n2 damp back to 0
func TestFires(t *testing.T) {
	n1, n2 := make2connectedNeurons(t, MAXSIG-1)
	n1.Fire()
	time.Sleep(DT)
	if n2.Potential == 0 {
		t.Errorf("The signal didn't pass between neuron 1 and neuron 2")
	}
	waitAllNeuronsOff(t, 10*time.Second, n1, n2)
}

// TestAccumulates, basic test n1 fires several times until n2 fires
func TestAccumulates(t *testing.T) {
	n1, n2 := make2connectedNeurons(t, MAXSIG/3)
	for i := 0; i < 5; i++ {
		n1.Fire()
		if i == 4 && n2.Potential >= 0 {
			t.Error("Error: n2 didn't fire")
		}
		if i == 5 && n2.Potential >= 0 {
			// assumes LOWEND * DT > Firing Delay which should allways be the case
			t.Error("Error: n2 should still be in recovering phase")
		}
		time.Sleep(2 * DT)
	}

	for _, den := range n2.Dendrites {
		n2.log("syn : %v", den)
	}
	n1.Fire()
	waitAllNeuronsOff(t, 10*time.Second, n1, n2)
}
