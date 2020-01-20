package network

import (
	"testing"
)

func printlog(t *testing.T, log chan string) {
	for msg := range log {
		t.Log(msg)
	}
}

// TestFires, basic test n1 fires and n2 damp back to 0
func TestFires(t *testing.T) {
	net := NewNetwork([]int{1, 1})
	net.RandomConnectLayers(1, 0, 0)
	//state := NewState()
	t.Log(net)

	net.Nmap[1].Fire()
	i := 0
	for !net.isIdle() && i < 200 {
		t.Log(net)
		t.Log(net.excited)
		net.Update()
		i++
	}
	t.Log(net)
}
