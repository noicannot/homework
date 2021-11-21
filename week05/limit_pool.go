package main

import "sync"

type limitPoolManager struct {
	max     int
	tickets chan *struct{}
	lock    *sync.RWMutex
}

func NewLimitPoolManager(max int) *limitPoolManager {
	lpm := new(limitPoolManager)
	tickets := make(chan *struct{}, max)
	for i := 0; i < max; i++ {
		tickets <- &struct{}{}
	}
	lpm.max = max
	lpm.tickets = tickets
	lpm.lock = &sync.RWMutex{}
	return lpm
}

func (this *limitPoolManager) ReturnAll() {
	this.lock.Lock()
	defer this.lock.Unlock()
	if len(this.tickets) == 0 {
		for i := 0; i < this.max; i++ {
			this.tickets <- &struct{}{}
		}
	}
}

func (this *limitPoolManager) GetTicket() bool {
	this.lock.RLock()
	defer this.lock.RUnlock()
	select {
	case <-this.tickets:
		return true
	default:
		return false
	}
}

func (this *limitPoolManager) GetRemaind() int {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return len(this.tickets)
}
