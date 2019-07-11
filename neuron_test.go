package neurons

import (
	"testing"
	"time"
)

func TestFire(t *testing.T) {
	var logchan = NewLog(t.Log)
	n := NewNeuron()
	n1 := NewNeuron()
	n.Log, n1.Log = logchan, logchan
	Connect(n, n1)
	t.Log(n, n1)
	time.Sleep(5 * time.Second)
	n.Fire()
	time.Sleep(5 * time.Second)
}
