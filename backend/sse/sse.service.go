package sse

import "fmt"

type Service struct {
	data *Data
}

func NewService() *Service {
	return &Service{
		data: &Data{
			Clients:    make(map[chan DataEvent]bool),
			Register:   make(chan chan DataEvent),
			Unregister: make(chan chan DataEvent),
			Broadcast:  make(chan DataEvent),
		},
	}
}

func (s *Service) ManageClients() {
	for {
		select {
		case client := <-s.data.Register:
			// Tambah client ke list
			fmt.Println("manageClients register")
			s.data.Clients[client] = true
		case client := <-s.data.Unregister:
			// Hapus client dari list
			fmt.Println("manageClients unregister")
			delete(s.data.Clients, client)
			close(client) // Tutup channel biar gak bocor
		case message := <-s.data.Broadcast:
			// Kirim pesan ke semua client
			fmt.Println("manageClients broadcast", message)
			for client := range s.data.Clients {
				client <- message
			}
		}
	}
}

func (s *Service) Register(client chan DataEvent) {
	s.data.Register <- client
}

func (s *Service) Unregister(client chan DataEvent) {
	s.data.Unregister <- client
}

func (s *Service) Broadcast(key string, action string) {
	s.data.Broadcast <- DataEvent{
		Key:    key,
		Action: action,
	}
}
