package model

import (
	"sync"
	"sync/atomic"

	log "github.com/sirupsen/logrus"
)

type Counter struct {
	currentCounter int64
	waitGroup      *sync.WaitGroup
	doneChan       chan bool
	once           sync.Once
}

func Initialize() *Counter {
	counter := newCounter()
	return counter
}

func newCounter() *Counter {
	// For long term counter persistance between shutdowns, we can use a database to retrieve
	// the last counter value and set it in CurrentCounter in below struct.
	return &Counter{
		currentCounter: 0,
		waitGroup:      &sync.WaitGroup{},
		doneChan:       make(chan bool),
	}
}

func (c *Counter) AutoGenerateSequenceNumber(rcvChan chan int64) {
	defer c.waitGroup.Done()
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic in AutoGenerateSequenceNumber:", r)
			// Store the last counter value in database here.
		}
	}()
	select {
	case <-c.doneChan:
		return
	default:
		rcvChan <- atomic.AddInt64(&c.currentCounter, 1)
	}
}

func (c *Counter) GetCounter() int64 {
	rcvChan := make(chan int64)
	c.waitGroup.Add(1)
	go c.AutoGenerateSequenceNumber(rcvChan)
	return <-rcvChan
}

func (c *Counter) Stop() {
	c.once.Do(func() {
		c.waitGroup.Wait()
		close(c.doneChan)
	})
	// To persist the last counter between shutdowns, we can use a database to store the last counter value at this place.
}

func (c *Counter) GetCurrentCounter() int64 {
	return c.currentCounter
}
