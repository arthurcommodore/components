package sse

type Dispatcher struct {
	hub *Hub
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		hub: GetHub(),
	}
}

func (d *Dispatcher) SendToHub(id string, data string) {
	evt := Event{
		ID:   id,
		Data: data,
	}
	d.hub.Send(id, evt)
}

func (d *Dispatcher) Brodcast(ids []string, data string) {
	for _, id := range ids {
		evt := Event{
			ID:   id,
			Data: data,
		}
		d.hub.Send(id, evt)
	}
}
