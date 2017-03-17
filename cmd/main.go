package main

import (
	"log"

	"github.com/nlacasse/monome"
)

type dis struct{}

func (d *dis) HandleConnect(g *monome.Grid) {
	log.Printf("HandleConnect")

	for {
		keyEv := <-g.Ev
		log.Printf("GOT KEY %x", keyEv)
		if keyEv.T == monome.KeyUp {
			g.SetLED(keyEv.X, keyEv.Y, true)
		} else {
			g.SetLED(keyEv.X, keyEv.Y, false)
		}
	}
}

func (d *dis) HandleDisconnect(g *monome.Grid) {
	log.Printf("HandleDisconnect")
}

func main() {
	m := monome.New(&dis{})
	m.Start()
}
