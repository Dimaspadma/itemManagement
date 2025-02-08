package sse

import (
	"fmt"
	"net/http"
	"time"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) BroadcastEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Accel-Buffering", "no")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")

	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
	}

	ctx := r.Context()

	fmt.Println("Client Connected")
	clientChan := make(chan DataEvent)
	h.service.Register(clientChan)

	for {
		select {
		case msg := <-clientChan:
			data := fmt.Sprintf("data: {\"time\": \"%s\", \"key\":\"%s\", \"action\": \"%s\"}\n\n", time.Now().Format(time.RFC850), msg.Key, msg.Action)
			w.Write([]byte(data))
			flusher.Flush()

		case <-ctx.Done():
			fmt.Println("Client disconnected!")
			h.service.Unregister(clientChan)
			return
		}
	}
}
