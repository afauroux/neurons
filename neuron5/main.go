package main

import (
	"fmt"
	"time"
)

func main() {
	shape := []int{3, 3, 3}

	net := NewNetwork(shape)
	net.RandomConnectLayers(100, 0, 0)
	go func() {
		for {
			net.Update()
			//fmt.Println(net)
			fmt.Println(net.NN[0][0].Axones)
			time.Sleep(time.Millisecond * 30)
		}
	}()
	CreateCanvas(net)

}
