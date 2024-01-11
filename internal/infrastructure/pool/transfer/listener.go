package pool

type Listener struct {
	ch   chan Job
	size int
}

func NewListener(size int) *Listener {
	ch := make(chan Job, size)
	return &Listener{ch: ch, size: size}
}

func (listener *Listener) listen() {
	for i := 0; i < listener.size; i++ {
		go func(ch chan Job) {
			<-ch
		}(listener.ch)
	}
}
