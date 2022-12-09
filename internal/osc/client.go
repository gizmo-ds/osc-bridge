package osc

import (
	"net"

	"github.com/hypebeast/go-osc/osc"
)

func Send(addr string, packet osc.Packet) error {
	_addr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	conn, err := net.DialUDP("udp", nil, _addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	data, err := packet.MarshalBinary()
	if err != nil {
		return err
	}

	if _, err = conn.Write(data); err != nil {
		return err
	}
	return nil
}
