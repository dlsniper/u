package main

import (
	"sync"
	"time"
)

var (
	buffer      []struct{}
	bufferMutex = &sync.Mutex{}
)

func addToChan(val struct{}, cnt chan<- struct{}) {
	bufferMutex.Lock()
	defer func() {
		bufferMutex.Unlock()
		if len(buffer) == 500{
			dumpBuffer("dump from count")
		}
	}()

	buffer = append(buffer, val)
}

func dumpBuffer(reason string) {
	bufferMutex.Lock()
	defer bufferMutex.Unlock()
	if len(buffer) == 0 {
		return
	}
	println(reason)
	buffer = []struct{}{}
}

func monitorBuffer() {
	interval := time.NewTicker(1 * time.Millisecond)
	defer interval.Stop()
	for {
		select {
		case <-interval.C:
				dumpBuffer("dump from interval")
		}
	}
}

func main() {
	chn := make(chan struct{}, 20000)
	go monitorBuffer()
	go func() {
		for i := 0; i < 999; i++ {
			addToChan(struct{}{}, chn)
		}
	}()

	time.Sleep(100 * time.Second)
}
