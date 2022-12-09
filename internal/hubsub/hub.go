package hubsub

import (
	"sync"
)

type Hub[T any, Y SubConfig] struct {
	m    sync.RWMutex
	size int
	subs map[*Sub[T, Y]]struct{}
}

func NewHub[T any, Y SubConfig](size int) *Hub[T, Y] {
	return &Hub[T, Y]{
		size: size,
		subs: make(map[*Sub[T, Y]]struct{}),
	}
}

func (h *Hub[T, Y]) Sub(conf Y, handler func(*T, Y)) *Sub[T, Y] {
	h.m.Lock()
	defer h.m.Unlock()

	for s := range h.subs {
		if s.conf.Key() == conf.Key() {
			return nil
		}
	}
	sub := &Sub[T, Y]{
		conf:    conf,
		msg:     make(chan *T, h.size),
		quit:    make(chan struct{}),
		handler: handler,
	}
	h.subs[sub] = struct{}{}

	go sub.run()
	return sub
}

func (h *Hub[T, Y]) Publish(msg *T, key ...string) {
	h.m.RLock()
	for s := range h.subs {
		if len(key) == 0 || key[0] == s.conf.Key() {
			s.Publish(msg)
		}
	}
	h.m.RUnlock()
}

func (h *Hub[T, Y]) Configs() []Y {
	h.m.RLock()
	var subs []Y
	for s := range h.subs {
		subs = append(subs, s.conf)
	}
	h.m.RUnlock()
	return subs
}

func (h *Hub[T, Y]) UnSub(key string) {
	h.m.Lock()
	for s := range h.subs {
		if s.conf.Key() == key {
			close(s.quit)
			delete(h.subs, s)
		}
	}
	h.m.Unlock()
}
