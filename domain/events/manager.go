package events

import "sync"

var em *eventManager
var once sync.Once
var mu sync.Mutex

type eventManager struct {
	Handlers map[string][]EventHandler
	Pipe     chan Event
}

func NewEventManager() *eventManager {
	once.Do(func() {
		em = &eventManager{
			Handlers: make(map[string][]EventHandler),
			Pipe:     make(chan Event),
		}
	})
	return em
}

func (e *eventManager) Subscribe(event Event, handler EventHandler) {
	mu.Lock()
	defer mu.Unlock()
	em.Handlers[event.Name()] = append(em.Handlers[event.Name()], handler)
}

func (e *eventManager) Emit(event Event) {
	em.Pipe <- event
}

func (e *eventManager) Start() {
	go func() {
		for event := range e.Pipe {
			e.handleEvent(event)
		}
	}()
}

func (e *eventManager) handleEvent(event Event) {
	mu.Lock()
	defer mu.Unlock()
	handlers := em.Handlers[event.Name()]
    for _, handler := range handlers {
        go handler(event)
    }
}