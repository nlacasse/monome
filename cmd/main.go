package main

import (
	"log"

	"github.com/nlacasse/monome"
)

type dis struct{}

func (d *dis) HandleDevice(g *monome.Grid) {
	log.Printf("HandleDevice")
}

func (d *dis) HandleConnect(g *monome.Grid) {
	log.Printf("HandleConnect")
}

func (d *dis) HandleDisconnect(g *monome.Grid) {
	log.Printf("HandleDisconnect")
}

func main() {
	m := monome.New(&dis{})
	m.Start()
}
