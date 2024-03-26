package events

type Event interface {
	Name() string
}


type EventHandler func(event Event)


