// Package neurons is an attempt at a neural network simulation
// with each neurons beeing a goroutine that receive signals of others
// via channels in a LTP style backpropagation of discrete signal
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/afauroux/neurons/gui"
	"github.com/afauroux/neurons/network"
)

// Getxt get some text from user
func Getxt() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	return text, err
}
func log(input chan string, nn *network.Network) {
	for msg := range input {
		fmt.Println(nn, msg)
	}
}

func main() {
	random := false
	var shape []int
	if random {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(6) + 2
		shape = make([]int, n)
		for i := 0; i < n; i++ {
			shape[i] = rand.Intn(5) + 1
		}
	} else {
		shape = []int{3, 1, 3}
	}
	fmt.Println(shape)
	logchan := make(chan string)
	nn := network.New(shape)
	for _, n := range nn.Nmap {
		n.Log = logchan
	}
	go log(logchan, nn)
	fmt.Println(nn)
	gui.CreateCanvas(nn)

}
