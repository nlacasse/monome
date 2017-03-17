package monome

import (
	"fmt"
	"log"

	"github.com/hypebeast/go-osc/osc"
)

const (
	gridServerPort = 45451
)

type KeyEvType int

const (
	KeyDown KeyEvType = iota
	KeyUp
)

type KeyEv struct {
	X int
	Y int
	T KeyEvType
}

type LEDType int

const (
	LEDOn LEDType = iota
	LEDOff
)

type Grid struct {
	s      *osc.Server
	c      *osc.Client
	prefix string
	ready  chan struct{}
	Ev     chan KeyEv
}

func NewGrid(port int32) *Grid {
	g := &Grid{
		s:     &osc.Server{Addr: fmt.Sprintf("127.0.0.1:%d", gridServerPort)},
		c:     osc.NewClient("127.0.0.1", int(port)),
		ready: make(chan struct{}),
		Ev:    make(chan KeyEv),
	}

	g.s.Handle("/sys/prefix", g.handlePrefix)
	g.s.Handle("/sys/size", g.handleSize)
	g.s.Handle("/sys/id", g.handleId)
	g.s.Handle("/sys/host", g.handleHost)
	g.s.Handle("/sys/rotation", g.handleRotation)
	g.s.Handle("/sys/disconnect", g.handleDisconnect)

	go g.s.ListenAndServe()

	g.c.Send(osc.NewMessage("/sys/port", int32(gridServerPort)))
	g.c.Send(osc.NewMessage("/sys/info", int32(gridServerPort)))

	//  else if address == keyAddr then handleKey(msg)
	//  else if address == encAddr then handleEnc(msg)
	//  else if address == tiltAddr then handleTilt(msg)

	return g
}

func (g *Grid) handlePrefix(msg *osc.Message) {
	log.Printf("handlePrefix: %v", msg)
	g.prefix = msg.Arguments[0].(string)
	g.s.Handle(g.prefix+"/grid/key", g.handleKey)
	g.ready <- struct{}{}
}

func (g *Grid) handleSize(msg *osc.Message) {
	log.Printf("handleSize: %v", msg)
}

func (g *Grid) handleId(msg *osc.Message) {
	log.Printf("handleId: %v", msg)
}

func (g *Grid) handleHost(msg *osc.Message) {
	log.Printf("handleHost: %v", msg)
}

func (g *Grid) handleRotation(msg *osc.Message) {
	log.Printf("handleRotation: %v", msg)
}

func (g *Grid) handleDisconnect(msg *osc.Message) {
	log.Printf("handleDisconnect: %v", msg)
}

func (g *Grid) handleKey(msg *osc.Message) {
	log.Printf("handleKey: %v", msg)
	g.Ev <- KeyEv{
		X: int(msg.Arguments[0].(int32)),
		Y: int(msg.Arguments[1].(int32)),
		T: KeyEvType(msg.Arguments[2].(int32)),
	}
	log.Printf("handleKey down here")
}

func (g *Grid) SetLED(x, y int, on bool) {
	onInt := 0
	if on {
		onInt = 1
	}
	g.c.Send(osc.NewMessage(g.prefix+"/grid/led/set", int32(x), int32(y), int32(onInt)))
}
