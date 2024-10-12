package api

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"net/http"
)

type Handler struct {
	writer *kafka.Writer
}

func NewHandler(w *kafka.Writer) *Handler {
	return &Handler{
		writer: w,
	}
}

func (h *Handler) Data(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	data := struct {
		A int
		B int
	}{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "decode body", http.StatusInternalServerError)
		return
	}
	sum := data.A + data.B

	sumByte, err := json.Marshal(sum)
	if err != nil {
		http.Error(w, "marshal sum", http.StatusInternalServerError)
		return
	}

	err = h.writer.WriteMessages(ctx, kafka.Message{Value: sumByte})
	if err != nil {
		http.Error(w, "write message in kafka", http.StatusInternalServerError)
		return
	}
}
