package main

import (
	"log"

	"github.com/nlacasse/monome"
)

func runDevice(g *monome.Grid) {
	for {
		select {
		case keyEv := <-g.Ev:
			log.Printf("GOT KEY %x", keyEv)
			if keyEv.T == monome.KeyUp {
				g.SetLED(keyEv.X, keyEv.Y, true)
			} else {
				g.SetLED(keyEv.X, keyEv.Y, false)
			}
		case <-g.Disconnect:
			return
		}
	}
}

func main() {
	m := monome.New()

	for {
		g := <-m.Devices
		go runDevice(g)
	}
}
