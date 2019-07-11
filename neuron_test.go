package neurons

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

var logchan = NewLog()

// NewLog creates a one line logger
func NewLog() chan string {
	tick := time.NewTicker(dtNeurones)
	var logchan = make(chan string)
	register := make(map[string]string) // each neuron got one only field
	go func() {

		for {
			select {
			case t := <-tick.C:
				str := ""
				for _, v := range register {
					str += "/" + v + "/"
				}

				fmt.Print(t.Format("15:04:05") + " -> " + str + "\r")
			case msg := <-logchan:
				id := strings.Split(msg, ":")[0]
				register[id] = msg
			}

		}
	}()
	return logchan
}

func TestFire(t *testing.T) {
	n := NewNeuron()
	n1 := NewNeuron()
	Connect(n, n1)
	fmt.Println(n, n1)
	time.Sleep(5 * time.Second)
	n.Fire()
	time.Sleep(5 * time.Second)
}
