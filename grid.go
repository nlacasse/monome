package monome

import ()

type Grid struct {
	port int32
}

func NewGrid(port int32) *Grid {
	return &Grid{
		port: port,
	}
}
