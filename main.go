package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/afauroux/neurons/gui"
	"github.com/afauroux/neurons/nnetwork"
)

// Getxt get some text from user
func Getxt() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	return text, err
}
func log(input chan string) {
	for msg := range input {
		fmt.Println(msg)
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
		shape = []int{3, 3, 3, 3}
	}
	fmt.Println(shape)
	nmap := nnetwork.MakeNeuralNetwork(shape, true)
	logchan := make(chan string)
	go log(logchan)
	nmap[0][1].Log = logchan
	nmap[0][2].Log = logchan
	gui.CreateCanvas(nmap)
}
