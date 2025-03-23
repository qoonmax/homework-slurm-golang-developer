package main

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"os"
)

type key string

const (
	userIDKey    key = "userID"
	requestIDKey key = "requestID"
)

type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	// Извлекаем значения из контекста
	if userID, ok := ctx.Value(userIDKey).(string); ok {
		r.AddAttrs(slog.String("userID", userID))
	}
	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		r.AddAttrs(slog.String("requestID", requestID))
	}
	return h.Handler.Handle(ctx, r)
}

func main() {
	baseHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(&ContextHandler{Handler: baseHandler})
	slog.SetDefault(logger)

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, userIDKey, uuid.New().String())
	ctx = context.WithValue(ctx, requestIDKey, "req-"+uuid.New().String())

	slog.InfoContext(ctx, "Request received")

	service(ctx)
}

func service(ctx context.Context) {
	slog.InfoContext(ctx, "Service called")
	repository(ctx)
}

func repository(ctx context.Context) {
	slog.InfoContext(ctx, "Repository called")
}
