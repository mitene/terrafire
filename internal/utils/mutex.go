package utils

/*
A Mutex is a mutual exclusion lock which uses a channel as lock.

This Mutex can be used in select statement, whereas sync.Mutex cannot.
*/
type Mutex struct {
	ch chan interface{}
}

func NewMutex() *Mutex {
	return &Mutex{ch: make(chan interface{}, 1)}
}

func (m *Mutex) Lock() chan interface{} {
	return m.ch
}

func (m *Mutex) UnLock() {
	<-m.ch
}

func ExampleMutex() {
	mtx := NewMutex()

	select {
	case mtx.Lock() <- nil:
		func() {
			defer mtx.UnLock()

			// do something
		}()
	}
}
