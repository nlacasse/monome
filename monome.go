package monome

import (
	"fmt"

	"github.com/hypebeast/go-osc/osc"
)

const (
	sport = 45450
	cport = 12002
)

type Monome struct {
	s *osc.Server
	c *osc.Client
}

func New() *Monome {
	m := &Monome{
		s: &osc.Server{Addr: fmt.Sprintf("127.0.0.1:%d", sport)},
		c: osc.NewClient("127.0.0.1", cport),
	}
	m.s.Handle("/serialosc/device", m.handle)
	m.s.Handle("/serialosc/add", m.handle)

	m.c.Send(osc.NewMessage("/serialosc/list", "127.0.0.1", int32(sport)))
	m.c.Send(osc.NewMessage("/serialosc/notify", "127.0.0.1", int32(sport)))

	return m
}

func (m *Monome) handle(msg *osc.Message) {
	fmt.Printf("%+v\n", msg)
}

func (m *Monome) Start() {
	m.s.ListenAndServe()
}
