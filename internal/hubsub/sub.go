package hubsub

type Sub[T any, Y SubConfig] struct {
	conf    Y
	msg     chan *T
	quit    chan struct{}
	handler func(*T, Y)
}
type SubConfig interface {
	Key() string
}

func (s *Sub[T, Y]) Publish(msg *T) {
	s.msg <- msg
}

func (s *Sub[T, Y]) run() {
	for {
		select {
		case msg := <-s.msg:
			s.handler(msg, s.conf)
		case <-s.quit:
			return
		}
	}
}
