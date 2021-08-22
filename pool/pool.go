package pool

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"
)

type Pool struct {
	timeout      time.Duration
	maxCount     int
	mu           sync.Mutex
	currentCount int
	list         chan *http.Client
}

func NewPool(maxCount int, timeout time.Duration) *Pool {
	p := &Pool{
		timeout:      timeout,
		maxCount:     maxCount,
		mu:           sync.Mutex{},
		currentCount: maxCount,
		list:         make(chan *http.Client, maxCount),
	}
	for i := 0; i < maxCount; i++ {
		p.list <- &http.Client{}
	}
	return p
}

func (p *Pool) Get() (*http.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), p.timeout)
	select {
	case c := <-p.list:
		p.mu.Lock()
		defer p.mu.Unlock()
		p.currentCount--
		return c, nil
	case <-ctx.Done():
		return nil, errors.New("ctx time out")
	}
}

func (p *Pool) Release(client *http.Client) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.currentCount >= p.maxCount {
		return errors.New("channel overflow")
	}
	p.list <- client
	p.currentCount++
	return nil
}
