package osc

import (
	"osc-bridge/internal/config"
	"osc-bridge/internal/hubsub"

	"github.com/hypebeast/go-osc/osc"
)

type Server struct {
	d      *Dispatcher
	server *osc.Server
}

func NewServer(addr string, hub *hubsub.Hub[osc.Message, config.Bridge]) *Server {
	d := &Dispatcher{
		patterns: make(map[string][]string),
		publish:  hub.Publish,
	}
	for _, b := range hub.Configs() {
		for _, p := range b.Patterns {
			d.patterns[p] = append(d.patterns[p], b.Key())
		}
	}
	return &Server{
		d: d,
		server: &osc.Server{
			Addr:       addr,
			Dispatcher: d,
		},
	}
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}
