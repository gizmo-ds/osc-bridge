package osc

import (
	"path"

	"github.com/hypebeast/go-osc/osc"
)

type Dispatcher struct {
	patterns map[string][]string
	publish  func(*osc.Message, ...string)
}

func (d *Dispatcher) Dispatch(packet osc.Packet) {
	switch p := packet.(type) {
	default:
		return
	case *osc.Message:
		for pattern, keys := range d.patterns {
			if m, _ := path.Match(pattern, p.Address); m {
				for _, key := range keys {
					d.publish(p, key)
				}
			}
		}
	}
}
