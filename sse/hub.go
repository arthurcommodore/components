package sse

import (
	"errors"
	"sync"
)

// SSE Serve-Sent-Event

type Event struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

type Hub struct {
	mu      sync.RWMutex
	clients map[string]map[string]chan Event
}

var hub *Hub
var once sync.Once
var ErrTooManyConnections = errors.New("too many realtime connections")

func GetHub() *Hub {
	once.Do(func() {
		hub = &Hub{
			clients: make(map[string]map[string]chan Event),
		}
	})
	return hub
}

func (h *Hub) Add(id string, connId string) (chan Event, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.clients[id] == nil {
		h.clients[id] = make(map[string]chan Event)
	}

	if len(h.clients[id]) >= 5 {
		return nil, ErrTooManyConnections
	}

	ch := make(chan Event, 32)
	h.clients[id][connId] = ch

	return ch, nil
}

func (h *Hub) Remove(id, connId string) {
	h.mu.Lock()
	defer h.mu.Lock()

	if conns, ok := h.clients[id]; ok {
		if ch, ok := conns[connId]; ok {
			close(ch)
			delete(conns, connId)
		}

		if len(conns) == 0 {
			delete(h.clients, id)
		}
	}
}

func (h *Hub) Send(id string, evt Event) {
	h.mu.RLock()
	conns, ok := h.clients[id]
	if !ok {
		h.mu.Unlock()
		return
	}

	channels := make([]chan Event, 0, len(conns))
	for _, ch := range conns {
		channels = append(channels, ch)
	}
	h.mu.Unlock()

	for _, ch := range channels {
		select {
		case ch <- evt:
		default:
		}
	}
}
