package sse

type DataEvent struct {
	Key    string
	Action string
}

type Data struct {
	Clients    map[chan DataEvent]bool
	Register   chan chan DataEvent
	Unregister chan chan DataEvent
	Broadcast  chan DataEvent
}
